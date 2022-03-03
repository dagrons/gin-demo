package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func GetAppPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))

	return path[:index]
}

func GetEnvString(key string, defaultValue string) string {
	val := os.Getenv(key)
	if len(val) == 0 {
		return defaultValue
	}
	return val
}

func GetEnvInt(key string, defaultValue int) int {
	str := os.Getenv(key)
	val, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		return int(val)
	}
	return defaultValue
}

func MustGetEnvInt(key string) int {
	str := os.Getenv(key)
	val, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		panic("failed to get env int")
	}
	return int(val)
}

func MustGetEnvString(key string) string {
	val := os.Getenv(key)
	if len(val) == 0 {
		panic(fmt.Sprintf("%s cannot be empty", key))
	}
	return val
}
