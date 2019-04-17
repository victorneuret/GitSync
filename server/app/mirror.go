package app

import (
	"github.com/victorneuret/GitSync/config"
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