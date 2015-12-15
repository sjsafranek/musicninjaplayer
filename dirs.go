package main

import (
	"os"
	"path"
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
	if runtime.GOOS == "windows" || runtime.GOOS == "linux" {
		BASE_DIR = path.Join(homeDir(),"Music")
		BASE_DIR = strings.Replace(BASE_DIR,"\\","/",-1)
		_, err := os.Stat(BASE_DIR)
		if err != nil {
			Info.Println("Creating %s", BASE_DIR)
			os.Mkdir(BASE_DIR, os.ModeSticky | 0755)
		}
	} else {
		Error.Fatal("This OS is not supported!")
	}
}

func setup_ninja_dir() {
	if runtime.GOOS == "windows" || runtime.GOOS == "linux" {
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

func setup_ninja_music_dir() {
	if runtime.GOOS == "windows" || runtime.GOOS == "linux" {
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

func setup_ninja_static_dir() {
	if runtime.GOOS == "windows" || runtime.GOOS == "linux" {
		_, err := os.Stat(path.Join(BASE_DIR,"Ninja", "static"))
		if err != nil {
			Info.Printf("Creating %s", path.Join(BASE_DIR,"Ninja", "static"))
			os.Mkdir(path.Join(BASE_DIR,"Ninja", "static"), os.ModeSticky | 0755)
		}
		STATIC_DIR = path.Join(BASE_DIR,"Ninja", "static")
	} else {
		Error.Fatal("This OS is not supported!")
	}
}



func downloadFromUrl(url string, outDir string) {
	// https://github.com/thbar/golang-playground/blob/master/download-files.go

	tokens := strings.Split(url, "/")
	fileName := path.Join(outDir, tokens[len(tokens)-1])

	// CHECK IF FILE EXISTS
	_, file_err := os.Stat(fileName)
	if file_err == nil {
		Info.Printf("File %s exists", fileName)
		return
	} else {
		Info.Printf("File %s not found", fileName)
	}

	Info.Println("Downloading", url, "to", fileName)

	// TODO: check file existence first with io.IsExist
	output, err := os.Create(fileName)
	if err != nil {
		Info.Println("Error while creating", fileName, "-", err)
		return
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		Info.Println("Error while downloading", url, "-", err)
		return
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		Info.Println("Error while downloading", url, "-", err)
		return
	}

	Info.Println(n, "bytes downloaded.")


}



func dir_init() {
	// Create Application Directories
	setup_music_dir()
	setup_ninja_dir()
	setup_ninja_music_dir()
	setup_ninja_static_dir()

	downloadFromUrl("https://raw.githubusercontent.com/sjsafranek/musicninjaplayer/master/static/logo.png", STATIC_DIR)
	downloadFromUrl("https://raw.githubusercontent.com/sjsafranek/musicninjaplayer/master/static/favicon.ico", STATIC_DIR)
	downloadFromUrl("https://raw.githubusercontent.com/sjsafranek/musicninjaplayer/master/static/error.png", STATIC_DIR)

}