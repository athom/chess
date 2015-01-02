package main

import (
	"bufio"
	"fmt"
	"net"
	"time"

	"github.com/athom/chess"
	"gopkg.in/mgo.v2/bson"
)

type Client struct {
	id string
	//incoming  chan string
	incoming chan *clientIncomming
	outgoing chan *chess.Message
	reader   *bufio.Reader
	writer   *bufio.Writer
	moveable bool
	Side     chess.Side
}

func (this *Client) Read() {
	for {
		line, err := this.reader.ReadString('\n')
		if err != nil {
			continue
		}

		var event = GOING
		if line == "QUIT\n" {
			event = CLOSED
		}
		ci := &clientIncomming{
			clientId: this.id,
			event:    event,
			content:  line,
		}
		this.incoming <- ci
	}
}

func (this *Client) Write() {
	for data := range this.outgoing {
		//client.writer.WriteString(data)
		//log.Println(data)
		//log.Println(data.ToJson())
		this.writer.Write(data.ToJson())
		time.Sleep(100 * time.Millisecond) //TODO solve the sent togather bug, and remove this time waiting magic
		this.writer.Flush()
	}
}

func (this *Client) Ready() {
	ci := &clientIncomming{
		clientId: this.id,
		event:    READY,
	}
	this.incoming <- ci
}

func (this *Client) Listen() {
	go this.Read()
	go this.Write()
}

func NewClient(connection net.Conn) *Client {
	writer := bufio.NewWriter(connection)
	reader := bufio.NewReader(connection)

	client := &Client{
		id:       bson.NewObjectId().Hex(),
		incoming: make(chan *clientIncomming),
		outgoing: make(chan *chess.Message),
		reader:   reader,
		writer:   writer,
	}

	client.Listen()
	return client
}

type Event int

const (
	GOING Event = 0
	READY Event = iota
	CLOSED
)

type clientIncomming struct {
	clientId string
	event    Event
	content  string
}

type GameHall struct {
	cmdParser *chess.CmdParser
	game      *chess.Game
	clients   []*Client
	joins     chan net.Conn
	incoming  chan *clientIncomming
	outgoing  chan *chess.Message
}

func (this *GameHall) BroadcastTips(data string) {
	for _, client := range this.clients {
		msg := chess.NewMessage(data, false)
		client.outgoing <- msg
	}
	//for i, _ := range this.clients {
	//msg := chess.NewMessage(data, false)
	//this.clients[i].outgoing <- msg
	//}
}

func (this *GameHall) AnnounceGameOver(err error) {
	for _, client := range this.clients {
                fmt.Println(client.Side)
		var text string
		if err == chess.ErrGameOverBlackWin {
			if client.Side == chess.BLACK {
				text = "Game Over! you win :)\n"
			} else {
				text = "Game Over! you lose :)\n"
			}
		}
		if err == chess.ErrGameOverWhiteWin {
			if client.Side == chess.WHITE {
				text = "Game Over! you win :)\n"
			} else {
				text = "Game Over! you lose :)\n"
			}
		}
		msg := chess.NewMessage(text, false)
		client.outgoing <- msg
	}
}

func (this *GameHall) BroadcastBoard() {
	this.broadcastBoard(func(c *Client) bool {
		return c.moveable
	})
}
func (this *GameHall) BroadcastBoardOver() {
	this.broadcastBoard(func(c *Client) bool {
		return false
	})
}
func (this *GameHall) broadcastBoard(f func(*Client) bool) {
	for _, client := range this.clients {
		data := this.game.TextView(client.Side)
		msg := chess.NewMessage(data, f(client))
		client.outgoing <- msg
	}
}

func (this *GameHall) toggleTurn() {
	for _, client := range this.clients {
		client.moveable = !client.moveable
	}
}

func (this *GameHall) removeClient(clientId string) {
	for i, client := range this.clients {
		if client.id == clientId {
			this.clients = append(this.clients[0:i], this.clients[i+1:]...)
			break
		}
	}
}

// TODO turn checking
func (this *GameHall) HandleCmd(ci *clientIncomming) {
	fromClient := this.findClientById(ci.clientId)
	if fromClient == nil {
		panic("invalid id")
	}

	switch ci.event {
	case CLOSED:
		this.removeClient(ci.clientId)
		if len(this.clients) == 1 {
			this.clients[0].moveable = true
			this.clients[0].Side = chess.BLACK

		}
		this.BroadcastTips("opponent(" + ci.clientId + ")quit! waiting for antoher player...\n")
		return
	case READY:
		if len(this.clients) < 2 {
			//fromClient.outgoing <- "welcome to yeer's chess\nwaiting for antoher player..."
			fromClient.moveable = true
			fromClient.Side = chess.BLACK
			fromClient.outgoing <- chess.NewMessage("welcome to yeer's chess", false)
			fromClient.outgoing <- chess.NewMessage("waiting for antoher player...", false)
			//fromClient.outgoing <- chess.NewMessage("welcome to yeer's chess\nwaiting for antoher player...", false)
		} else {
			fromClient.Side = chess.WHITE
			fromClient.outgoing <- chess.NewMessage("welcome to yeer's chess\n", false)
			//time.Sleep(100 * time.Millisecond) //TODO solve the sent togather bug, and remove this time waiting magic
			this.BroadcastTips("found opponent, game start!")
			//time.Sleep(100 * time.Millisecond) //TODO solve the sent togather bug, and remove this time waiting magic
			this.BroadcastBoard()
		}
		return
	case GOING:
	}

	cmd := ci.content
	fromPos, toPos, err := this.cmdParser.Parse(cmd)
	if err != nil {
		fromClient.outgoing <- chess.NewMessage("Cmd error, please retry\n", true)
		return
	}

	if !fromClient.moveable {
		fromClient.outgoing <- chess.NewMessage("Not your turn to move, please wait\n", false)
		return
	}

	err = this.game.Move(fromPos, toPos, fromClient.Side)
	if err != nil {
		if err == chess.ErrGameOverBlackWin || err == chess.ErrGameOverWhiteWin {
			this.BroadcastBoardOver()
			this.AnnounceGameOver(err)
			return
		}

		fromClient.outgoing <- chess.NewMessage("Cmd error, please retry\n", true)
		return
	}
	this.toggleTurn()
	this.BroadcastBoard()
}

func (this *GameHall) Join(connection net.Conn) {
	client := NewClient(connection)
	this.clients = append(this.clients, client)

	// small hall :)
	if len(this.clients) > 2 {
		return
	}

	go func() {
		for {
			ci := <-client.incoming
			this.incoming <- ci
			if ci.event == CLOSED {
				break
			}
		}
	}()
	client.Ready()
}

func (this *GameHall) findClientById(id string) (r *Client) {
	for _, client := range this.clients {
		if client.id == id {
			return client
		}
	}
	return
}

func (this *GameHall) Listen() {
	go func() {
		for {
			select {
			case data := <-this.incoming:
				this.HandleCmd(data)
			case conn := <-this.joins:
				fmt.Println("Join")
				this.Join(conn)
			}
		}
	}()
}

func NewGameHall() *GameHall {
	gameHall := &GameHall{
		clients:   make([]*Client, 0),
		joins:     make(chan net.Conn),
		incoming:  make(chan *clientIncomming),
		outgoing:  make(chan *chess.Message),
		game:      chess.NewGame(6, chess.NewTextFormatter()),
		cmdParser: chess.NewCmdParser(),
	}

	gameHall.Listen()
	return gameHall
}

func main() {
	gameHall := NewGameHall()
	listener, _ := net.Listen("tcp", ":6666")
	for {
		conn, _ := listener.Accept()
		gameHall.joins <- conn
	}
}
