package main

import (
	"log"
	"net/http"
)

var Log1 *log.Logger
var Conf *Config
var Repo *Repository

func main() {
	ReadConfig()
	InitLogger()
	http.HandleFunc("/promci", MyHandler)
	Log1.Println("Starting HTTP server on port 8866")
	err := http.ListenAndServe(":8866", nil)
	if err != nil {
		Log1.Fatal("Error starting HTTP server: ", err)
	}
}
