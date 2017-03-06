package meta

import (
	"strings"
)

type Index struct {
	IsSparse bool
	IsUnique bool
	Fields   IndexFields
}

type IndexField struct {
	Field
	Desc bool
}

type IndexFields []IndexField

func (this IndexFields) FuncTypedArgs() string {
	args := make([]string, len(this))
	for i, one := range this {
		args[i] = one.Name.String() + " " + one.Type.GoType()
	}

	return strings.Join(args, ", ")
}

func (this IndexFields) FuncArgs() string {
	args := make([]string, len(this))
	for i, one := range this {
		args[i] = one.Name.String()
	}

	return strings.Join(args, ", ")
}
