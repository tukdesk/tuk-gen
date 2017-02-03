package generator

import (
	"bytes"
	"fmt"
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
		"generator/tmpl/sqlconn.tmpl",
		"generator/tmpl/sql_script_table.tmpl",
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

func outputNormalFile(dir, filename, tplName string, data interface{}) error {
	buf := new(bytes.Buffer)
	if err := tpl.ExecuteTemplate(buf, tplName, data); err != nil {
		return err
	}

	return writeToFile(buf, dir, filename)
}

func outputGoFile(dir, filename, tplName string, data interface{}) error {
	buf := new(bytes.Buffer)
	if err := tpl.ExecuteTemplate(buf, tplName, data); err != nil {
		return err
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Println(buf.String())
		return err
	}

	return writeToFile(bytes.NewReader(formatted), dir, filename)
}

func writeToFile(r io.Reader, dir, filename string) error {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	path := filepath.Join(dir, filename)

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, r)

	return err
}
