package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no commands provided")
		os.Exit(1)
	}
	cmd := os.Args[1]
	switch cmd {
	case "add":
		writeJson()
	case "list":
		list()
	case "clear":
		clearJson()
	case "delete":
		deleteObject()
	}
}
