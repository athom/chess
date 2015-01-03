package chess

import (
	"sync"

	"gopkg.in/mgo.v2/bson"
)

func NewRoom(players ...*Player) (r *Room) {
	r = &Room{}
	r.id = bson.NewObjectId().Hex()
	for _, p := range players {
		r.Join(p)
	}
	r.run()
	return
}

type Room struct {
	sync.RWMutex
	id          string
	game        *Game
	playerBlack *Player
	playerWhite *Player
	watchers    []*Player
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
	defer this.initGame()

	if this.playerBlack == nil {
		this.playerBlack = player
		return
	}

	if this.playerWhite == nil {
		this.playerWhite = player
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

func (this *Room) run() {
	go func() {
	}()
}
