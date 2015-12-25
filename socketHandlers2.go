package main

import (
	"net/http"
	"html/template"
	"golang.org/x/net/websocket"
	"os/exec"
	"io/ioutil"
	"path"
	"bytes"
)

/*
type SocketMessage struct {
	Action  string    `json:"action"`
	Song    string    `json:"song"`
}

func clientErrorHandler(w http.ResponseWriter, r *http.Request) {
	Error.Println("%s client had an error", r.RemoteAddr)
	http.ServeFile(w, r, "./static/error.png")
}
*/


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
		player.Next()
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





/*

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



*/







func webSocketHandler2(ws *websocket.Conn) {
	var data SocketMessage
	player := MusicPlayer{ Ws: ws, Id: 0 }
	for {
		if err := websocket.JSON.Receive(ws, &data); err != nil {
			stopMusic()
			Error.Println(err)
			return
		} else {
			switch data.Action {
				case "play":
					current_song_name = data.Song
					if current_song_name == "" {
						current_song_name = player.Random()
					}
					player.Play(current_song_name)	
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




func socketClientHandler2(w http.ResponseWriter, r *http.Request) {
	files := getMusicFiles()
	song_list := ""
	for _, v := range files {
		song_list += `
					<tr class="song">
						<td id="` + v + `">` + v + `</td>
					</tr>`
	}

	Info.Printf("%s something is happening", r.RemoteAddr)
	tmpl := template.New("music controller template")
	page :=`
<head>

	<title>Music Ninja Player</title>

	<meta charset="utf-8">
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	<meta name="viewport" content="width=device-width, initial-scale=1">

	<link rel="stylesheet" href="/static/css/font-awesome.min.css">
	<script src="/static/js/jquery.min.js"></script>

	<!-- Latest compiled and minified CSS -->
	<link rel="stylesheet" href="/static/css/bootstrap.min.css">
	<!-- Optional theme -->
	<link rel="stylesheet" href="/static/css/bootstrap-theme.min.css">

	<style>
		
		body { padding-top: 50px; }
		#error {
			background-color: white;
			text-align: center;
			font-weight: bold;
			z-index: 15;
			position: absolute;
			bottom: 0;
			top: 0;
			right: 0;
			left: 0;
		}

	</style>

</head>
<body>

	<div id="error" style="display:none;">
		<img src="/error" alt="error" style="margin-top:5px;">
		<div id="error_message"></div>
	</div>

	<nav class="navbar navbar-inverse navbar-fixed-top">
		<div class="container">
			<div class="navbar-header">
				<button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar" aria-expanded="false" aria-controls="navbar">
					<span class="sr-only">Toggle navigation</span>
					<span class="icon-bar"></span>
					<span class="icon-bar"></span>
					<span class="icon-bar"></span>
				</button>
				<a class="navbar-brand" href="#">Music Ninja Player</a>
			</div>
			<div id="navbar" class="collapse navbar-collapse">
				<ul class="nav navbar-nav">
					<li class="active"><a href="#">Home</a></li>
					<li><a href="#about">About</a></li>
					<li><a href="#contact">Contact</a></li>
				</ul>
			</div><!--/.nav-collapse -->
		</div>
	</nav><!-- /.nav -->

	<div class="container">
		<div class="row">
			<div class="col-md-2">
				<img src="/static/logo.png" alt="logo" style="height:170px; width:170px;">
			</div>
			<div class="col-md-4">
				<h3>Music Ninja Player</h3>
				<div>
					<label><b>Current: </b></label><span id="current"></span>
				</div>
				<div>
					<button type='button' title="back track" id="back" class="btn btn-primary"><i id="back" class="fa fa-backward"></i></button>
					<button type='button' title="random track" id="play" class="btn btn-primary"><i id="play" class="fa fa-play"></i></button>
					<button type='button' title="stop music" id="stop" class="btn btn-primary"><i id="stop" class="fa fa-stop"></i></button>
					<button type='button' title="next track" id="next" class="btn btn-primary"><i id="next" class="fa fa-forward"></i></button>
				</div>
			</div>
			<div class="col-md-6">
				<!-- <h3>Playlist</h3> -->
				<h3></h3>
				<table class="table table-striped table-bordered table-hover">
					<tr>
						<th class="success">Songs</th>
					</tr>
					` + song_list + `
				</table>
			</div>
		</div> <!-- /.row -->
	</div><!-- /.container -->

	<!-- Latest compiled and minified JavaScript -->
	<script src="/static/js/bootstrap.min.js"></script>
	<script>

		// Setup Websocket
		function getWebSocket() {
			console.log("Opening websocket");
			var url = "ws://" +window.location.host + "/ws2";

			ws = new WebSocket(url);
			ws.onopen = function(e) { 
				console.log("Websocket is open");
			};
			ws.onmessage = function(e) {
				var data = JSON.parse(e.data);
				console.log("Data recieved:",data);
				$("#current")[0].innerText = " " + data.song;
			};
			ws.onclose = function(e) { 
				console.log("Websocket is closed"); 
				$("#error").css("display","block");
				$("#error_message").text("Connection error");
			}
			ws.onerror = function(e) { console.log(e); }
			return ws;
		}

		var ws = getWebSocket();


		function playSong(event) {
			console.log("Sending WebSocket request")
			var msg = {
				action: event.target.id,
				song: null
			};
			var payload = JSON.stringify(msg);
			try {
				ws.send(payload);
			}
			catch(err) {
				console.log(err);
				alert(err);
				window.location = "/error";
			}
		}
		$("button").on("click", playSong);

		function chooseSong(event) {
			console.log("Sending WebSocket request")
			var msg = {
				action: "play",
				song: event.target.id
			}
			var payload = JSON.stringify(msg);
			try {
				ws.send(payload);
			}
			catch(err) {
				console.log(err);
				alert(err);
				window.location = "/error";
			}
		}
		$("tr").on("click",chooseSong)


	</script>

</body>
	`
	tmpl, _ = tmpl.Parse(page)
	tmpl.Execute(w, nil) 
}

