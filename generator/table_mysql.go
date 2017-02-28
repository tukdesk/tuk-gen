package generator

import (
	"fmt"
	"strings"

	"github.com/tukdesk/tuk-gen/meta"
)

func ColumnDefForMySQL(opt Option, field meta.Field) string {
	str := field.Column

	length := field.Length
	lengthRequired := false
	lengthAvailable := true

	unsignedNumeric := false

	switch field.Type {
	case meta.FieldTypeUInt8:
		str += " TINYINT"
		unsignedNumeric = true

	case meta.FieldTypeUInt16:
		str += " SMALLINT"
		unsignedNumeric = true

	case meta.FieldTypeUInt32:
		str += " INT"
		unsignedNumeric = true

	case meta.FieldTypeUInt64:
		str += " BIGINT"
		unsignedNumeric = true

	case meta.FieldTypeString:
		str += " VARCHAR"
		lengthRequired = true

	case meta.FieldTypeInt8:
		str += " TINYINT"

	case meta.FieldTypeInt16:
		str += " SMALLINT"

	case meta.FieldTypeInt32:
		str += " INT"

	case meta.FieldTypeInt64, meta.FieldTypeId:
		str += " BIGINT"

	case meta.FieldTypeFloat64:
		str += " DOUBLE"

	case meta.FieldTypeBool:
		str += " BOOLEAN"
		lengthAvailable = false

	case meta.FieldTypeEnum:
		str += " INT"

	case meta.FieldTypeTimestamp:
		str += " BIGINT"

	case meta.FieldTypeDatetime:
		str += " DATETIME"

	case meta.FieldTypeBytes:
		str += " BLOB"

	case meta.FieldTypeText:
		str += " TEXT"
		lengthAvailable = false

	default:
		panic(fmt.Errorf("unsupport data type %s of field %s for mysql", field.Type, field.Column))
	}

	if lengthAvailable {
		if lengthRequired && length == 0 && opt.DefaultStringLength > 0 {
			length = opt.DefaultStringLength
			if length <= 0 {
				length = 1
			}
		}

		if length > 0 {
			str += fmt.Sprintf(" (%d) ", length)
		}
	}

	if unsignedNumeric {
		str += " UNSIGNED"
	}

	if field.IsPrimaryKey {
		if field.Type.AutoIncrable() && field.AutoIncr {
			str += " AUTO_INCREMENT"
		}

		str += " PRIMARY KEY"
	}

	if !field.IsPrimaryKey && !field.Nullable {
		str += " NOT NULL"
		if field.Type.HasDefault() && field.DefaultValue != nil {
			str += " DEFAULT " + field.Default()
		}
	}

	return str
}

func IndexDefForMySQL(index meta.Index) string {
	names := make([]string, len(index.Fields))
	idxDefs := make([]string, len(index.Fields))
	for i, f := range index.Fields {
		names[i] = f.Column

		c := f.Column
		if f.Desc {
			c += " DESC"
		}
		idxDefs[i] = c
	}

	var typStr string
	if index.IsUnique {
		typStr = "UNIQUE KEY"
	} else {
		typStr = "INDEX"
	}

	idxName := fmt.Sprintf("ix_%s", strings.Join(names, "_"))

	columns := fmt.Sprintf("(%s)", strings.Join(idxDefs, ", "))
	return fmt.Sprintf("%s %s %s", typStr, idxName, columns)
}
