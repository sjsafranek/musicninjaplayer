package main

import (
	"os/exec"
	"io/ioutil"
	"path"
)

var current_song_name string
var current_song_id = 0 

func getMusicFiles() []string {
	results := []string{}
	files, _ := ioutil.ReadDir(MUSIC_DIR)
	for i := 0; i < len(files); i++ {
		results = append(results,files[i].Name())
	}
	return results
}

func randomSong() string {
	files, _ := ioutil.ReadDir(MUSIC_DIR)
	if len(files) != 0 {
		i := randInt(0,len(files))
		if files[i].Name() == current_song_name {
			return randomSong()
		} else {
			current_song_name = files[i].Name()
			current_song_id = i
			return files[i].Name()
		}
	} else {
		Warning.Println("No files found")
		return "No music files"
	}
	return "this shouldnt happen"
}

func backSong() {
	files, _ := ioutil.ReadDir(MUSIC_DIR)
	if len(files) != 0 {
		current_song_id = modulo((current_song_id - 1), len(files))
		current_song_name = files[current_song_id].Name()
	} else {
		current_song_name = "No music files"
	}
}

func nextSong() {
	files, _ := ioutil.ReadDir(MUSIC_DIR)
	if len(files) != 0 {
		current_song_id = (current_song_id + 1) % len(files)
		current_song_name = files[current_song_id].Name()
	} else {
		current_song_name = "No music files"
	}
}

func playMusic(song string) {
	stopMusic()
	cmd := "play"
	args := []string{ path.Join(MUSIC_DIR, song) }
	_, err := exec.Command(cmd, args...).Output()
	if err != nil {
		Warning.Println(err)
		// raspberry pi: raspbian
		// cmd := "omxpalyer -o local"
		args := []string{ path.Join(MUSIC_DIR, song) }
		_, err2 := exec.Command("omxpalyer -o local", args...).Output()
		if err2 == nil {
			Info.Printf("Playing %s", song)
		} else {
			Warning.Println(err2)
		}
	} else {
		Info.Printf("Playing %s", song)
	}
}

func stopMusic() {
	Info.Printf("Stopping music")
	cmd := "killall"
	args := []string{"play"}
	exec.Command(cmd, args...).Output()
	current_song_name = ""
}
