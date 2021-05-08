package pkg

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	//Zone          int    `json:"zone"`
	Bucket        string `json:"Bucket"`
	UseHTTPS      bool   `json:"use_https"`
	UseCdnDomains bool   `json:"use_cdn_domains"`
	Domain        string `json:"domain"`
	SubDir        string `json:"sub_dir"`
}

const (
	ConfigDirname  = ".config/typora-qiniu-uploader"
	ConfigFilename = "config.json"
)

func LoadConfig(path string) (config *Config) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(b)
	}

	err = json.Unmarshal(b, &config)
	if err != nil {
		panic(err)
	}
	return
}

// NewConfig generate an empty config
func NewConfig() *Config {
	return &Config{
		AccessKey: "",
		SecretKey: "",
		//Zone:          0,
		Bucket:        "",
		UseHTTPS:      true,
		UseCdnDomains: true,
		Domain:        "",
	}
}
