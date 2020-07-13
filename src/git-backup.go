package main

import (
	"fmt"
	"os"

	"github.com/go-git/go-git"
)

func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	url := "https://github.com/AlexOberhofer/SDL2-GNUBoy.git"
	directory := "./backup"

	CreateDirIfNotExist(directory)

	git.PlainClone(directory, false, &git.CloneOptions{
		URL:               url,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})

	fmt.Printf("Completed!\n")
}
