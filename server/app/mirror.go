package app

import (
	"github.com/victorneuret/GitSync/config"
	"github.com/victorneuret/GitSync/database"
	"os/exec"
)

func SetMirror(name string, blihLogin string, login string) bool {
	blihSSH := config.Config.BlihSSH + ":/" + blihLogin + "/" + name

	cmd := exec.Command("git", "clone", "--mirror", "git@github.com:" + login + "/" + name + ".git")
	cmd.Dir = config.Config.RepoPath
	_, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}

	cmd = exec.Command("git", "remote", "set-url", "--push", "origin", blihSSH)
	cmd.Dir = config.Config.RepoPath + "/" + name + ".git"
	_, err = cmd.CombinedOutput()
	if err != nil {
		return false
	}
	return true
}

func SyncMirror(name string, login string) bool {
	if !database.DB.Where(&database.Repo{Name: name, Owner: login}).First(&database.Repo{}).RecordNotFound() {
		return false
	}

	var user database.User
	if database.DB.Where(&database.User{Login: login}).First(&user).RecordNotFound() {
		return false
	}

	cmd := exec.Command("git", "fetch", "-p", "origin")
	cmd.Dir = config.Config.RepoPath + name + ".git"
	_, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}

	cmd = exec.Command("git", "push", "--mirror")
	cmd.Dir = config.Config.RepoPath + name + ".git"
	_, err = cmd.CombinedOutput()
	if err != nil {
		return false
	}
	return true
}