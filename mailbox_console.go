package chess

import (
	"bufio"
	"net"
)

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
