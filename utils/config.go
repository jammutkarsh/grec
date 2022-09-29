package utils

import (
	"fmt"
	"os"
)

func GetConfigPath() string {
	home, _ := os.UserHomeDir()
	return home + "/.grec/"
}

func CreateConfigDir() {
	path := GetConfigPath()
	_, err := os.Stat(path)
	if !os.IsNotExist(err) {
		return
	} else {
		err = os.Mkdir(path, 0755)
		if err != nil {
			_ = fmt.Errorf("error creating config directory: %v", err)
			os.Exit(1)
		}
	}
}

func CreateFile(filename string) {
	path := GetConfigPath()
	_, err := os.OpenFile(path+filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		_ = fmt.Errorf("error creating file: %v", err)
		os.Exit(1)
	}
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return os.IsNotExist(err)
}

func GetDBFile() string {
	return GetConfigPath() + "db.json"
}
