package main

import (
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

func ReadConfig() {
	config_file_key := "config.file"
	config_file, found := ReadArgs(config_file_key)
	if !found {
		config_file = "promci.yml"
	}
	//println("read config file" + config_file)
	Conf = &Config{}
	dataBytes, err := os.ReadFile(config_file)
	if err != nil {
		panic(err.Error())
	}
	err = yaml.Unmarshal(dataBytes, Conf)
	if err != nil {
		panic(err.Error())
	}
	//println(Conf.GitServer.GroupURL)
}

func ReadArgs(key string) (string, bool) {
	key = "--" + key
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, key) {
			name, value, found := strings.Cut(arg, "=")
			if found && name == key && value != "" {
				return value, found
			}
		}
	}
	return "", false
}
