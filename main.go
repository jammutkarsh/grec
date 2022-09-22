package main

import (
	"fmt"
	"log"
	"os"

	"github.com/JammUtkarsh/grec/subcommands"
	"github.com/JammUtkarsh/grec/utils"
)

func init() {
	utils.FileExists(utils.GetFilePath())
	utils.FileExists(utils.GetConfigPath() + "logs.txt")
	f, err := os.OpenFile(utils.GetConfigPath()+"logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	log.SetOutput(f)
	if err != nil {
		fmt.Println(err)
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
