package utils

import (
	"io/ioutil"
	"os"
)

func FileExists(filename string) {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		os.NewFile(0, filename)
		clear := []byte("[]")
		_ = ioutil.WriteFile(filename, clear, 0755)
	}
}

func GetConfigPath() string {
	home, _ := os.UserHomeDir()
	return home + "/.grec/"
}

func GetFilePath() string {
	return GetConfigPath() + "list.json"
}
