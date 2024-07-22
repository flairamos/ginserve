package bootstrap

import (
	"log"
	"os"

	"github.com/minio/minio-go"
	"gopkg.in/yaml.v3"
)

type minioConfig struct {
	AccessKeyID     string `yaml:"accessKeyID"`
	SecretAccessKey string `yaml:"secretAccessKey"`
	Endpoint        string `yaml:"endpoint"`
	UseSSL          bool   `yaml:"useSSL"`
	Bucket          string `yaml:"bucket"`
}

type MinioConfig struct {
	Config minioConfig `yaml:"minio_config"`
}

var (
	MinioClient *minio.Client
	Bucket      string
)

func MinioInit(path string) error {
	dataBytes, err := os.ReadFile(path)
	if err != nil {
		log.Println("ReadFile Postgres err: ", err)
		return err
	}
	conf := MinioConfig{}
	err = yaml.Unmarshal(dataBytes, &conf)
	if err != nil {
		log.Println("Unmarshal Postgres err:", err)
		return err
	}
	MinioClient, err = minio.New(conf.Config.Endpoint, conf.Config.AccessKeyID, conf.Config.SecretAccessKey, conf.Config.UseSSL)
	if err != nil {
		log.Println("minio client create err: ", err)
		return err
	}
	log.Println("minio client created")
	// 判断bucket是否存在
	exists, err := MinioClient.BucketExists(conf.Config.Bucket)
	if err != nil {
		log.Println("BucketExists err: ", err)
		return err
	}
	if !exists {
		err = MinioClient.MakeBucket(conf.Config.Bucket, "local")
		if err != nil {
			log.Println("MakeBucket err: ", err)
			return err
		}
	}
	Bucket = conf.Config.Bucket
	return nil
}
