package main

import (
	"fmt"
	"os"
)

func main() {
	var saveFileUnixTime int64
	for {
		f1, _ := os.Lstat("file.1")
		f1UnixTime := f1.ModTime().Unix()
		if saveFileUnixTime != f1UnixTime {
			saveFileUnixTime = f1UnixTime
			fmt.Println(f1.Name(), saveFileUnixTime)
		}
	}
}
