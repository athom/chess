package main

import (
	"encoding/json"

	"github.com/athom/chess"

	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
)

func handleSignal(conn net.Conn) {
	_, _ = conn.Write([]byte("QUIT\n"))
	os.Exit(1)
}

func main() {
	conn, err := net.Dial("tcp", ":6666")
	checkError(err)

	go func() {
		var bytes [10000]byte
		for {
			var n int
			n, err = conn.Read(bytes[:])
			gs := &chess.GameState{}
			err := json.Unmarshal(bytes[:n], &gs)
			if err != nil {
				panic(err)
			}
			ui := chess.NewConsoleUI(gs)
			ui.Render()

			if gs.MyBoardInfo != nil && gs.MyBoardInfo.Movable {
				fmt.Print(">")
			}
		}

	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, os.Kill)
		<-c
		handleSignal(conn)
		_, _ = conn.Write([]byte("QUIT\n"))
	}()

	// read input
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		_, err = conn.Write([]byte(input))
	}
}
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
}
