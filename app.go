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

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/test", testHandler)
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/api/v1/play", playMusicHandler)
	http.HandleFunc("/api/v1/stop", stopMusicHandler)
	http.HandleFunc("/api/v1/back", backTrackHandler)
	http.HandleFunc("/api/v1/next", nextTrackHandler)

	serveSingle("/favicon.ico", path.Join(STATIC_DIR, "favicon.ico"))
	serveSingle("/static/logo.png", path.Join(STATIC_DIR, "logo.png"))

	serveSingle("/js/jquery.min.js", path.Join(STATIC_DIR, "jquery.min.js"))

	serveSingle("/js/bootstrap.min.js", path.Join(STATIC_DIR, "bootstrap.min.js"))
	serveSingle("/css/bootstrap.min.css", path.Join(STATIC_DIR, "bootstrap.min.css"))
	serveSingle("/css/bootstrap-theme.min.css", path.Join(STATIC_DIR, "bootstrap-theme.min.css"))

	serveSingle("/4.5.0/css/font-awesome.min.css", path.Join(STATIC_DIR, "font-awesome.min.css"))
	serveSingle("/4.5.0/fonts/fontawesome-webfont.woff2?v=4.5.0", path.Join(STATIC_DIR, "fontawesome-webfont.woff2"))
	serveSingle("/4.5.0/fonts/fontawesome-webfont.ttf?v=4.5.0", path.Join(STATIC_DIR, "fontawesome-webfont.ttf"))

	// Start app
	Info.Printf("Magic happens on port %s...", *port)
	err := http.ListenAndServe(":" + *port, nil)
	if err != nil {
		panic(err)
	}

}

