package main

import (
	"os"
	"path"
	"path/filepath"
	"strings"
	"runtime"
	"net/http"
	"io"
)

var BASE_DIR string
var NINJA_DIR string
var MUSIC_DIR  string
var STATIC_DIR string

func setup_music_dir() {
	if runtime.GOOS == "windows" || runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		BASE_DIR = path.Join(homeDir(),"Music")
		BASE_DIR = strings.Replace(BASE_DIR,"\\","/",-1)
		_, err := os.Stat(BASE_DIR)
		if err != nil {
			Info.Printf("Creating %s", BASE_DIR)
			os.Mkdir(BASE_DIR, os.ModeSticky | 0755)
		}
	} else {
		Error.Fatal("This OS is not supported!")
	}
}

func ninja_dir() {
		if runtime.GOOS == "windows" || runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		_, err := os.Stat(path.Join(BASE_DIR,"Ninja"))
		if err != nil {
			Info.Printf("Creating %s", path.Join(BASE_DIR,"Ninja"))
			os.Mkdir(path.Join(BASE_DIR,"Ninja"), os.ModeSticky | 0755)
		}
		NINJA_DIR = path.Join(BASE_DIR,"Ninja")
	} else {
		Error.Fatal("This OS is not supported!")
	}
}

func ninja_music_dir() {
		if runtime.GOOS == "windows" || runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		_, err := os.Stat(path.Join(BASE_DIR,"Ninja", "music"))
		if err != nil {
			Info.Printf("Creating %s", path.Join(BASE_DIR,"Ninja", "music"))
			os.Mkdir(path.Join(BASE_DIR,"Ninja", "music"), os.ModeSticky | 0755)
		}
		MUSIC_DIR = path.Join(BASE_DIR,"Ninja", "music")
	} else {
		Error.Fatal("This OS is not supported!")
	}
}

func static_dir() {
	if runtime.GOOS == "windows" || runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		_, err := os.Stat(path.Join(BASE_DIR,"Ninja", "static"))
		if err != nil {
			Info.Printf("Creating %s", path.Join(BASE_DIR,"Ninja", "static"))
			os.Mkdir(path.Join(BASE_DIR,"Ninja", "static"), os.ModeSticky | 0755)
		}
		STATIC_DIR = path.Join(BASE_DIR,"Ninja", "static")
		css_dir()
		js_dir()
		fonts_dir()
	} else {
		Error.Fatal("This OS is not supported!")
	}
}



func css_dir() {
	if runtime.GOOS == "windows" || runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		_, err := os.Stat(path.Join(STATIC_DIR, "css"))
		if err != nil {
			Info.Printf("Creating %s", path.Join(STATIC_DIR,"css"))
			os.Mkdir(path.Join(STATIC_DIR,"css"), os.ModeSticky | 0755)
		}
	} else {
		Error.Fatal("This OS is not supported!")
	}
}

func js_dir() {
	if runtime.GOOS == "windows" || runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		_, err := os.Stat(path.Join(STATIC_DIR, "js"))
		if err != nil {
			Info.Printf("Creating %s", path.Join(STATIC_DIR,"js"))
			os.Mkdir(path.Join(STATIC_DIR,"js"), os.ModeSticky | 0755)
		}
	} else {
		Error.Fatal("This OS is not supported!")
	}
}

func fonts_dir() {
	if runtime.GOOS == "windows" || runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		_, err := os.Stat(path.Join(STATIC_DIR, "fonts"))
		if err != nil {
			Info.Printf("Creating %s", path.Join(STATIC_DIR,"fonts"))
			os.Mkdir(path.Join(STATIC_DIR,"fonts"), os.ModeSticky | 0755)
		}
	} else {
		Error.Fatal("This OS is not supported!")
	}
}





func downloadFromUrl(url string, outDir string) {
	// https://github.com/thbar/golang-playground/blob/master/download-files.go

	tokens := strings.Split(url, "/")
	fileName := path.Join(outDir, tokens[len(tokens)-1])
	fileName = strings.Split(fileName,"?")[0]

	// CHECK IF FILE EXISTS
	_, file_err := os.Stat(fileName)
	if file_err == nil {
		return
	}

	Info.Println("Downloading", url)

	output, err := os.Create(fileName)
	if err != nil {
		Error.Println("Error while creating", filepath.Base(fileName), "-", err)
		return
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		Error.Println("Error while downloading", url, "-", err)
		return
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	// _, err = io.Copy(output, response.Body)
	if err != nil {
		Error.Println("Error while downloading", url, "-", err)
		return
	}
	Info.Println(n, "bytes downloaded.")

}



func dir_init() {
	Info.Println("Running installation for", runtime.GOOS)
	// Create Application Directories
	Info.Println("Checking app directories")
	setup_music_dir()
	ninja_dir()
	ninja_music_dir()
	static_dir()

	// Icons & Images
	Info.Println("Checking resource files")
	downloadFromUrl("https://raw.githubusercontent.com/sjsafranek/musicninjaplayer/master/static/logo.png", STATIC_DIR)
	downloadFromUrl("https://raw.githubusercontent.com/sjsafranek/musicninjaplayer/master/static/favicon.ico", STATIC_DIR)
	downloadFromUrl("https://raw.githubusercontent.com/sjsafranek/musicninjaplayer/master/static/error.png", STATIC_DIR)

	// Javascript and Stylesheets
	downloadFromUrl("https://maxcdn.bootstrapcdn.com/font-awesome/4.5.0/fonts/fontawesome-webfont.woff2?v=4.5.0", path.Join(STATIC_DIR,"fonts"))
	downloadFromUrl("https://maxcdn.bootstrapcdn.com/font-awesome/4.5.0/fonts/fontawesome-webfont.ttf?v=4.5.0", path.Join(STATIC_DIR,"fonts"))
	downloadFromUrl("https://maxcdn.bootstrapcdn.com/font-awesome/4.5.0/css/font-awesome.min.css", path.Join(STATIC_DIR,"css"))
	downloadFromUrl("https://ajax.googleapis.com/ajax/libs/jquery/2.1.3/jquery.min.js", path.Join(STATIC_DIR,"js"))
	downloadFromUrl("https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/js/bootstrap.min.js", path.Join(STATIC_DIR,"js"))
	downloadFromUrl("https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/css/bootstrap.min.css", path.Join(STATIC_DIR,"css"))
	downloadFromUrl("https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/css/bootstrap-theme.min.css", path.Join(STATIC_DIR,"css"))

}




/*

https://gobyexample.com/methods


*/