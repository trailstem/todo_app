package utils

import (
	"io"
	"log"
	"os"
)


func LoggingSettings(logFile string) {
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
	//OpenFileでlogfileを読み込んで、読み書き、作成、追記を指定して、パーミッションが0666と指定
}
