package main

import (
	"log"
	"net/http"
	"runtime/debug"

	"github.com/athom/chess"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/cors"
	"github.com/martini-contrib/gzip"
	"github.com/martini-contrib/render"
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

	m.Use(martini.Static("../webclient/"))

	m.Use(render.Renderer(render.Options{
		Directory: "../webclient/", // Specify what path to load the templates from.
		//Layout:          "layout",                       // Specify a layout template. Layouts can call {{ yield }} to render the current template.
		Extensions: []string{".tmpl", ".html"}, // Specify extensions to load for templates.
		//Funcs:           []template.FuncMap{AppHelpers}, // Specify helper function maps for templates to access.
		//Delims:          render.Delims{"{[{", "}]}"},    // Sets delimiters to the specified strings.
		//Charset:         "UTF-8",                        // Sets encoding for json and html content-types. Default is "UTF-8".
		//IndentJSON:      true,                           // Output human readable JSON
		//IndentXML:       true,                           // Output human readable XML
		//HTMLContentType: "application/xhtml+xml",        // Output XHTML content type instead of default "text/html"
	}))

	// routeers
	//m.Get("/ws", websocket.Handler(BuildConnection).ServeHTTP)
	wsConfig, err := websocket.NewConfig("ws://localhost:7201", "http://*")
	if err != nil {
		panic(err)
	}

	m.Get("/ws", websocket.Server{
		Config:  *wsConfig,
		Handler: BuildConnection,
	}.ServeHTTP)

	m.Get("/", func(r render.Render) {
		r.HTML(200, "index", nil)
	})

	http.ListenAndServe(":7200", m)
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

	log.Println("ws connected ..")
	mb := chess.NewWebsocketMailBox(ws)
	gameHall.Join(mb)
	mb.Run()
}
