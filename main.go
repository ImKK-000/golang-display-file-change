package main

import (
	"fmt"
	"io/ioutil"
	"path"
)

var ignoreFile = map[string]bool{}

func WalkInDirectories(previousPathFile, currentDirectory string, tabSize int) {
	currentPathFile := path.Join(previousPathFile, currentDirectory)
	files, _ := ioutil.ReadDir(currentPathFile)
	for _, file := range files {
		if ignoreFile[file.Name()] {
			continue
		}
		printFileNameWithPrefix(file.Name(), tabSize)
		if file.IsDir() {
			WalkInDirectories(currentPathFile, file.Name(), tabSize+1)
		}
	}
}

func printFileNameWithPrefix(filename string, tabSize int) {
	for i := 0; i < tabSize; i++ {
		fmt.Print("|  ")
	}
	fmt.Println("|-", filename)
}

func main() {
	// ignore files
	ignoreFile[".git"] = true
	ignoreFile[".DS_Store"] = true

	// start to walk directory from current path
	WalkInDirectories(path.Dir("."), ".", 0)
}
