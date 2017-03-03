package meta

import (
	"fmt"
	"strings"
)

type Object struct {
	Package    string
	Name       string
	Engine     string
	DB         string
	Table      string
	PrimaryKey Field
	Indexes    []Index
	Partition  Partition

	Fields   []Field
	FieldMap map[string]Field
}

func (this *Object) Columns() []string {
	columns := make([]string, len(this.Fields)+1)
	columns[0] = this.PrimaryKey.Column

	for i, one := range this.Fields {
		columns[i+1] = one.Column
	}

	return columns
}

func (this *Object) ColumnsJoin(sep string) string {
	colunms := this.Columns()

	return strings.Join(colunms, sep)
}

func (this *Object) TableIdentifier() string {
	if this.DB == "" {
		return this.Table
	}

	return fmt.Sprintf("%s.%s", this.DB, this.Table)
}
