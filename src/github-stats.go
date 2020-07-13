package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/github"
)

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
