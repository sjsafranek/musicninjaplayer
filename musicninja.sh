#!/bin/bash

export GOPATH=$HOME/Desktop/musicninjaplayer

function build_pi {
	$HOME/go/bin/go build *.go
}

function build {
	go build *.go
}

function run_pi {
	$HOME/go/bin/go run *.go
}

function run {
	go run *.go
}

case "$1" in
	"build")
		echo "Building app..."
		# $HOME/go/bin/go build *.go
		build_pi || build
		echo "Complete!!"
		;;
	"run")
		# $HOME/go/bin/go run *.go
		run_pi || run
		;;
	*)
		echo "Unknown argument: $1"
	;;
esac

