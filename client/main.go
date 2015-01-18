package main

import (
	"encoding/json"

	"github.com/athom/chess"
	"github.com/athom/easysignal"

	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":6666")
	checkError(err)

	easysignal.OnProcessKilled(func() {
		_, _ = conn.Write([]byte("QUIT\n"))
		os.Exit(1)
	})

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
