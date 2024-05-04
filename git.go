package main

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func GitPull() {
	path := Repo.Directory
	Log1.Println("git pull to " + path)

	// We instantiate a new repository targeting the given path (the .git folder)

	r, err := git.PlainOpen(path)
	if err != nil {
		Log1.Println(err.Error())
		return
	}

	// Get the working directory for the repository
	w, err := r.Worktree()
	if err != nil {
		Log1.Println(err.Error())
		return
	}

	// Pull the latest changes from the origin remote and merge into the current branch
	Log1.Println("git pull origin")

	err = w.Pull(&git.PullOptions{
		RemoteName: "origin",
		Auth: &http.BasicAuth{
			Username: Conf.GitServer.User,
			Password: Conf.GitServer.AccessToken,
		},
	})
	if err != nil {
		Log1.Println(err.Error())
		return
	}

	// Print the latest commit that was just pulled
	ref, err := r.Head()
	if err != nil {
		Log1.Println(err.Error())
		return
	}
	commit, err := r.CommitObject(ref.Hash())
	if err != nil {
		Log1.Println(err.Error())
		return
	}

	fmt.Println(commit)
}

func GitClone() {

	path := Repo.Directory
	//url := Conf.BuildRepositoryAccessUrl()
	url := Conf.GitServer.GroupURL + "/" + Repo.Name + ".git"
	Log1.Println("git clone " + url)

	_, err := git.PlainClone(path, false, &git.CloneOptions{
		URL: url,
		Auth: &http.BasicAuth{
			Username: Conf.GitServer.User,
			Password: Conf.GitServer.AccessToken,
		},
		//	Progress: os.Stdout,
	})

	if err != nil {
		Log1.Println(err.Error())
		return
	}
}
