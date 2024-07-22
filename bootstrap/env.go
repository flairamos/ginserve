package bootstrap

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"gopkg.in/yaml.v3"
)

var Server envconfig

type envconfig struct {
	Env env `yaml:"env"`
}

type env struct {
	Run  string `yaml:"run"`
	Port string `yaml:"port"`
}

func EnvInit(path string) string {
	var file string
	system := runtime.GOOS
	if system == "linux" {
		file = path + "/config/app.yaml"
	} else if system == "windows" {
		file = path + "\\config\\app.yaml"
	} else {
		log.Fatal("Unsupported platform")
	}
	dataBytes, err := os.ReadFile(file)
	if err != nil {
		log.Fatal("ReadFile Env err:", err)
		return ""
	}
	conf := envconfig{}
	err = yaml.Unmarshal(dataBytes, &conf)
	if err != nil {
		log.Fatal("Unmarshal Env err:", err)
		return ""
	}
	Server = conf
	if system == "linux" {
		return path + fmt.Sprintf("/config/app_%s.yaml", conf.Env.Run)
	} else {
		return path + fmt.Sprintf("\\config\\app_%s.yaml", conf.Env.Run)
	}
}
