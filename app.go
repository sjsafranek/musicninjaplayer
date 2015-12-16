package main

import (
	"os"
	"time"
	"io/ioutil"
	"flag"
	"net/http"
	"path"
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
	// serveSingle("/sitemap.xml", "./sitemap.xml")
	// serveSingle("/favicon.ico", "./static/favicon.ico")
	// // serveSingle("/robots.txt", "./static/robots.txt")
	// serveSingle("/logo.png", "./static/logo.png")

	serveSingle("/favicon.ico", path.Join(STATIC_DIR, "favicon.ico"))
	serveSingle("/logo.png", path.Join(STATIC_DIR, "logo.png"))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/api/v1/play", playMusicHandler)
	http.HandleFunc("/api/v1/stop", stopMusicHandler)
	http.HandleFunc("/api/v1/back", backTrackHandler)
	http.HandleFunc("/api/v1/next", nextTrackHandler)
	
	// Start app
	Info.Printf("Magic happens on port %s...", *port)
	err := http.ListenAndServe(":" + *port, nil)
	if err != nil {
		panic(err)
	}

}

