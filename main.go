package main

import (
	"log"
	"os"
	"webserver/server"
)

// Reads the name of the html files
// in the given directory and its subdirectories
func ReadDirRecursive(dir string, templates *[]string) {

	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			ReadDirRecursive(dir+"/"+file.Name(), templates)
		} else if file.Name()[len(file.Name())-5:] == ".html" {
			*templates = append(*templates, dir+"/"+file.Name())
		}
	}
}

func main() {
	s := server.New()

	templates := make([]string, 0)
	ReadDirRecursive("./templates", &templates)
	err := s.Start(":80", templates)
	if err != nil {
		log.Fatal(err)
	}
}
