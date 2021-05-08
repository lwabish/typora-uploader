package main

import (
	"encoding/json"
	"github.com/lwabish/typora-qiniu-uploader/pkg"
	"io/ioutil"
	"log"
	"os"
	"path"
)

var images []string

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln(err)
	}

	configFilePath := path.Join(homeDir, pkg.ConfigDirname, pkg.ConfigFilename)
	defaultConfig, err := json.MarshalIndent(pkg.NewConfig(), "", "")
	if err != nil {
		log.Fatalln(err)
	}

	if !pkg.PathExists(configFilePath) {
		if err = os.MkdirAll(path.Join(homeDir, pkg.ConfigDirname), 0755); err != nil {
			log.Fatalln("Can't Create Config Dir: ", err)
		}
		if err = ioutil.WriteFile(configFilePath, defaultConfig, 0644); err != nil {
			log.Fatalln(err)
		}
		log.Fatalf("Can't find config file, a default one has been written to:  %s, please edit it and rerun this program", configFilePath)
	}

	images = os.Args[1:]
	log.Println("Images to upload accepted from args: ", images)

	config := pkg.LoadConfig(configFilePath)
	qClient := pkg.NewQiNiuClient(config.AccessKey, config.SecretKey, config.Bucket, config.UseHTTPS, config.UseCdnDomains, config.Domain, config.SubDir)
	qClient.UploadImages(images)

}
