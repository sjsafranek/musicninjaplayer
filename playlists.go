package main

import (
	"encoding/json"
	"io/ioutil"
)

var Playlists = getPlaylists()

type Playlist struct {
	Dir        string      `json"dir"`
	Name       string      `json:"name"`
	Songs      []string    `json:"songs"`
}

type MusicData struct {
	Playlists  []Playlist  `json:"playlists"`
}

func getPlaylists() MusicData {
	content, err := ioutil.ReadFile("music.json")
	if err!=nil{
		Error.Print("Error:",err)
	}
	var music MusicData
	err=json.Unmarshal(content, &music)
	if err!=nil{
		Error.Print("Error:",err)
	}
	return music
}
