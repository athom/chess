package chess

import "gopkg.in/mgo.v2/bson"

func NewPlayer(mb MailBox) (r *Player) {
	r = &Player{
		id:          bson.NewObjectId().Hex(),
		mailBox:     mb,
		playerState: make(chan *PlayerState),
		gameState:   make(chan *GameState),
	}
	r.run()
	return
}

type Player struct {
	id          string
	side        Side
	playerState chan *PlayerState
	gameState   chan *GameState
	mailBox     MailBox
}

func (this *Player) IsWatcher() bool {
	return this.side != WHITE && this.side != BLACK
}

func (this *Player) Ready() {
	this.playerState <- &PlayerState{Id: this.id, State: IN_READY}
}

func (this *Player) Move() {
	this.playerState <- &PlayerState{Id: this.id, State: IN_MOVE}
}

func (this *Player) GiveUp() {
	this.playerState <- &PlayerState{Id: this.id, State: IN_GIVEUP}
}

func (this *Player) AbortGame() {
	this.playerState <- &PlayerState{Id: this.id, State: IN_ABORT}
}

func (this *Player) execute() {
	for {
		ps := this.mailBox.Receive()
		this.playerState <- ps
		ps.Id = this.id
		ps.Side = this.side
		if ps.State == IN_ABORT {
			break
		}
	}
}

func (this *Player) report() {
	for gs := range this.gameState {
		this.mailBox.Send(gs)
	}
}

func (this *Player) run() {
	go this.execute()
	go this.report()
}
