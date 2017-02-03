package generator

import (
	"fmt"
	"strings"

	"github.com/tukdesk/tuk-gen/meta"
)

const postgresIndexFormat = `;

CREATE%s INDEX %s IF NOT EXISTS
ON %s
(%s)
`

func ColumnDefForPostgres(opt Option, field meta.Field) string {
	str := field.Column

	length := field.Length

	switch field.Type {
	case meta.FieldTypeString:
		str += " VARCHAR"

		if length == 0 {
			length = opt.DefaultStringLength
		}

		if length <= 0 {
			length = 1
		}

		str += fmt.Sprintf("(%d)", length)

	case meta.FieldTypeInt8, meta.FieldTypeInt16:
		str += " SMALLINT"

	case meta.FieldTypeInt32:
		str += " INT"

	case meta.FieldTypeInt64:
		str += " BIGINT"

	case meta.FieldTypeId:
		str += " BIGSERIAL"

	case meta.FieldTypeFloat64:
		str += " DOUBLE PRECISION"

	case meta.FieldTypeBool:
		str += " BOOLEAN"

	case meta.FieldTypeEnum:
		str += " INT"

	case meta.FieldTypeTimestamp:
		str += " BIGINT"

	case meta.FieldTypeDatetime:
		str += " TIMESTAMP"

	case meta.FieldTypeBytes:
		str += " BYTEA"

	default:
		panic(fmt.Errorf("unsupport data type %s of field %s for postgres", field.Type, field.Column))
	}

	if field.IsPrimaryKey {
		str += " PRIMARY KEY"
	}

	if !field.IsPrimaryKey && !field.Nullable {
		str += " NOT NULL"
	}

	return str
}

func IndexDefForPostgres(table string, index meta.Index) string {
	names := make([]string, len(index.Fields))
	for i, f := range index.Fields {
		names[i] = f.Column
	}

	unique := ""
	if index.IsUnique {
		unique = " UNIQUE"
	}

	idxName := fmt.Sprintf("ix_%s", strings.Join(names, "_"))

	columns := strings.Join(names, ", ")

	return fmt.Sprintf(postgresIndexFormat, unique, idxName, table, columns)
}
