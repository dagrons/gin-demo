package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Level int

var (
	F *os.File

	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	logger     *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func init() {
	filePath := getLogFileFullPath()
	F = openLogFile(filePath)

	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

func Debug(format string, v ...interface{}) {
	setPrefix(DEBUG)
	logger.Printf(format, v)
}

func Info(format string, v ...interface{}) {
	setPrefix(INFO)
	logger.Printf(format, v)
}

func Warn(format string, v ...interface{}) {
	setPrefix(WARNING)
	logger.Printf(format, v)
}

func Error(format string, v ...interface{}) {
	setPrefix(ERROR)
	logger.Printf(format, v)
}

func Fatal(format string, v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalf(format, v)
}

func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}
