package generator

import (
	"fmt"
	"strings"

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

	lines := make(Lines, 0, len(this.PrimaryKey)+len(this.Fields)+len(this.Indexes)) // primay key + fields + indexes
	pkCols := make([]string, 0, len(this.PrimaryKey))

	switch engine {
	case db.MYSQL:
		for _, pk := range this.PrimaryKey {
			lines = append(lines, ColumnDefForMySQL(this.opt, pk))
			pkCols = append(pkCols, pk.Column)
		}

		for _, f := range this.Fields {
			lines = append(lines, ColumnDefForMySQL(this.opt, f))
		}

		for _, i := range this.Indexes {
			lines = append(lines, IndexDefForMySQL(i))
		}

	case db.POSTGRESQL:
		for _, pk := range this.PrimaryKey {
			lines = append(lines, ColumnDefForPostgres(this.opt, pk))
			pkCols = append(pkCols, pk.Column)
		}

		for _, f := range this.Fields {
			lines = append(lines, ColumnDefForPostgres(this.opt, f))
		}

	}

	lines = append(lines, fmt.Sprintf("PRIMARY KEY (%s)", strings.Join(pkCols, ", ")))

	return lines
}

func (this *Table) ExtraLines() []string {
	lines := make([]string, 0)

	engine := this.Object.Engine

	switch engine {
	case db.MYSQL:
		lines = append(lines,
			"ENGINE = InnoDB",
			"DEFAULT CHARSET utf8mb4",
		)

		if partitionLines := PartitionLinesForMySQL(this.Partition); len(partitionLines) > 0 {
			lines = append(lines, "PARTITION BY")
			lines = append(lines, partitionLines...)
		}

	case db.POSTGRESQL:
		for _, i := range this.Object.Indexes {
			lines = append(lines, IndexDefForPostgres(this.Object.TableIdentifier(), i))
		}
	}

	return lines
}
