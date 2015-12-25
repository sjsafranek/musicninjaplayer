package main

import (
	"net/http"
	"html/template"
	"golang.org/x/net/websocket"
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
	Id     string  `json:"id"`
	Ws     *websocket.Conn
}
func (player *MusicPlayer) Play(new_track string) {
	player.Track = new_track
	go func() {
		playMusic(player.Track)
		Info.Println("Song is finished")
	}()
	resp := ApiReturn{ Message:player.Track, Results:"ok", Action:"play", Song:player.Track }
	websocket.JSON.Send(player.Ws, resp)
}
func (player *MusicPlayer) Stop() {
	player.Track = ""
	stopMusic()
	resp := ApiReturn{ Message:"Silence!!", Results:"ok", Action:"stop", Song:player.Track }
	websocket.JSON.Send(player.Ws, resp)
}
func (player *MusicPlayer) Back() {
	backSong()
	player.Play(current_song_name)
}
func (player *MusicPlayer) Next() {
	nextSong()
	player.Play(current_song_name)
}



func webSocketHandler2(ws *websocket.Conn) {
	var data SocketMessage
	player := MusicPlayer{ Ws: ws }
	for {
		if err := websocket.JSON.Receive(ws, &data); err != nil {
			stopMusic()
			Error.Println(err)
			return
		} else {
			switch data.Action {
				case "play":
					if data.Song != "" {
						current_song_name = data.Song
					} else {
						current_song_name = randomSong()
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

