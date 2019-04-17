package app

import (
	"github.com/victorneuret/GitSync/config"
	"os/exec"
)

func CreateBlihRepo(name string, login string, token string) bool {
	repoPath := config.Config.RepoPath + name
	blihSSH := config.Config.BlihSSH + ":/" + login + "/" + name

	cmd := exec.Command("blih", "-u", login, "-t", token, "repository", "create", name)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}
	cmd = exec.Command("blih", "-u", login, "-t", token, "repository", "setacl", name, "ramassage-tek", "r")
	_, err = cmd.CombinedOutput()
	if err != nil {
		return false
	}
	cmd = exec.Command("mkdir", "-p", repoPath)
	_, err = cmd.CombinedOutput()
	if err != nil {
		return false
	}
	cmd = exec.Command("git", "clone", blihSSH, repoPath)
	_, err = cmd.CombinedOutput()
	if err != nil {
		return false
	}
	return true
}