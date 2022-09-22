package subcommands

import (
	"io/ioutil"

	"github.com/JammUtkarsh/grec/utils"
)

func Clear() {
	clear := []byte("[]")
	_ = ioutil.WriteFile(utils.GetFilePath(), clear, 0755)
}
