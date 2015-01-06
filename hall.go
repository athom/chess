package chess

import (
	"fmt"
	"net"
	"sync"
)

func NewHall() (r *Hall) {
	r = &Hall{
		Joins: make(chan net.Conn),
	}
	r.run()
	return
}

type Hall struct {
	sync.Mutex
	game    *Game
	rooms   []*Room
	players []*Player
	Joins   chan net.Conn
}

func (this *Hall) Players() []*Player {
	this.Lock()
	defer this.Unlock()
	return this.players
}

func (this *Hall) Rooms() []*Room {
	this.Lock()
	defer this.Unlock()
	return this.rooms
}

func (this *Hall) Join(conn net.Conn) {
	this.Lock()
	defer this.Unlock()
	player := NewPlayer(conn)
	this.players = append(this.players, player)
	this.matchPlayer(player)
}

func (this *Hall) matchPlayer(player *Player) {
	for _, room := range this.rooms {
		if room.JoinableForPlay() {
			room.Join(player)
			return
		}
	}
	this.rooms = append(this.rooms, NewRoom(player))
}

func (this *Hall) run() {
	go func() {
		for {
			select {
			case conn := <-this.Joins:
				fmt.Println("Join")
				this.Join(conn)
			}
		}
	}()
}
