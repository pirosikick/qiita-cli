package main

import (
	"encoding/json"
	"io/ioutil"
)

var ConfigFile = ".qiitarc"

type Config struct {
	Token string
}

func LoadConfig() Config {
	str, _ := ioutil.ReadFile(ConfigFile)

	var config Config
	json.Unmarshal(str, &config)

	return config
}
