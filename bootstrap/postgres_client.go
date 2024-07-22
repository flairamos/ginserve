package bootstrap

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type config struct {
	Adepter  string `yaml:"adepter"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	Ssl      string `yaml:"ssl"`
}

type PostgresConfig struct {
	Config config `yaml:"db_config"`
}

var (
	DB *gorm.DB
)

func PostgresInit(path string) error {
	dataBytes, err := os.ReadFile(path)
	if err != nil {
		log.Println("ReadFile Postgres err: ", err)
		return err
	}
	conf := PostgresConfig{}
	err = yaml.Unmarshal(dataBytes, &conf)
	if err != nil {
		log.Println("Unmarshal Postgres err:", err)
		return err
	}
	dataSource := fmt.Sprintf("%s://%s:%s@%s:%s/%s",
		conf.Config.Adepter, conf.Config.User, conf.Config.Password, conf.Config.Host, conf.Config.Port, conf.Config.Database)
	db, err := gorm.Open(postgres.Open(dataSource), &gorm.Config{})
	if err != nil {
		log.Println("Open Postgres err:", err)
		return err
	}
	DB = db
	log.Println("postgres client created")
	return nil
}
