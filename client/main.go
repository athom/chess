package main

import (
	"bytes"
	"encoding/json"
	"io"

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
	if len(os.Args) != 2 {
		//fmt.Println("Usage: ", os.Args[0], "host")
		//os.Exit(1)
	}
	//host := os.Args[1]
	//host := "localhost"
	conn, err := net.Dial("tcp", ":6666")
	checkError(err)

	go func() {
		var bytes [10000]byte
		for {
			var n int
			n, err = conn.Read(bytes[:])
			//if false {
			msg := &chess.Message{}
			err := json.Unmarshal(bytes[:n], &msg)
			if err != nil {
				//handleSignal(conn)
				panic(err)
			}
			//}

			fmt.Println("\n" + msg.UI)
			if msg.MyTurn {
				fmt.Print(">")
			}
			//fmt.Println(string(bytes[0:n]))
		}

	}()

	go func() {
		//_, err = conn.Write([]byte("HEAD"))
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, os.Kill)
		// Block until a signal is received.
		<-c
		handleSignal(conn)
		_, _ = conn.Write([]byte("QUIT\n"))
	}()

	reader := bufio.NewReader(os.Stdin)
	for {
		//var output [1003]byte
		//_, err = conn.Read(output[0:])
		//fmt.Println("received: ", string(output[0:]))
		input, _ := reader.ReadString('\n')
		_, err = conn.Write([]byte(input))
		//fmt.Println(input)

	}
}
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
}

func readFully(conn net.Conn) ([]byte, error) {
	//defer conn.Close()
	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}
