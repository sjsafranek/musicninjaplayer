package main

import (
	"os/exec"
	"io/ioutil"
	"path"
)

var current_song_name string
var current_song_id = 0 

// func getFiles() []os.FileInfo {
// 	files, _ := ioutil.ReadDir(MUSIC_DIR)
// 	return files
// }

func randomSong() string {
	// files, _ := ioutil.ReadDir("./music")
	files, _ := ioutil.ReadDir(MUSIC_DIR)
	if len(files) != 0 {
		i := randInt(0,len(files))
		if files[i].Name() == current_song_name || files[i].Name() == ".gitignore" {
			return randomSong()
		} else {
			current_song_name = files[i].Name()
			current_song_id = i
			return files[i].Name()
		}
	} else {
		Error.Fatal("No files found")
		return "No music files"
	}
}

func backSong() {
	// files, _ := ioutil.ReadDir("./music")
	files, _ := ioutil.ReadDir(MUSIC_DIR)
	
	current_song_id = modulo((current_song_id - 1), len(files))
	if files[current_song_id].Name() != ".gitignore" {
		current_song_name = files[current_song_id].Name()
	} else {
		backSong()
	}
	
}

func nextSong() {
	// files, _ := ioutil.ReadDir("./music")
	files, _ := ioutil.ReadDir(MUSIC_DIR)
	current_song_id = (current_song_id + 1) % len(files)
	if files[current_song_id].Name() != ".gitignore" {
		current_song_name = files[current_song_id].Name()
	} else {
		nextSong()
	}
}

// func playMusic(song ...string) {
//	 stopMusic()
//	 cmd := "play"
//	 args := []string{"music/" + song[0]}
//	 _, err := exec.Command(cmd, args...).Output()
//	 if err != nil {
//		 Error.Println(err)
//	 } else {
//		 Info.Printf("Playing %s", song[0])
//	 }
// }

func playMusic(song string) {
	stopMusic()
	cmd := "play"
	// args := []string{"music/" + song}
	args := []string{ path.Join(MUSIC_DIR, song) }
	_, err := exec.Command(cmd, args...).Output()
	if err != nil {
		Error.Println(err)
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
