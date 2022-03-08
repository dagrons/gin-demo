package logging

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/spf13/viper"
)

var (
	LogSavePath string
	LogSaveName string
	LogFileExt  = "log"
	TimeFormat  = "20060102"
)

type option func()

func WithViperConfig() option {
	return func() {
		LogSavePath = viper.GetString("logs.log_dir")
		LogSaveName = viper.GetString("logs.log_save_name")
		fmt.Print("1,", LogSaveName, LogSavePath)
	}
}

func Option(opts ...option) {
	for _, opt := range opts {
		opt()
	}
}

func Init(opts ...option) {
	Option(opts...)
	filePath := getLogFileFullPath()
	F = openLogFile(filePath)
	logger = log.New(F, DefaultPrefix, log.LstdFlags|log.Llongfile)
}

func getLogFilePath() string {
	return fmt.Sprintf("%s", LogSavePath)
}

func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)

	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

func openLogFile(filePath string) *os.File {
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		mkDir()
	case os.IsPermission(err):
		log.Fatal("Permission: %v", err)
	}
	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to OpenFile: %v", err)
	}
	return handle
}

func mkDir() {
	dir, _ := os.Getwd()
	err := os.MkdirAll(path.Join(dir, getLogFilePath()), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
