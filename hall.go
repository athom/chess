package chess

import (
	"fmt"
	"sync"
)

func NewHall() (r *Hall) {
	r = &Hall{
		Joins: make(chan MailBox),
	}
	r.run()
	return
}

type Hall struct {
	sync.Mutex
	game    *Game
	rooms   []*Room
	players []*Player
	Joins   chan MailBox
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

func (this *Hall) Join(conn MailBox) {
	this.Lock()
	defer this.Unlock()
	player := NewPlayer(conn)
	this.players = append(this.players, player)
	this.matchPlayer(player)
}

func (this *Hall) JoinRoom(conn MailBox, slug string) {
	this.Lock()
	defer this.Unlock()
	player := NewPlayer(conn)
	this.players = append(this.players, player)
	this.matchRoomPlayer(player, slug)
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

func (this *Hall) matchRoomPlayer(player *Player, slug string) {
	for _, room := range this.rooms {
		if room.slug == slug {
			room.Join(player)
			return
		}
	}

	this.rooms = append(this.rooms, NewRoomWithSlug(slug, player))
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
