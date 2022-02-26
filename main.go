package main

import (
	"fmt"
	"os"
)

func main() {
	if !FileExists(fileLocation) {
		_, err := os.Create(fileLocation)
		clear()
		if err != nil {
			panic(err)
		}
	}
	if len(os.Args) < 2 {
		fmt.Println("no commands provided")
		os.Exit(1)
	}
	cmd := os.Args[1]
	switch cmd {
	case "add":
		add()
	case "ls":
		list()
	case "clear":
		clear()
	case "delete":
		delete()
	case "help":
		help()
	default:
		help()
	}
}
