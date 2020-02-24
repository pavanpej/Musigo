package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// constant to store .mp3 file extension string
const Mp3Ext string = ".mp3"

func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == Mp3Ext {
			name := strings.TrimSuffix(info.Name(), Mp3Ext)
			*files = append(*files, name)
		}
		return nil
	}
}

func recursiveDirectoryParse(path string) []string {
	var music []string
	err := filepath.Walk(path, visit(&music))
	if err != nil {
		panic(err)
	}
	return music
}

func singleDirectoryParse(path string) []string {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatal("Path does not exist")
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	var music []string
	for _, f := range files {
		if !f.IsDir() && filepath.Ext(f.Name()) == Mp3Ext {
			music = append(music, f.Name())
		}
	}

	return music
}

func main() {
	// store final music file names in this variable
	var music []string

	// handle flag for recursive/single search
	var rec *bool
	rec = flag.Bool("r", false, "search recursively or not")
	flag.Parse()

	// by default, the current dir
	root := "."

	// TODO handle multiple folders and output into different files for each folder
	args := os.Args[1:]
	if len(args) != 0 {
		for _, arg := range args {
			if !strings.HasPrefix(arg, "-") {
				root = arg
				break
			}
		}
	}

	if *rec {
		music = recursiveDirectoryParse(root)
	} else {
		music = singleDirectoryParse(root)
	}

	for _, song := range music {
		fmt.Println(song)
	}

}
