package main

import (
	"log"
	"net/http"
	"runtime/debug"

	"github.com/athom/chess"
	"github.com/bom-d-van/goutil/printutils"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/cors"
	"github.com/martini-contrib/gzip"
	"github.com/martini-contrib/render"
	"golang.org/x/net/websocket"
	//"github.com/gorilla/websocket"
)

var gameHall = chess.NewHall()

func main() {

	m := martini.Classic()
	m.Use(gzip.All())
	m.Use(cors.Allow(&cors.Options{
		//AllowOrigins: []string{"http://*"},
		AllowOrigins:     []string{"http://localhost:6001"},
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
	//m.Use(render.Renderer())

	// routeers
	m.Get("/ws", websocket.Handler(BuildConnection).ServeHTTP)
	//m.Get("/ws", BuildConnection)

	//m.Get("/", func() string {
	//return "Hello Welcome to Yeer Chess"
	//})
	m.Get("/", func(r render.Render) {
		r.HTML(200, "index", nil)
	})

	http.ListenAndServe(":7200", m)
	m.Run()
}

func BuildConnection(ws *websocket.Conn) {
	//func BuildConnection(params martini.Params, w http.ResponseWriter, r *http.Request, ws *websocket.Conn) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("********** WebSocket Error: %+v ***********\n", err)
			log.Println(debug.Stack())
		}
	}()

	log.Println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	printutils.PrettyPrint(ws)
	gameHall.Join(ws)
}
