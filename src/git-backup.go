package main

import (
	"flag"
	"fmt"
	"os"
	"time"
	"strings"

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

func DoAClone(url string) {
	fmt.Printf("Attempting to clone repository: %s\n", url)

	//pull repo name
	urlStrings := strings.Split(url, "/")
	repoNameLong := urlStrings[len(urlStrings) - 1]	
	repoNameLongSplit := strings.Split(repoNameLong, ".")
	repoName := repoNameLongSplit[0]

	CreateDirIfNotExist(repoName)

	git.PlainClone(repoName, false, &git.CloneOptions{
		URL:               url,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})
}

//TODO probably move to file operations file
func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}

//TODO (actually make some flags...)
func Usage() {
	fmt.Printf("Usage: -q <username> \n")
}

//Return current timestamp (to be used for backup filename)
func GetCurrentTimeStamp() string {
	return time.Now().Format(time.RFC850)
}

func main() {

	username := flag.String("q", "", "Github Username to query")
	url := flag.String("cs", "", "Single Repository URL to clone")
	flag.Parse()

	fmt.Printf("Backup Utility started: %s\n", GetCurrentTimeStamp())

	//directory := "./my-repo"

	//fmt.Printf("The following repositories will be cloned: \n")

	//ListRepositories()

	if *username != "" {
		GetStats(*username)
	}

	if *url != "" {
		DoAClone(*url)
	}

	fmt.Printf("Backup Utility completed: %s\n", GetCurrentTimeStamp())
}
