package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	// ファイルを開く
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	// ログの出力方法を指定
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	// ログの設定
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	// ログの出力先を設定
	log.SetOutput(multiLogFile)
}
