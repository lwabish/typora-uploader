package pkg

import (
	"io"
	"log"
	"os"
	"path"
	"sync"
)

var (
	once   sync.Once
	logger *log.Logger
)

const logFileName = "tqu.log"

func InitOrGetLogger() *log.Logger {
	once.Do(func() {

		log.SetFlags(log.Lshortfile | log.LstdFlags)
		log.Println("Init logger...")

		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatalln(err)
		}

		if !PathExists(path.Join(homeDir, ConfigDirname)) {
			if err = os.MkdirAll(path.Join(homeDir, ConfigDirname), 0755); err != nil {
				log.Fatalln("Can't Create Config Dir: ", err)
			}
		}

		f, err := os.OpenFile(path.Join(homeDir, ConfigDirname, logFileName), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatalln(err)
		}
		writers := []io.Writer{
			f,
			os.Stdout,
		}
		multipleWriter := io.MultiWriter(writers...)
		logger = log.New(multipleWriter, "", log.Lshortfile|log.LstdFlags)
		printHeader()
		logger.Println("Init logger done...")
	})
	return logger
}

func printHeader() {
	logger.Println()
	logger.Println("**********************tqu started**********************")
	logger.Println()
}
