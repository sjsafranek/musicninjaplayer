package main

import (
	"net/http"
	"encoding/json"
	"time"
)


// Serves static files
// http://stackoverflow.com/questions/14086063/serve-homepage-and-static-content-from-root
func serveSingle(pattern string, filename string) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		Info.Printf("%s something is happening...", r.RemoteAddr)
		http.ServeFile(w, r, filename)
	})
}

// Serves error page
func errorHandler(w http.ResponseWriter, r *http.Request, err error) {
	Error.Printf("%s %v}", r.RemoteAddr, err)
	http.ServeFile(w, r, "./static/error.png")
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	type Ping struct {
		Message	   string	  `json:"message"`
		Registered	time.Time   `json:"registered"`
		Runtime	   float64	 `json:"runtime"`
	}
	Info.Printf("%s something is happening...", r.RemoteAddr)
	resp := Ping{Message: "Pong", Registered: start_time, Runtime: time.Since(start_time).Seconds()}
	json.NewEncoder(w).Encode(resp)
}

