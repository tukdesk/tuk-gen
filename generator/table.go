package generator

import (
	"github.com/tukdesk/base/sql/db"
	"github.com/tukdesk/tuk-gen/meta"
)

type Lines []string

func (this Lines) MaxIdx() int {
	return len(this) - 1
}

type Table struct {
	*meta.Object
	opt Option
}

func (this *Table) DefLines() Lines {
	engine := this.Object.Engine

	lines := make(Lines, 0, 1+len(this.Fields)+len(this.Indexes)) // primay key + fields + indexes

	switch engine {
	case db.MYSQL:
		lines = append(lines, ColumnDefForMySQL(this.opt, this.PrimaryKey, true))
		for _, f := range this.Fields {
			lines = append(lines, ColumnDefForMySQL(this.opt, f, false))
		}

		for _, i := range this.Indexes {
			lines = append(lines, IndexDefForMySQL(i))
		}

	case db.POSTGRESQL:

	}

	return lines
}

func (this *Table) ExtraLines() []string {
	lines := make([]string, 0)

	engine := this.Object.Engine

	switch engine {
	case db.MYSQL:
		lines = append(lines,
			"ENGINE = InnoDB",
			"DEFAULT CHARSET utf8",
		)

	case db.POSTGRESQL:

	}

	return lines
}
