package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/JammUtkarsh/grec/subcommands"
	"github.com/JammUtkarsh/grec/utils"
)

func init() {
	utils.CreateConfigDir()
	if utils.FileExists(utils.GetConfigPath() + "db.json") {
		utils.CreateFile("db.json")
		_ = ioutil.WriteFile(utils.GetDBFile(), []byte("[]"), 0755)
	}
	if utils.FileExists(utils.GetConfigPath() + "logs.txt") {
		utils.CreateFile("logs.txt")
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no commands provided")
		subcommands.Help()
		os.Exit(1)
	}
	cmd := os.Args[1]
	switch cmd {
	case "add":
		subcommands.Add()
	case "ls":
		subcommands.List()
	case "paste":
		subcommands.Paste()
	case "clear":
		subcommands.Clear()
	case "delete":
		subcommands.Delete()
	case "help":
		subcommands.Help()
	default:
		subcommands.Help()
	}
}
