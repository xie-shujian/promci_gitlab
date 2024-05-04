package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
)

func MyHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("X-Gitlab-Token")
	if Conf.PromciApp.AccessToken != token {
		Log1.Println("token not match")
		return
	}

	//repository

	body, err := io.ReadAll(r.Body)
	if err != nil {
		Log1.Println(err.Error())
		return
	}
	defer r.Body.Close()
	var myEvent PushEvent
	err2 := json.Unmarshal(body, &myEvent)
	if err2 != nil {
		Log1.Println(err2.Error())
		return
	}
	repo_url := myEvent.Repository.GitHTTPURL
	Log1.Println(repo_url)

	Conf.FindAndSetRepo(repo_url)

	if Repo == nil {
		Log1.Println("Repository not found")
		return
	}
	//git pull or clone
	git_dir := Repo.Directory + "/.git"
	_, dirErr := os.Stat(git_dir)
	if errors.Is(dirErr, os.ErrNotExist) {
		Log1.Println("Directory not found")
		GitClone()
	} else {
		GitPull()
	}
	//reload config
	Reload()
}
