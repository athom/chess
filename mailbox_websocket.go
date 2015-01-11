package chess

import "golang.org/x/net/websocket"

func NewWebsocketMailBox(ws *websocket.Conn) (r *WebsocketMailBox) {
	r = &WebsocketMailBox{
		ws:     ws,
		parser: NewPlayerStateParser(),
		In:     make(chan []byte),
	}
	return
}

type WebsocketMailBox struct {
	ws     *websocket.Conn
	parser *PlayerStateParser
	In     chan []byte
}

func (this *WebsocketMailBox) Receive() (r *PlayerState) {
	for msg := range this.In {
		r = this.parser.Parse(string(msg))
		return
	}
	return
}
func (this *WebsocketMailBox) Send(gs *GameState) {
	this.ws.Write(gs.ToJson())
	return
}

func (this *WebsocketMailBox) Run() {
	for {
		var b []byte
		this.ws.Read(b)
		if b != nil {
			this.In <- b
		}
	}
}
