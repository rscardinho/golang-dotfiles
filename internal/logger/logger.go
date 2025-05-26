package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/rscardinho/golang-dotfiles/cmd/helpers"
)

var (
	logDir = "logs"
)

func Load() *os.File {
	os.MkdirAll(logDir, os.ModePerm)

	logFilename := filepath.Join(logDir, fmt.Sprintf("install-%s.log", time.Now().Format("2006-01-02-15-04-05")))
	logPath, err := helpers.RelativeFilePath(logFilename)
	if err != nil {
		log.Fatalf("Failed to build log file path: %v", err)
	}

	logFile, err := os.Create(logPath)
	if err != nil {
		log.Fatalf("Failed to create log file: %v", err)
	}

	return logFile
}

func Filename(file os.File) string {
	return filepath.Base(file.Name())
}
