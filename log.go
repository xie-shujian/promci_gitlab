package main

import (
	"log"
	"os"
)

func InitLogger() {
	fileName := Conf.PromciApp.LogFile

	logWriter, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic("create log file failed " + fileName + err.Error())
	}
	Log1 = log.New(logWriter, "", log.LstdFlags)
}
