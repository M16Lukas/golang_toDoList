package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSetting(logFile string) {
	// O_RDWR   : ファイルの読み書き
	// O_CREATE : ファイルの作成
	// O_APPEND : ファイルの追記
	logfile, err := os.OpenFile(
		logFile,
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0666,
	)

	if err != nil {
		log.Fatal(err)
	}

	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}
