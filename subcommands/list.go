package subcommands

import (
	"fmt"

	"github.com/JammUtkarsh/grec/utils"
	"github.com/pterm/pterm"
)

func List() {
	readData := utils.ReadJson()
	for _, data := range readData {
		pterm.DefaultBulletList.WithItems([]pterm.BulletListItem{
			{Level: 0, Text: data, TextStyle: pterm.NewStyle(pterm.FgWhite), BulletStyle: pterm.NewStyle(pterm.FgWhite)},
		}).Render()
	}
}
func Paste() {
	readData := utils.ReadJson()
	fmt.Println(readData[len(readData)-1])
}
