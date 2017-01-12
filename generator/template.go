package generator

import (
	"bytes"
	"go/format"
	"io"
	"os"
	"path/filepath"
	"text/template"

	"github.com/tukdesk/tuk-gen/generator/tmpl"
)

var tpl *template.Template

func init() {
	tpl = template.New("tuk-gen/tmpl")

	files := []string{
		"generator/tmpl/struct.tmpl",
		"generator/tmpl/sqlmgr.tmpl",
	}

	for _, filename := range files {
		data, err := tmpl.Asset(filename)
		if err != nil {
			panic(err)
		}

		if _, err = tpl.Parse(string(data)); err != nil {
			panic(err)
		}
	}
}

func outputFile(dir, filename, tplName string, data interface{}) error {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	path := filepath.Join(dir, filename)

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	buf := new(bytes.Buffer)
	if err := tpl.ExecuteTemplate(buf, tplName, data); err != nil {
		return err
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}

	_, err = io.Copy(file, bytes.NewReader(formatted))

	return err
}
