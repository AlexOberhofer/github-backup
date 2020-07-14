/* ********************************************************************************
*    Git Operations
*    Copyright (C) 2020  Alex Oberhofer
*
*    This program is free software: you can redistribute it and/or modify
*    it under the terms of the GNU General Public License as published by
*    the Free Software Foundation, either version 3 of the License, or
*    (at your option) any later version.
*
*    This program is distributed in the hope that it will be useful,
*    but WITHOUT ANY WARRANTY; without even the implied warranty of
*    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
*    GNU General Public License for more details.
*
*    You should have received a copy of the GNU General Public License
*    along with this program.  If not, see <https://www.gnu.org/licenses/>.
********************************************************************************/

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-git/go-git"
	"github.com/google/go-github/github"
)

func DoAClone(url string) {
	fmt.Printf("Attempting to clone repository: %s\n", url)

	//pull repo name
	urlStrings := strings.Split(url, "/")
	repoNameLong := urlStrings[len(urlStrings)-1]
	repoNameLongSplit := strings.Split(repoNameLong, ".")
	repoName := repoNameLongSplit[0]

	CreateDirIfNotExist(repoName)

	git.PlainClone(repoName, false, &git.CloneOptions{
		URL:               url,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})
}

func DoACloneDir(url string, dir string) {
	fmt.Printf("# Attempting to clone repository: %s\n", url)

	//pull repo name
	urlStrings := strings.Split(url, "/")
	repoNameLong := urlStrings[len(urlStrings)-1]
	repoNameLongSplit := strings.Split(repoNameLong, ".")
	repoName := repoNameLongSplit[0]

	PathToClone := dir + repoName

	CreateDirIfNotExist(PathToClone)

	git.PlainClone(PathToClone, false, &git.CloneOptions{
		URL:               url,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})
}

func CloneAllPublicRepos(username string) {
	ctx := context.Background()

	client := github.NewClient(nil)

	repos, _, err := client.Repositories.List(ctx, username, nil)

	for _, repos := range repos {
		DoACloneDir(repos.GetGitURL(), "backup/")
	}

	if err != nil {
		log.Fatal(err)
	}
}

func GetStats(username string) {

	ctx := context.Background()

	client := github.NewClient(nil)

	repos, _, err := client.Repositories.List(ctx, username, nil)

	fmt.Printf("###################################################################################\n")
	fmt.Printf("#\n")
	fmt.Printf("# Compiling (Public) Repository details for %s: \n", username)
	fmt.Printf("#\n")
	fmt.Printf("###################################################################################\n")

	for _, repos := range repos {
		fmt.Printf("# Repository Name: %s\n", repos.GetName())
		fmt.Printf("# Git URL: %s\n", repos.GetGitURL())
	}

	if err != nil {
		fmt.Printf("###################################################################################\n")
		fmt.Printf("# Whoops... Are you sure that's the correct username? \n")
		fmt.Printf("###################################################################################\n")
		os.Exit(1)
	}

	fmt.Printf("###################################################################################\n")
	fmt.Printf("# Query completed without error\n")
	fmt.Printf("###################################################################################\n")
}
