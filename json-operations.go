package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

const fileLocation = "./.hidden/lists.json"

func readJson() (elements []string) {
	fileBytes, err := ioutil.ReadFile(fileLocation)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(fileBytes, &elements)
	if err != nil {
		panic(err)
	}
	return elements
}

func writeJson() {
	addCMD := flag.NewFlagSet("add", flag.ExitOnError)
	data := addCMD.String("d", "", "enter the data")
	// data is not getting the value.
	err := addCMD.Parse(os.Args[2:])
	element := append(readJson(), *data)
	infoByte, err := json.Marshal(element)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(fileLocation, infoByte, 0644)
	if err != nil {
		panic(err)
	}
}

func list() {
	readData := readJson()
	for i, data := range readData {
		fmt.Printf("|%d|\t%s\t|\n", i+1, data)
	}
}
