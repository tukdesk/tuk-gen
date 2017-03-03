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
		str += " AUTO_INCREMENT PRIMARY KEY"
	}

	if !field.IsPrimaryKey && !field.Nullable {
		str += " NOT NULL"
	}

	return str
}

func IndexDefForMySQL(index meta.Index) string {
	names := make([]string, len(index.Fields))
	for i, f := range index.Fields {
		names[i] = f.Column
	}

	var typStr string
	if index.IsUnique {
		typStr = "UNIQUE KEY"
	} else {
		typStr = "INDEX"
	}

	idxName := fmt.Sprintf("ix_%s", strings.Join(names, "_"))

	columns := fmt.Sprintf("(%s)", strings.Join(names, ", "))
	return fmt.Sprintf("%s %s %s", typStr, idxName, columns)
}

func PartitionLinesForMySQL(partition meta.Partition) []string {
	lines := make([]string, 0)

	partitionType := strings.ToUpper(strings.TrimSpace(partition.Type))
	var defValid bool

	switch partitionType {
	case "HASH":
		if partition.Expr == "" {
			return nil
		}

		var linearPart string
		if partition.Linear {
			linearPart = "LINEAR "
		}

		lines = append(lines,
			fmt.Sprintf("\t%s%s (%s)", linearPart, partitionType, partition.Expr),
		)

		if partition.Num > 0 {
			lines = append(lines,
				fmt.Sprintf("PARTITIONS %d", partition.Num),
			)
		}

	case "KEY":
		if len(partition.Columns) == 0 {
			return nil
		}

		var linearPart string
		if partition.Linear {
			linearPart = "LINEAR "
		}

		var algorithmPart string
		if partition.Algorithm > 0 {
			algorithmPart = fmt.Sprintf("ALGORITHM=%d ", partition.Algorithm)
		}

		lines = append(lines,
			fmt.Sprintf("\t%s%s %s(%s)", linearPart, partitionType, algorithmPart, strings.Join(partition.Columns, ", ")),
		)

		if partition.Num > 0 {
			lines = append(lines,
				fmt.Sprintf("PARTITIONS %d", partition.Num),
			)
		}

	case "RANGE", "LIST":
		if partition.Expr == "" && len(partition.Columns) == 0 {
			return nil
		}

		defValid = true

		if partition.Expr != "" {
			lines = append(lines,
				fmt.Sprintf("\t%s (%s)", partitionType, partition.Expr),
			)
		} else {
			lines = append(lines,
				fmt.Sprintf("\t%s COLUMNS (%s)", partitionType, strings.Join(partition.Columns, ", ")),
			)
		}

	default:
		return nil
	}

	if defValid && len(partition.Defs) > 0 {
		lines = append(lines, "(")
		maxPDef := len(partition.Defs) - 1
		for i, one := range partition.Defs {
			var op string
			var valPart string
			if strings.ToUpper(one.Op) == "IN" {
				op = "IN"
				valPart = fmt.Sprintf("(%s)", strings.Join(one.Values, ", "))
			} else {
				op = "LESS THAN"
				if one.MaxValue {
					valPart = "MAXVALUE"
				} else if one.Expr != "" {
					valPart = fmt.Sprintf("(%s)", one.Expr)
				} else {
					valPart = fmt.Sprintf("(%s)", strings.Join(one.Values, ", "))
				}
			}

			defLine := fmt.Sprintf("PARTITION %s VALUES %s %s", one.Name, op, valPart)
			if i < maxPDef {
				defLine += ","
			}
			lines = append(lines, defLine)
		}
		lines = append(lines, ")")
	}

	return lines
}
