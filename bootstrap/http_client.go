package bootstrap

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type HttpClientConfig struct {
	Config httConfig `json:"origin_proxy" yaml:"origin_proxy"`
}

type httConfig struct {
	Scheme      string `json:"scheme" yaml:"scheme"`
	Host        string `json:"host" yaml:"host"`
	Port        int    `json:"port" yaml:"port"`
	ContextPath string `json:"context_path" yaml:"context_path"`
}

var (
	HttpClient httpClient
)

type httpClient struct {
	Client      http.Client
	Host        string
	Port        int
	ContextPath string
	Scheme      string
}

func HttpClientInit(path string) error {
	dataBytes, err := os.ReadFile(path)
	if err != nil {
		log.Println("ReadFile HttpClient err: ", err)
		return err
	}
	conf := HttpClientConfig{}
	err = yaml.Unmarshal(dataBytes, &conf)
	if err != nil {
		log.Println("Unmarshal HttpClient err:", err)
		return err
	}
	HttpClient = httpClient{
		Client:      http.Client{},
		Scheme:      conf.Config.Scheme,
		Host:        conf.Config.Host,
		ContextPath: conf.Config.ContextPath,
		Port:        conf.Config.Port,
	}

	fmt.Println("HttpClient init success", HttpClient)
	return nil
}

func (h *httpClient) Get(url string) (resp *http.Response, err error) {
	var abusolute string
	if strings.HasPrefix(url, "/") {
		abusolute = fmt.Sprintf("%s://%s:%d%s", h.Scheme, h.Host, h.Port, h.ContextPath) + url
	} else {
		abusolute = fmt.Sprintf("%s://%s:%d%s", h.Scheme, h.Host, h.Port, h.ContextPath) + "/" + url
	}
	log.Println("the get request url", abusolute)
	return HttpClient.Client.Get(abusolute)
}

// 仅支持json格式参数
func (h *httpClient) Post(url string, body io.Reader) (resp *http.Response, err error) {
	var abusolute string
	if strings.HasPrefix(url, "/") {
		abusolute = fmt.Sprintf("%s://%s:%d%s", h.Scheme, h.Host, h.Port, h.ContextPath) + url
	} else {
		abusolute = fmt.Sprintf("%s://%s:%d%s", h.Scheme, h.Host, h.Port, h.ContextPath) + "/" + url
	}
	log.Println("the post request url", abusolute)
	return HttpClient.Client.Post(abusolute, "application/json", body)
}
