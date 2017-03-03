package meta

import (
	"fmt"
	"strings"
)

type Object struct {
	Package   string
	Name      string
	Engine    string
	DB        string
	Table     string
	Indexes   []Index
	Partition Partition

	PrimaryKey Fields
	Fields     Fields
	FieldMap   map[string]Field
}

func (this *Object) Columns() []string {
	columns := make([]string, 0, len(this.PrimaryKey)+len(this.Fields))
	for _, pk := range this.PrimaryKey {
		columns = append(columns, pk.Column)
	}

	for _, one := range this.Fields {
		columns = append(columns, one.Column)
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
