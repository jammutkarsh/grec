package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func PrettyJSON(v interface{}) (prettyJson string, err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func ReadJson() (elements []string) {
	fileBytes, err := ioutil.ReadFile(GetFilePath())
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(fileBytes, &elements)
	if err != nil {
		log.Println(err)
	}
	return elements
}
