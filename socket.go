package main

import (
	"golang.org/x/net/websocket"
)

func webSocketHandler(ws *websocket.Conn) {
	var data SocketMessage
	player := MusicPlayer{ Ws: ws, Id: 0 }
	for {
		if err := websocket.JSON.Receive(player.Ws, &data); err != nil {
			player.Stop()
			Error.Println(err)
			return
		} else {
			switch data.Action {
				case "play":
					player.Track = data.Song
					if player.Track == "" {
						player.Track = player.Random()
					}
					player.Play(player.Track)	
				case "back":
					player.Back()
				case "next":
					player.Next()
				default: // "stop"
					player.Stop()
			}
		}
		Info.Printf("Received: %s %s", data.Action, data.Song)
	}
}




