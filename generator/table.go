package generator

import (
	"github.com/tukdesk/tuk-gen/meta"
	"github.com/tukdesk/tuk-gen/sql/db"
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
		lines = append(lines, ColumnDefForMySQL(this.opt, this.PrimaryKey))
		for _, f := range this.Fields {
			lines = append(lines, ColumnDefForMySQL(this.opt, f))
		}

		for _, i := range this.Indexes {
			lines = append(lines, IndexDefForMySQL(i))
		}

	case db.POSTGRESQL:
		lines = append(lines, ColumnDefForPostgres(this.opt, this.PrimaryKey))
		for _, f := range this.Fields {
			lines = append(lines, ColumnDefForPostgres(this.opt, f))
		}

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
		for _, i := range this.Object.Indexes {
			lines = append(lines, IndexDefForPostgres(this.Object.TableIdentifier(), i))
		}
	}

	return lines
}
