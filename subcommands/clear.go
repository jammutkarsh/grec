package subcommands

import (
	"io/ioutil"

	"github.com/JammUtkarsh/grec/utils"
)

func Clear() {
	_ = ioutil.WriteFile(utils.GetDBFile(), []byte("[]"), 0755)
}
