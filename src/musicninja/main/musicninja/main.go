package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/websocket"
	"musicninja/app"
	"net/http"
	"path"
)

var (
	port int
)

func init() {
	// Parse command line for arguements
	flag.IntVar(&port, "p", 8080, "server port")
	flag.StringVar(&app.Apikey, "k", "test", "apikey")
	flag.Parse()
}

// Main
func main() {

	// Mandatory root-based resources
	app.ServeSingle("/favicon.ico", path.Join("/static", "favicon.ico"))
	// serveSingle("/sitemap.xml", "./sitemap.xml")
	// serveSingle("/robots.txt", "./static/robots.txt")

	// Static Files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Main Routes
	http.HandleFunc("/music", app.ClientHandler)
	http.HandleFunc("/ping", app.PingHandler)
	http.HandleFunc("/error", app.ErrorHandler)

	// Web Socket
	http.Handle("/ws", websocket.Handler(app.WebSocketHandler))

	bind := fmt.Sprintf(":%v", port)
	// Start app
	app.Info.Printf("Magic happens on port %v...", port)
	err := http.ListenAndServe(bind, nil)
	if err != nil {
		panic(err)
	}

}
