package chess

import (
	"fmt"
	"net"
	"sync"
)

func NewHall() (r *Hall) {
	r = &Hall{
		joins: make(chan net.Conn),
	}
	r.run()
	return
}

type Hall struct {
	sync.Mutex
	cmdParser *CmdParser
	game      *Game
	players   []*Player
	joins     chan net.Conn
}

func (this *Hall) Players() []*Player {
	this.Lock()
	defer this.Unlock()
	return this.players
}

func (this *Hall) Join(conn net.Conn) {
	this.Lock()
	defer this.Unlock()
	player := NewPlayer(conn)
	this.players = append(this.players, player)
}

func (this *Hall) MatchPlayer(p *Player) {
}

func (this *Hall) run() {
	go func() {
		for {
			select {
			case conn := <-this.joins:
				fmt.Println("Join")
				this.Join(conn)
			}
		}
	}()
}
