package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v25/github"
	"golang.org/x/oauth2"
)

func main() {
	// get access token from args
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Print("Usage: ozymandias [AccessToken]")
		os.Exit(0)
	}

	accessToken := args[0]

	// connect to github api
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list all organizations
	orgs, _, err := client.Organizations.List(ctx, "", nil)
	if err != nil {
		panic(err)
	}
	fmt.Print("%v", orgs)

	// list all repositories
	repos, _, err := client.Repositories.List(ctx, "", nil)
	if err != nil {
		panic(err)
	}
	fmt.Print("%v", repos)
}
