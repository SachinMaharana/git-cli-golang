package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func init() {
	fmt.Println(`Starting...`)
	if os.Getenv("GITHUB_AUTH_TOKEN") == "" {
		log.Fatal("Unauthorized: No token present.")
		os.Exit(1)
	}
}

var (
	name        = flag.String("name", "", "Name of Repo to be created....")
	description = flag.String("desc", "", "Description of created repo....")
)

func createRepo(token string, name *string, description *string) (repo *github.Repository, err error) {
	ctx := context.Background()
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tokenContex := oauth2.NewClient(ctx, tokenSource)
	client := github.NewClient(tokenContex)
	r := &github.Repository{Name: name, Description: description}
	repo, _, err = client.Repositories.Create(ctx, "", r)
	return repo, err
}

func cloneRepo(repo *github.Repository) (err error) {
	repoName := repo.GetName()
	var cloneError error
	if repoName != "" {
		url := fmt.Sprintf("https://github.com/SachinMaharana/%s", repoName)
		cmd := exec.Command("git", "clone", url)
		err := cmd.Run()
		cloneError = err
	}
	return cloneError
}

func main() {
	flag.Parse()
	token := os.Getenv("GITHUB_AUTH_TOKEN")

	if *name == "" {
		log.Fatal("No name: New Repos must be given name")
	}

	repo, err := createRepo(token, name, description)

	check(err)

	fmt.Println(`Repo Created:`, repo.GetName())

	cloneError := cloneRepo(repo)

	check(cloneError)

	log.Println("Git Clone done for", repo.GetName())

	createFile(repo.GetName(), "README.md")

	content := fmt.Sprintf("#%s", repo.GetName())

	writeFile(repo.GetName(), "README.md", content)

	log.Println(`Ready....Start Coding....`)

	// https://github.com/SachinMaharana/test_repo_working.git
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func createFile(path string, fileName string) {
	file, err := os.Create(filepath.Join(path, fileName))
	check(err)
	log.Println(`File Created.`)
	defer file.Close()
}

func writeFile(newPath string, fileName string, content string) {
	file, err := os.OpenFile(filepath.Join(newPath, fileName), os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666)
	check(err)
	byteSlice := []byte(content)
	_, errr := file.Write(byteSlice)
	check(errr)
	defer file.Close()
}
