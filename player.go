package main

import (
	"bytes"
	"golang.org/x/net/websocket"
	"io/ioutil"
	"os/exec"
	"path"
	"path/filepath"
)

type ApiReturn struct {
	Action   string `json:"action"`
	Message  string `json:"message"`
	Results  string `json:"results"`
	Song     string `json:"song"`
	Playlist string `json:"playlist"`
}

type MusicPlayer struct {
	Track string `json:"track"`
	Id    int    `json:"id"`
	Dir   string `json:"list"`
	Ws    *websocket.Conn
}

func (player *MusicPlayer) Play(new_track string) {
	player.Stop()
	player.Track = new_track
	Info.Println(player.Track)
	go func(player *MusicPlayer) {
		Info.Printf("Playing %s", player.Track)
		cmd := "/usr/bin/omxplayer"
		// args := []string{ "-o","local", path.Join(player.Dir, player.Track) }
		args := []string{"-o", "local", player.Track}
		_, err := exec.Command(cmd, args...).Output()
		if err != nil {
			Error.Println(err)
		} else {
			Info.Println("Song is finished")
		}
		// player.Next()
	}(player)
	resp := ApiReturn{Message: player.Track, Results: "ok", Action: "play", Song: player.Track}
	websocket.JSON.Send(player.Ws, resp)
}
func (player *MusicPlayer) Stop() {
	player.Track = ""
	Trace.Printf("Stopping music")
	cmd := exec.Command("killall", "omxplayer.bin")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		Error.Println(err)
		if len(stderr.String()) != 0 {
			Error.Println(stderr.String())
		}
	}
	if out.String() != "" {
		Info.Println(out.String())
	}
	resp := ApiReturn{Message: "Silence!!", Results: "ok", Action: "stop", Song: player.Track}
	websocket.JSON.Send(player.Ws, resp)
}
func (player *MusicPlayer) Back() {
	files, _ := ioutil.ReadDir(player.Dir)
	player.Track = "No music files"
	if len(files) != 0 {
		player.Id = modulo((player.Id - 1), len(files))
		// player.Track = files[player.Id].Name()
		player.Track = path.Join(player.Dir, files[player.Id].Name())
		if files[player.Id].IsDir() { //
			player.Back()
		}
	}
	player.Play(player.Track)
}
func (player *MusicPlayer) Next() {
	files, _ := ioutil.ReadDir(player.Dir)
	player.Track = "No music files"
	if len(files) != 0 {
		player.Id = (player.Id + 1) % len(files)
		// player.Track = files[player.Id].Name()
		player.Track = path.Join(player.Dir, files[player.Id].Name())
		if files[player.Id].IsDir() {
			player.Next()
		}
	}
	player.Play(player.Track)
}
func (player *MusicPlayer) Random() string {
	files, _ := ioutil.ReadDir(player.Dir)
	if len(files) != 0 {
		i := randInt(0, len(files))
		if files[i].Name() == player.Track || files[i].IsDir() {
			return player.Random()
		} else {
			player.Track = path.Join(player.Dir, files[i].Name())
			player.Id = i
			return path.Join(player.Dir, files[i].Name())
		}
	} else {
		Warning.Println("No files found")
		return "No music files"
	}
	return "this shouldnt happen"
}
func (player *MusicPlayer) Playlist(directory string) {
	if directory == "" {
		directory = MUSIC_DIR
	}
	player.Dir = directory
	player.Id = 0
	files := getFilesInDirectory(directory)
	folders := getFoldersInDirectory(directory)
	var song_list string
	if filepath.ToSlash(directory) == filepath.ToSlash(MUSIC_DIR) {
		song_list = `<tr>
						<th class="success">Songs</th>
					</tr>`
	} else {
		song_list = `<tr class="back_directory">
						<th class="success" id="` + filepath.Dir(directory) + `"><i class="fa fa-caret-square-o-left"></i> ` + filepath.Base(directory) + `</th>
					</tr>`
	}
	for _, v := range folders {
		song_list += `
					<tr class="playlist">
						<td class="warning" id="` + path.Join(directory, v) + `"><i class="fa fa-caret-square-o-right"></i> ` + v + `</td>
					</tr>`
	}
	for _, v := range files {
		song_list += `
					<tr class="song">
						<td id="` + path.Join(directory, v) + `">` + v + `</td>
					</tr>`
	}
	resp := ApiReturn{Message: "Get playlist", Results: "ok", Action: "playlist", Song: "", Playlist: song_list}
	websocket.JSON.Send(player.Ws, resp)
}
