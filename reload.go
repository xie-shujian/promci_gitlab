package main

import "net/http"

func Reload() {
	if Repo.ReloadUrl == "" {
		return
	}
	req, err := http.NewRequest("POST", Repo.ReloadUrl, http.NoBody)
	if err != nil {
		Log1.Println(err.Error())
		return
	}
	req.Header.Set("Content-Type", "")

	resp, err := http.Post(req.URL.String(), "", http.NoBody)
	if err != nil {
		Log1.Println(err.Error())
		return
	}
	Log1.Println("Reload success " + resp.Status)
}
