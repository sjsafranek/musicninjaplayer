package main

import (
	"os"
	"time"
	"io/ioutil"
	"flag"
	"net/http"
)

// Defaults
var start_time = time.Now()
const PORT string = "8080"
const CONF string = "music.json"

func main() {

	// Parse command line for arguements
	port := flag.String("port", PORT, "server port")
	// conf := flag.String("conf", CONF, "config json file")
	flag.Parse()

	// Setup log
	log_init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	// Mandatory root-based resources
	serveSingle("/sitemap.xml", "./sitemap.xml")
	serveSingle("/favicon.ico", "./static/favicon.ico")
	serveSingle("/robots.txt", "./static/robots.txt")
	serveSingle("/logo.png", "./static/logo.png")
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/api/v1/play", playMusicHandler)
	http.HandleFunc("/api/v1/stop", stopMusicHandler)
	http.HandleFunc("/api/v1/back", backTrackHandler)
	http.HandleFunc("/api/v1/next", nextTrackHandler)
	// http.HandleFunc("/api/v1/list", apiListMusic)
	http.HandleFunc("/api/v1/playlists", playlistsHandler)
	
	// Start app
	Info.Printf("Magic happens on port %s...", *port)
	err := http.ListenAndServe(":" + *port, nil)
	if err != nil {
		panic(err)
	}

}

