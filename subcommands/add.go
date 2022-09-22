package subcommands

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/JammUtkarsh/grec/utils"
)

func Add() {
	addCMD := flag.NewFlagSet("add", flag.ExitOnError)
	data := addCMD.String("d", "", "Enter some text.")
	if len(os.Args) < 3 {
		fmt.Print(`Options:	
	-d [value]
	-d "multiple values"
`)
		os.Exit(0)
	}
	err := addCMD.Parse(os.Args[2:])
	if err != nil {
		log.Println(err)
	}
	element := append(utils.ReadJson(), *data)
	infoByte, err := json.Marshal(element)
	if err != nil {
		log.Println(err)
	}
	err = ioutil.WriteFile(utils.GetFilePath(), infoByte, 0644)
	if err != nil {
		log.Println(err)
	}
}
