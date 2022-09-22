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

func Delete() {
	delCMD := flag.NewFlagSet("delete", flag.ExitOnError)
	data := delCMD.Int("d", 0, "delete an entry")
	if len(os.Args) < 3 {
		fmt.Print(`Options:	
	-d [value]
`)
		os.Exit(0)
	}
	err := delCMD.Parse(os.Args[2:])
	if err != nil {
		log.Println(err)
	}
	jsonObjects := utils.ReadJson()
	var newJsonObjects []string
	for i, entry := range jsonObjects {
		if *data-1 == i {
			continue
		}
		newJsonObjects = append(newJsonObjects, entry)
	}
	infoByte, err := json.Marshal(newJsonObjects)
	if err != nil {
		log.Println(err)
	}
	err = ioutil.WriteFile(utils.GetFilePath(), infoByte, 0644)
	if err != nil {
		log.Println(err)
	}
}
