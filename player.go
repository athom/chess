package chess

import (
	"bufio"
	"fmt"
	"net"

	"gopkg.in/mgo.v2/bson"
)

type Player struct {
	id        string
	incoming  chan string
	outgoing  chan string
	reader    *bufio.Reader
	writer    *bufio.Writer
	firstHand bool
}

func (this *Player) Read() {
	for {
		line, err := this.reader.ReadString('\n')
		if err != nil {
			continue
		}
		this.incoming <- line
	}
}

func (this *Player) Write() {
	for data := range this.outgoing {
		//this.writer.WriteString(data)
		this.writer.Write(NewMessage(data, this.firstHand).ToJson())
		this.writer.Flush()
		fmt.Println("out: ", data)
	}
}

func (this *Player) Listen() {
	go this.Read()
	go this.Write()
}

func NewPlayer(connection net.Conn) *Player {
	writer := bufio.NewWriter(connection)
	reader := bufio.NewReader(connection)

	this := &Player{
		id:       bson.NewObjectId().Hex(),
		incoming: make(chan string),
		outgoing: make(chan string),
		reader:   reader,
		writer:   writer,
	}

	this.Listen()
	return this
}
