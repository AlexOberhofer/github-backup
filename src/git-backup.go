package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-git/go-git"
)

var repositories [5]string = [5]string{
	"https://github.com/AlexOberhofer/SDL2-GNUBoy.git",
	"https://github.com/AlexOberhofer/Space-Invaders.git",
	"https://github.com/AlexOberhofer/Chip-8-Emulator.git",
	"https://github.com/AlexOberhofer/pylife.git",
	"https://github.com/AlexOberhofer/sdl2-doom.git"}

func ListRepositories() {
	for _, element := range repositories {
		fmt.Printf(element)
		fmt.Printf("\n")
	}
}

func DoAllClones(url string, directory string) {
	fmt.Printf("Attempting to clone repository: %s\n", url)

	git.PlainClone(directory, false, &git.CloneOptions{
		URL:               url,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})
}

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

	fmt.Printf("The following repositories will be cloned: \n")

	ListRepositories()

	//GetStats()

	time := GetCurrentTimeStamp()

	CreateDirIfNotExist(directory)

	fmt.Printf(time)
	fmt.Printf("\n")

	fmt.Printf("Attempting to clone repository: %s\n", url)

	//git.PlainClone(directory, false, &git.CloneOptions{
	//URL:               url,
	//	RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	//})

	fmt.Printf("Completed!\n")
}
