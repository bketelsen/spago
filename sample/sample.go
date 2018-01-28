package sample

import (
	"github.com/gobuffalo/packr"
)

const templatePrefix = "templates/"

var pack packr.Box

func init() {
	pack = packr.NewBox("./assets")
}

// Templates returns a template
func Template(template string) []byte {
	return pack.Bytes(templates + name)
}
