package subcommands

import (
	"fmt"

	"github.com/JammUtkarsh/grec/utils"
)

func List() {
	readData := utils.ReadJson()
	for i, data := range readData {
		fmt.Printf("|%d|\t%s\t|\n", i+1, data)
	}
}
func Paste() {
	readData := utils.ReadJson()
	for _, data := range readData {
		fmt.Printf("%s", data)
	}
}
