package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func init() {
	fmt.Println(`Starting...`)
}

var (
	name        = flag.String("name", "", "Name of Repo to be created....")
	description = flag.String("description", "", "Description of created repo.")
)

func main() {
	flag.Parse()
	token := os.Getenv("GITHUB_AUTH_TOKEN")
	if token == "" {
		log.Fatal("Unauthorized: No token present.")
	}

	if *name == "" {
		log.Fatal("No name: New Repos must be given name")
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	r := &github.Repository{Name: name, Description: description}
	repo, _, err := client.Repositories.Create(ctx, "", r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(`Created: `, repo.GetName())
	if repo.GetName() != "" {
		url := fmt.Sprintf("https://github.com/SachinMaharana/%s", repo.GetName())
		cmd := exec.Command("git", "clone", url)
		out, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(out))
	}
	fmt.Println(`Ready....Start Coding....`)

	// https://github.com/SachinMaharana/test_repo_working.git
}
