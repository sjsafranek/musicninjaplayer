package main

import (
	"math/rand"
	// "reflect"
	// "unsafe"
	"time"
	"os/user"
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


// func BytesToString(b []byte) string {
// 	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
// 	sh := reflect.StringHeader{bh.Data, bh.Len}
// 	return *(*string)(unsafe.Pointer(&sh))
// }

// func StringToBytes(s string) []byte {
// 	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
// 	bh := reflect.SliceHeader{sh.Data, sh.Len, 0}
// 	return *(*[]byte)(unsafe.Pointer(&bh))
// }