package app

import (
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"context"
)

func connectUser(token string) *github.Client {
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	client := oauth2.NewClient(context.Background(), tokenSource)

	githubClient := github.NewClient(client)
	return githubClient
}

func CreateGitHubRepo(name string, private bool, token string) bool {
	client := connectUser(token)
	repo := &github.Repository{
		Name:    github.String(name),
		Private: github.Bool(private),
	}
	_, _, err := client.Repositories.Create(context.Background(), "", repo)
	if err != nil {
		return false
	}
	return true
}