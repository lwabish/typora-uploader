package main

import (
	"encoding/json"
	"fmt"
	"github.com/lwabish/typora-qiniu-uploader/pkg"
	"io/ioutil"
	"os"
	"path"
)

var images []string

var version string

func main() {

	// get logger
	log := pkg.InitOrGetLogger()

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln(err)
	}

	// export default config file if not exists
	configFilePath := path.Join(homeDir, pkg.ConfigDirname, pkg.ConfigFilename)
	if !pkg.PathExists(configFilePath) {
		defaultConfig, err := json.MarshalIndent(pkg.NewConfig(), "", "")
		if err != nil {
			log.Fatalln(err)
		}
		if err = ioutil.WriteFile(configFilePath, defaultConfig, 0644); err != nil {
			log.Fatalln(err)
		}
		log.Fatalf("Can't find config file, a default one has been written to:  %s, please edit it and rerun this program", configFilePath)
	}

	if len(os.Args) < 2 {
		printUsage()
	}

	images = os.Args[1:]
	log.Println("Images to upload accepted from args: ", images)

	config := pkg.LoadConfig(configFilePath)
	qClient := pkg.NewQiNiuClient(config.AccessKey, config.SecretKey, config.Bucket, config.UseHTTPS, config.UseCdnDomains, config.Domain, config.SubDir)
	imageUris := qClient.UploadImages(images)

	for _, imageUri := range imageUris {
		log.Println(imageUri)
	}
	log.Println("Jobs done!")
	for _, imageUri := range imageUris {
		// use fmt instead of log because typora will recognize image url from stdout
		fmt.Println(imageUri)
	}

}

func printUsage() {
	fmt.Println("typora-qiniu-uploader(tql):", version)
	fmt.Println()
	fmt.Println("Usage: tql <path1 path2 ...>")
	os.Exit(0)
}
