package main

import (
	"math/rand"
	"time"
	"os/user"
	"io/ioutil"
	// "reflect"
	// "fmt"
)

// Returns random int between min and max
func randInt(min, max int) int {
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	return rand.Intn(max - min) + min
}

// Checks to see if a string is in a list of strings
// Returns true or false
func stringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

// golang returns -1 module
// https://code.google.com/p/go/issues/detail?id=448
func modulo(x int, y int) int {
	z := x % y
	if z < 0 {
		z += y
	}
	return z
}

// Get home directory
// http://stackoverflow.com/questions/7922270/obtain-users-home-directory
func homeDir() string {
	usr, err := user.Current()
	if err != nil {
		Error.Fatal( err )
	}
	return usr.HomeDir
}

// Gets all files in a directory
func getFilesInDirectory(directory string) []string {
	results := []string{}
	files, _ := ioutil.ReadDir(directory)
	for i := 0; i < len(files); i++ {
		if files[i].IsDir() == false {
			results = append(results, files[i].Name())
		}
	}
	return results
}

// Gets all folders in directory
func getFoldersInDirectory(directory string) []string {
	folders := []string{}
	files, _ := ioutil.ReadDir(directory)
	for i := 0; i < len(files); i++ {
		if files[i].IsDir() {
			folders = append(folders, files[i].Name())
		}
	}
	return folders
}




// // Gets object methods
// func getMethods() {
// 	fooType := reflect.TypeOf(files[i])
// 	for i := 0; i < fooType.NumMethod(); i++ {
// 		method := fooType.Method(i)
// 		Info.Println(method.Name)
// 	}
// }

