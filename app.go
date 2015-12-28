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
var APIKEY string

// Main
func main() {

	// Setup log
	log_init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	dir_init()

	// Parse command line for arguements
	port := flag.String("port", PORT, "server port")
	apikey := flag.String("apikey", "1234", "apikey")
	flag.Parse()

	APIKEY = *apikey

	// Mandatory root-based resources
	serveSingle("/favicon.ico", path.Join(STATIC_DIR, "favicon.ico"))
	// serveSingle("/sitemap.xml", "./sitemap.xml")
	// serveSingle("/robots.txt", "./static/robots.txt")

	// Static Files
	fs := http.FileServer(http.Dir(STATIC_DIR))
	http.Handle("/static/",http.StripPrefix("/static/",fs))
	
	// Main Routes
	http.HandleFunc("/music", clientHandler)
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/error", errorHandler)

	// Web Socket
	http.Handle("/ws", websocket.Handler(webSocketHandler))

	// Start app
	Info.Printf("Magic happens on port %s...", *port)
	err := http.ListenAndServe(":" + *port, nil)
	if err != nil {
		panic(err)
	}

}

