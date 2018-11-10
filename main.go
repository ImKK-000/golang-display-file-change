package main

import (
	"fmt"
	"io/ioutil"
	"path"
)

var ignoreFile = map[string]bool{}

func WalkInDirectories(pathFile string) {
	files, _ := ioutil.ReadDir(pathFile)
	for _, file := range files {
		if ignoreFile[file.Name()] {
			continue
		}
		if file.IsDir() {
			fmt.Println(file.Name(), file.Size(), file.ModTime().Unix())
			WalkInDirectories(path.Join(pathFile, file.Name()))
		}
	}
}

func main() {
	// ignore files
	ignoreFile[".git"] = true

	// start to walk directory from current path
	WalkInDirectories(path.Dir(""))
}
