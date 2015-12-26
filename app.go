package main

import (
	"os"
	"time"
	"io/ioutil"
	"flag"
	"net/http"
	"path"
	"golang.org/x/net/websocket"
)

// Defaults
var start_time = time.Now()
const PORT string = "8080"
const CONF string = "music.json"


func main() {
	// Setup log
	log_init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	dir_init()

	// Parse command line for arguements
	port := flag.String("port", PORT, "server port")
	flag.Parse()

	// Mandatory root-based resources
	serveSingle("/favicon.ico", path.Join(STATIC_DIR, "favicon.ico"))
	// serveSingle("/sitemap.xml", "./sitemap.xml")
	// // serveSingle("/robots.txt", "./static/robots.txt")

	// Static Files
	fs := http.FileServer(http.Dir(STATIC_DIR))
	http.Handle("/static/",http.StripPrefix("/static/",fs))
	
	// Main Routes
	http.HandleFunc("/socket", socketClientHandler)
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/error", clientErrorHandler)

	//  Api Routes
	// http.HandleFunc("/api/v1/play", playMusicHandler)
	// http.HandleFunc("/api/v1/stop", stopMusicHandler)
	// http.HandleFunc("/api/v1/back", backTrackHandler)
	// http.HandleFunc("/api/v1/next", nextTrackHandler)

	// Web Socket
	http.Handle("/ws", websocket.Handler(webSocketHandler))


	// Start app
	Info.Printf("Magic happens on port %s...", *port)
	err := http.ListenAndServe(":" + *port, nil)
	if err != nil {
		panic(err)
	}

}

