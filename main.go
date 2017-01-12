package main

//go:generate go-bindata -o generator/tmpl/bindata.go -ignore bindata.go -pkg tmpl generator/tmpl/

import (
	"github.com/tukdesk/tuk-gen/cmd"
)

func main() {
	cmd.Execute()
}
