package main

import (
	"net/http"
	"html/template"
)


func testHandler(w http.ResponseWriter, r *http.Request) {
	files := getMusicFiles()
	song_list := ""
	for _, v := range files {
		song_list += `<tr class="song"><td id="` + v + `">` + v + `</td></tr>`
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

		/*  Table Style */
		table {
			border-collapse: collapse;
			width: 100%;
		}
		th, td {
			padding: 15px;
			text-align: left;
		}
		tr:nth-child(even) { background-color: #f2f2f2; }
		table, th, td { border: 1px solid #ddd; }
		th {
			background-color: #4CAF50;
			color: white;
		}
		tr:hover {
			/* background-color: #f5f5f5	*/
			background-color: #FF4040;
			border-color: #FF4040;
			box-shadow: inset 0 1px 1px rgba(0, 0, 0, 0.075), 0 0 8px rgba(255, 0, 0, 0.6);
		}

	</style>

</head>
<body>

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
					<button type='button' id="back"><i id="back" class="fa fa-backward"></i></button>
					<button type='button' id="play"><i id="play" class="fa fa-play"></i></button>
					<button type='button' id="stop"><i id="stop" class="fa fa-stop"></i></button>
					<button type='button' id="next"><i id="next" class="fa fa-forward"></i></button>
				</div>
			</div>
			<div class="col-md-6">
				<!-- <h3>Playlist</h3> -->
				<h3></h3>
				<table>
					<tr>
						<th>
							Songs
						</th>
					</tr>
					` + song_list + `
				</table>
			</div>
		</div> <!-- /.row -->
	</div><!-- /.container -->

	<!-- Latest compiled and minified JavaScript -->
	<script src="/static/js/bootstrap.min.js"></script>
	<script>

		function playSong(event) {
			$.get( "api/v1/" + event.target.id, function( data ) {
				data = $.parseJSON(data);
				console.log(data);
				$("#current")[0].innerText = " " + data.song;;
			});
		}
		$("button").on("click", playSong);

		function chooseSong(event) {
			$.ajax({
				url: "api/v1/play",
				data: "song=" + event.target.id,
				success: function( data ) {
					// data = $.parseJSON(data);
					console.log(data);
					$("#current")[0].innerText = " " + data.song;
				},
				dataType: "json"
			});
		}
		$("tr").on("click",chooseSong)

	</script>

</body>
	`
	tmpl, _ = tmpl.Parse(page)
	tmpl.Execute(w, nil) 
}

