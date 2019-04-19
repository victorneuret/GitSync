package app

import (
	"context"
	"github.com/google/go-github/github"
	"github.com/victorneuret/GitSync/config"
	"golang.org/x/oauth2"
	"net/http"
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

func CreateGitHubHook(name string, login string, token string) bool {
	client := connectUser(token)
	active := true

	hookConfig := map[string]interface{}{
		"url":          config.Config.URL + "/hook/?login=" + login + "&name=" + name,
		"content_type": "json",
	}
	myHook := &github.Hook{
		Config: hookConfig,
		Events: []string{"push"},
		Active: &active,
	}
	_, _, err := client.Repositories.CreateHook(context.Background(), login, name, myHook)
	if err != nil {
		return false
	}
	return true
}

func HandleWebHook(w http.ResponseWriter, r *http.Request) {
	login, err := r.URL.Query()["login"]
	if !err || len(login[0]) < 1 {
		return
	}

	name, err := r.URL.Query()["name"]
	if !err || len(name[0]) < 1 {
		return
	}

	SyncMirror(name[0], login[0])
}