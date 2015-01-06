package chess

import (
	"bufio"
	"net"

	"gopkg.in/mgo.v2/bson"
)

type MailBox interface {
	Receive() *PlayerState
	Send(*GameState)
}

func NewConsoleMailBox(conn net.Conn) (r MailBox) {
	r = &ConsoleMailBox{
		conn:   conn,
		reader: bufio.NewReader(conn),
		parser: NewPlayerStateParser(),
	}
	return
}

type ConsoleMailBox struct {
	conn   net.Conn
	reader *bufio.Reader
	parser *PlayerStateParser
}

func (this *ConsoleMailBox) Receive() (r *PlayerState) {
	line, err := this.reader.ReadString('\n')
	if err != nil {
		r = &PlayerState{State: IN_READY}
		return
	}

	r = this.parser.Parse(line)
	return
}
func (this *ConsoleMailBox) Send(gs *GameState) {
	this.conn.Write(gs.ToJson())
	return
}

func NewPlayer(conn net.Conn) (r *Player) {
	r = &Player{
		id:          bson.NewObjectId().Hex(),
		mailBox:     NewConsoleMailBox(conn),
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
