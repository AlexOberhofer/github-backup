package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
)

func GetStats() {

	ctx := context.Background()

	client := github.NewClient(nil)

	repos, _, err := client.Repositories.List(ctx, "alexoberhofer", nil)

	fmt.Printf("Stats: \n")

	for _, repos := range repos {
		fmt.Printf(repos.GetName())
		fmt.Printf("\n")
	}

	if err != nil {
		fmt.Printf("Whoops...\n")
	}
}
