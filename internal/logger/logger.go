package logger

import (
	"log"
	"os"
	"time"
)

func Logger(content string, verbose bool) {
	if !verbose {
		return
	}

	logger := log.New(os.Stdout, "", 0)
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	logger.SetPrefix("[" + currentTime + "] ")

	logger.Println(content)
}

func ErrorLogger(content error) {
	logger := log.New(os.Stdout, "", 0)
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	logger.SetPrefix("[" + currentTime + "] ")

	logger.Println(content)
}
