package main

import (
	"golang.org/x/net/websocket"
	"os/exec"
	"io/ioutil"
	"path"
	"bytes"
)

type ApiReturn struct {
	Action	 string	  `json:"action"`
	Message	string	  `json:"message"`
	Results	string	  `json:"results"`
	Song	   string	  `json:"song"`
}

type SocketMessage struct {
	Action  string    `json:"action"`
	Song    string    `json:"song"`
}

type MusicPlayer struct {
	Track  string  `json:"track"`
	Id     int  `json:"id"`
	Ws     *websocket.Conn
}
func (player *MusicPlayer) Play(new_track string) {
	player.Stop()
	player.Track = new_track
	go func(player *MusicPlayer) {
		cmd := "/usr/bin/omxplayer"
		args := []string{ "-o","local", path.Join(MUSIC_DIR, player.Track) }
		_, err := exec.Command(cmd, args...).Output()
		if err != nil {
			Warning.Println(err)
		} else {
			Info.Printf("Playing %s", player.Track)
		}
		Info.Println("Song is finished")
		// player.Next()
	}(player)
	resp := ApiReturn{ Message:player.Track, Results:"ok", Action:"play", Song:player.Track }
	websocket.JSON.Send(player.Ws, resp)
}
func (player *MusicPlayer) Stop() {
	player.Track = ""
	Info.Printf("Stopping music")
	cmd := exec.Command("killall", "omxplayer.bin")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		Error.Println(err)
		Error.Println(stderr.String())
	}
	Info.Println(out.String())
	resp := ApiReturn{ Message:"Silence!!", Results:"ok", Action:"stop", Song:player.Track }
	websocket.JSON.Send(player.Ws, resp)
}
func (player *MusicPlayer) Back() {
	files, _ := ioutil.ReadDir(MUSIC_DIR)
	player.Track = "No music files"
	if len(files) != 0 {
		player.Id = modulo((player.Id - 1), len(files))
		player.Track = files[player.Id].Name()
	}
	player.Play(player.Track)
}
func (player *MusicPlayer) Next() {
	files, _ := ioutil.ReadDir(MUSIC_DIR)
	player.Track = "No music files"
	if len(files) != 0 {
		player.Id = (player.Id + 1) % len(files)
		player.Track = files[player.Id].Name()
	}
	player.Play(player.Track)
}
func (player *MusicPlayer) Random() string {
	files, _ := ioutil.ReadDir(MUSIC_DIR)
	if len(files) != 0 {
		i := randInt(0,len(files))
		if files[i].Name() == player.Track {
			return player.Random()
		} else {
			player.Track = files[i].Name()
			player.Id = i
			return files[i].Name()
		}
	} else {
		Warning.Println("No files found")
		return "No music files"
	}
	return "this shouldnt happen"
}


func getMusicFiles() []string {
	results := []string{}
	files, _ := ioutil.ReadDir(MUSIC_DIR)
	for i := 0; i < len(files); i++ {
		results = append(results,files[i].Name())
	}
	return results
}
