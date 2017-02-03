package meta

import (
	"regexp"
)

var reFieldName = regexp.MustCompile("^[a-zA-Z0-9_]+$")

type FieldName string

func (this FieldName) String() string {
	return string(this)
}

func (this FieldName) Valid() bool {
	return reFieldName.MatchString(string(this))
}

type FieldType string

func (this FieldType) Valid() bool {
	for _, one := range availableFieldTypes {
		if one == this {
			return true
		}
	}

	return false
}

func (this FieldType) GoType() string {
	switch this {
	case FieldTypeId:
		return "types.Id"

	case FieldTypeEnum:
		return "types.Enum"

	case FieldTypeTimestamp:
		return "types.Timestamp"

	case FieldTypeDatetime:
		return "time.Time"

	case FieldTypeBytes:
		return "[]byte"

	default:
		return string(this)
	}
}

func (this FieldType) GoZero() string {
	switch this {
	case FieldTypeString:
		return "\"\""

	case FieldTypeFloat64:
		return "0.0"

	case FieldTypeBool:
		return "false"

	case FieldTypeDatetime:
		return "time.Time{}"

	case FieldTypeBytes:
		return "nil"

	default:
		return "0"
	}
}

const (
	FieldTypeId        FieldType = "id"
	FieldTypeString              = "string"
	FieldTypeInt8                = "int8"
	FieldTypeInt16               = "int16"
	FieldTypeInt32               = "int32"
	FieldTypeInt64               = "int64"
	FieldTypeUInt8               = "uint8"
	FieldTypeUInt16              = "uint16"
	FieldTypeUInt32              = "uint32"
	FieldTypeUInt64              = "uint64"
	FieldTypeFloat64             = "float64"
	FieldTypeBool                = "bool"
	FieldTypeEnum                = "enum"
	FieldTypeTimestamp           = "timestamp"
	FieldTypeDatetime            = "datetime"
	FieldTypeBytes               = "bytes"
)

var availableFieldTypes = []FieldType{
	FieldTypeId,
	FieldTypeString,
	FieldTypeInt8,
	FieldTypeInt16,
	FieldTypeInt32,
	FieldTypeInt64,
	FieldTypeFloat64,
	FieldTypeBool,
	FieldTypeEnum,
	FieldTypeTimestamp,
	FieldTypeDatetime,
	FieldTypeBytes,
}

type Field struct {
	Name         FieldName
	Type         FieldType
	Column       string
	Length       int
	Nullable     bool
	IsPrimaryKey bool
	DefaultValue interface{}
}
