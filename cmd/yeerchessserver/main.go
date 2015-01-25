package main

import (
	"net"

	"github.com/athom/chess"
)

func main() {
	gameHall := chess.NewHall()
	listener, _ := net.Listen("tcp", ":6666")
	for {
		conn, _ := listener.Accept()
		gameHall.Joins <- chess.NewConsoleMailBox(conn)
	}
}
