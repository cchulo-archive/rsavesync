package logger

import (
	"io"
	"log"
	"os"
	"time"
	"fmt"
	"path/filepath"
)

func InitLogger() (*log.Logger, *os.File, error) {
	logDir := os.ExpandEnv("$HOME/.config/rsavesync/logs")
	err := os.MkdirAll(logDir, 0755)
	if err != nil {
		return nil, nil, err
	}

	currentTime := time.Now().Format("2006-01-02")
	logFileName := fmt.Sprintf("rsavesync.%s.log", currentTime)
	logFilePath := filepath.Join(logDir, logFileName)

	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, nil, err
	}

	multiWriter := io.MultiWriter(os.Stdout, logFile)
	logger := log.New(multiWriter, "", log.LstdFlags)
	return logger, logFile, nil
}
