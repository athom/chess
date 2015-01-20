package chess

import (
	"sync"

	"gopkg.in/mgo.v2/bson"
)

func NewRoom(players ...*Player) (r *Room) {
	return NewRoomWithSlug("", players...)
}

func NewRoomWithSlug(slug string, players ...*Player) (r *Room) {
	r = &Room{
		id:          bson.NewObjectId().Hex(),
		slug:        slug,
		playerState: make(chan *PlayerState),
		gameState:   make(chan *GameState),
	}
	for _, p := range players {
		r.Join(p)
	}
	r.run()
	return
}

type Room struct {
	sync.RWMutex
	id          string
	slug        string
	game        *Game
	playerBlack *Player
	playerWhite *Player
	watchers    []*Player
	playerState chan *PlayerState
	gameState   chan *GameState
}

func (this *Room) JoinableForPlay() (r bool) {
	this.RLock()
	defer this.RUnlock()
	return this.playerBlack == nil || this.playerWhite == nil
}

func (this *Room) BlackPlayer() *Player {
	this.RLock()
	defer this.RUnlock()
	return this.playerBlack
}

func (this *Room) WhitePlayer() *Player {
	this.RLock()
	defer this.RUnlock()
	return this.playerWhite
}

func (this *Room) Watchers() []*Player {
	this.RLock()
	defer this.RUnlock()
	return this.watchers
}

func (this *Room) initGame() {
	if this.playerBlack != nil && this.playerWhite != nil {
		if this.game == nil {
			this.game = NewGame(6, NewTextFormatter())
		} else {
			this.game.reset()
		}
	} else {
		this.game = nil
	}
}

func (this *Room) Join(player *Player) {
	this.Lock()
	defer this.Unlock()

	this.connectPlayerState(player)

	if this.playerBlack == nil {
		player.side = BLACK
		//player.Ready()
		this.playerBlack = player
		this.broadcastWait()
		return
	}

	if this.playerWhite == nil {
		player.side = WHITE
		//player.Ready()
		this.playerWhite = player
		this.broadcastReady()
		return
	}

	this.watchers = append(this.watchers, player)
	return
}

func (this *Room) Leave(player *Player) {
	this.Lock()
	defer this.Unlock()
	defer this.initGame()
	if this.playerBlack == player {
		this.playerBlack = nil
		// turn white player to black when back is left and white still there
		if this.playerWhite != nil {
			this.playerBlack = this.playerWhite
			this.playerBlack.side = BLACK
			this.playerWhite = nil
		}
		return
	}
	if this.playerWhite == player {
		this.playerWhite = nil
	}
	for i, watcher := range this.watchers {
		if player == watcher {
			this.watchers = append(this.watchers[0:i], this.watchers[i+1:]...)
			return
		}
	}
}

func (this *Room) connectPlayerState(player *Player) {
	go func() {
		for {
			this.playerState <- <-player.playerState
		}
	}()
}

func (this *Room) run() {
	go func() {
		for {
			select {
			case ps := <-this.playerState:
				if ps == nil {
					continue
				}
				this.handlePlayerState(ps)
			}
		}
	}()
}

func (this *Room) handlePlayerState(ps *PlayerState) {
	switch ps.State {
	case IN_ABORT:
		this.broadcastLeave(ps.Id)
	case IN_MOVE:
		err := this.game.Move(ps.MoveInfo.FromPos, ps.MoveInfo.ToPos, ps.Side)
		if err == nil {
			this.broadcastMove()
			return
		}
		if err == ErrGameOverBlackWin {
			this.broadcastGameOver(BLACK)
			return
		}
		if err == ErrGameOverWhiteWin {
			this.broadcastGameOver(WHITE)
			return
		}
		player := this.findPlayer(ps.Id)
		if player == nil {
			return
		}
		player.gameState <- &GameState{
			State:       OUT_ILLEGAL_OPERATION,
			MyBoardInfo: NewMyBoardInfo(this.game.BoardInfo(), player.side),
		}
	case IN_ILLEAGLE_OPERATION:
		player := this.findPlayer(ps.Id)
		if player == nil {
			return
		}
		player.gameState <- &GameState{
			State:       OUT_ILLEGAL_OPERATION,
			MyBoardInfo: NewMyBoardInfo(this.game.BoardInfo(), player.side),
		}
	}
}

func (this *Room) broadcastWait() {
	this.playerBlack.gameState <- &GameState{
		State: OUT_WAIT,
	}
}
func (this *Room) broadcastReady() {
	this.initGame()

	this.playerBlack.gameState <- &GameState{
		State:       OUT_READY,
		MyBoardInfo: NewMyBoardInfo(this.game.BoardInfo(), BLACK),
	}

	if this.playerWhite != nil {
		this.playerWhite.gameState <- &GameState{
			State:       OUT_READY,
			MyBoardInfo: NewMyBoardInfo(this.game.BoardInfo(), WHITE),
		}
	}
}

func (this *Room) broadcastMove() {
	this.playerBlack.gameState <- &GameState{
		State:       OUT_BOARD_UPDATED,
		MyBoardInfo: NewMyBoardInfo(this.game.BoardInfo(), BLACK),
	}
	this.playerWhite.gameState <- &GameState{
		State:       OUT_BOARD_UPDATED,
		MyBoardInfo: NewMyBoardInfo(this.game.BoardInfo(), WHITE),
	}
	for _, watcher := range this.watchers {
		watcher.gameState <- &GameState{
			State:       OUT_BOARD_UPDATED,
			MyBoardInfo: NewMyBoardInfo(this.game.BoardInfo(), NONE),
		}
	}
}

func (this *Room) broadcastGameOver(winSide Side) {
	blackOverState := OUT_GAMEOVER_WIN
	whiteOverState := OUT_GAMEOVER_LOSE
	if winSide == WHITE {
		blackOverState = OUT_GAMEOVER_LOSE
		whiteOverState = OUT_GAMEOVER_WIN
	}

	this.playerBlack.gameState <- &GameState{
		State:       blackOverState,
		MyBoardInfo: GameOverBoardInfo(this.game.BoardInfo(), BLACK),
	}
	this.playerWhite.gameState <- &GameState{
		State:       whiteOverState,
		MyBoardInfo: GameOverBoardInfo(this.game.BoardInfo(), WHITE),
	}
	for _, watcher := range this.watchers {
		watcher.gameState <- &GameState{
			State:       OUT_GAMEOVER_FOR_WATCHER,
			MyBoardInfo: GameOverBoardInfo(this.game.BoardInfo(), NONE),
		}
	}
}

func (this *Room) broadcastLeave(id string) {
	this.playerBlack.gameState <- &GameState{
		State: OUT_OPPOENENT_ABORT,
	}

	if this.playerWhite != nil {
		this.playerWhite.gameState <- &GameState{
			State: OUT_OPPOENENT_ABORT,
		}
	}
	for _, watcher := range this.watchers {
		watcher.gameState <- &GameState{
			State: OUT_OPPOENENT_ABORT,
		}
	}

	player := this.findPlayer(id)
	this.Leave(player)
}

func (this *Room) findPlayer(id string) (r *Player) {
	if this.BlackPlayer().id == id {
		r = this.BlackPlayer()
		return
	}
	if this.WhitePlayer().id == id {
		r = this.WhitePlayer()
		return
	}
	for _, watcher := range this.Watchers() {
		if watcher.id == id {
			r = watcher
			return
		}
	}
	return
}
