package main

import (
	"log"
	"runtime/debug"
	"strings"

	"github.com/athom/chess"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/cors"
	"github.com/martini-contrib/gzip"
	"golang.org/x/net/websocket"
)

var gameHall = chess.NewHall()

func main() {

	m := martini.Classic()
	m.Use(gzip.All())
	m.Use(cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"PUT", "GET", "POST", "DELETE", "OPTIONS"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// NOTE dont check origin to make the client websocket possible
	wsConfig, err := websocket.NewConfig("ws://localhost:3001", "http://*")
	if err != nil {
		panic(err)
	}

	m.Get("/ws/:roomslug", websocket.Server{
		Config:  *wsConfig,
		Handler: BuildConnection,
	}.ServeHTTP)

	m.Get("/", func() string {
		return "Welcome to Yeer's chess, it works!"
	})

	m.Run()
}

func BuildConnection(ws *websocket.Conn) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("********** WebSocket Error: %+v ***********\n", err)
			log.Println(debug.Stack())
		}
	}()
	defer ws.Close()

	array := strings.Split(ws.Request().URL.Path, `/ws/`)
	if len(array) < 2 {
		panic("room slug not found")
	}

	slug := array[1]

	mb := chess.NewWebsocketMailBox(ws)
	//gameHall.Join(mb)
	gameHall.JoinRoom(mb, slug)
	mb.Run()
}
