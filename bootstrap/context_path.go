package bootstrap

import (
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

var ContextPath string

type contextPathConf struct {
	ContextPath string `yaml:"context_path"`
}

func ContextPathInit(path string) error {
	dataBytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("ReadFile err:", err)
		return err
	}
	conf := contextPathConf{}
	err = yaml.Unmarshal(dataBytes, &conf)
	if err != nil {
		log.Fatal("Unmarshal err:", err)
		return err
	}
	if strings.HasPrefix(conf.ContextPath, "/") {
		ContextPath = conf.ContextPath
	} else {
		ContextPath = "/" + conf.ContextPath
	}
	return nil
}
