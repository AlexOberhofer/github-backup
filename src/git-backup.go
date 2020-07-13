package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-git/go-git"
)

var repositories [10]string

func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}

func Usage() {
	fmt.Printf("Usage: \n")
}

func GetCurrentTimeStamp() string {
	return time.Now().Format(time.RFC850)
}

func main() {
	url := "https://github.com/AlexOberhofer/SDL2-GNUBoy.git"
	directory := "./backup"

	time := GetCurrentTimeStamp()

	CreateDirIfNotExist(directory)

	fmt.Printf(time)
	fmt.Printf("\n")

	fmt.Printf("Attempting to clone repository: %s\n", url)

	git.PlainClone(directory, false, &git.CloneOptions{
		URL:               url,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})

	fmt.Printf("Completed!\n")
}
