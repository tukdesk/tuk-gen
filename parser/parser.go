package parser

import (
	"io/ioutil"
	"sort"

	"gopkg.in/yaml.v2"

	"github.com/tukdesk/tuk-gen/meta"
)

var parsers = map[string]DefParser{}

type DefParser func(name string, def Def) (*meta.Object, error)

func RegisterDefParser(engine string, parser DefParser) {
	parsers[engine] = parser
}

type SortedObjects []*meta.Object

func (this SortedObjects) Len() int {
	return len(this)
}

func (this SortedObjects) Less(i, j int) bool {
	return this[i].Table < this[j].Table
}

func (this SortedObjects) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func Parse(filename, packageName string) ([]*meta.Object, error) {
	in, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	out := map[string]Def{}
	if err := yaml.Unmarshal(in, out); err != nil {
		return nil, err
	}

	objects := make([]*meta.Object, 0, len(out))
	for name, def := range out {
		parser := parsers[name]
		if parser == nil {
			parser = DefaultDefParser
		}

		object, err := parser(name, def)
		if err != nil {
			return nil, err
		}

		object.Package = packageName

		objects = append(objects, object)
	}

	sort.Sort(SortedObjects(objects))

	return objects, nil
}
