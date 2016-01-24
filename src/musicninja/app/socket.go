package app

import (
	"os"
	// "path"
	"golang.org/x/net/websocket"
)

type SocketMessage struct {
	Action string `json:"action"`
	Song   string `json:"song"`
}

// Create a websocket and keeps it open
func WebSocketHandler(ws *websocket.Conn) {
	player := MusicPlayer{Ws: ws, Id: 0, Dir: MUSIC_DIR}
	for {
		var data SocketMessage
		if err := websocket.JSON.Receive(player.Ws, &data); err != nil {
			player.Stop()
			Error.Println(err)
			return
		} else {
			switch data.Action {
			case "play":
				// Why is it receiving a directory path? Client doesn't appear to send it
				fileInfo, _ := os.Stat(data.Song)
				if data.Song == "" || fileInfo.IsDir() {
					player.Track = player.Random()
				} else {
					player.Track = data.Song
				}
				player.Play(player.Track)
			case "back":
				player.Back()
			case "next":
				player.Next()
			case "playlist":
				player.Playlist(data.Song)
			default: // "stop"
				player.Stop()
			}
		}
		Info.Printf("Received: %s %s", data.Action, data.Song)
	}
}
