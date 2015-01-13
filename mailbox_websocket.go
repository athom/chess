package chess

import (
	"io"

	"golang.org/x/net/websocket"
)

func NewWebsocketMailBox(ws *websocket.Conn) (r *WebsocketMailBox) {
	r = &WebsocketMailBox{
		ws:     ws,
		parser: NewPlayerStateParser(),
		ps:     make(chan *PlayerState),
	}
	return
}

type Move struct {
	FromPos Pos `json:"from_pos"`
	ToPos   Pos `json:"to_pos"`
}

type WebsocketMailBox struct {
	ws     *websocket.Conn
	parser *PlayerStateParser
	ps     chan *PlayerState
}

func (this *WebsocketMailBox) Receive() (r *PlayerState) {
	for ps := range this.ps {
		return ps
	}
	return
}

func (this *WebsocketMailBox) Send(gs *GameState) {
	this.ws.Write(gs.ToJson())
	return
}

func (this *WebsocketMailBox) Run() {
	for {
		mi := &MoveInfo{}
		err := websocket.JSON.Receive(this.ws, &mi)
		if err != nil && err == io.EOF {
			ps := &PlayerState{
				State: IN_ABORT,
			}
			this.ps <- ps
			break
		}

		ps := &PlayerState{
			State:    IN_MOVE,
			MoveInfo: mi,
		}
		this.ps <- ps
	}
}
