package main

import (
	"math/rand"
    "reflect"
    "unsafe"
    "time"
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



func BytesToString(b []byte) string {
    bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
    sh := reflect.StringHeader{bh.Data, bh.Len}
    return *(*string)(unsafe.Pointer(&sh))
}

func StringToBytes(s string) []byte {
    sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
    bh := reflect.SliceHeader{sh.Data, sh.Len, 0}
    return *(*[]byte)(unsafe.Pointer(&bh))
}