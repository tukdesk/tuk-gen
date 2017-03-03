package parser

import (
	"fmt"

	"github.com/tukdesk/tuk-gen/meta"
	"github.com/tukdesk/tuk-gen/util"
)

func init() {
	if !defaultPrimaryKey.Valid() {
		panic(fmt.Errorf("invalid default primary key"))
	}

	RegisterDefParser("", DefaultDefParser)
}

const (
	defaultPrimaryKey = meta.FieldName("Id")
)

func DefaultDefParser(name string, def Def) (*meta.Object, error) {
	obj := &meta.Object{
		Name:     name,
		Engine:   def.Engine,
		DB:       def.DB,
		Table:    def.Table,
		FieldMap: map[string]meta.Field{},
	}

	if obj.Table == "" {
		obj.Table = util.Camel2Snake(name)
	}

	pks := make([]meta.Field, 0)

	if len(def.PrimaryKey) == 0 {
		pkName := meta.FieldName(def.PrimaryKeyName)
		if pkName == "" {
			pkName = defaultPrimaryKey
		}

		defaultPK := meta.Field{
			Name:     pkName,
			Type:     meta.FieldTypeId,
			AutoIncr: true,
			Column:   util.Camel2Snake(pkName.String()),
		}
		pks = append(pks, defaultPK)

		obj.FieldMap[pkName.String()] = defaultPK

	} else {
		for i, pk := range def.PrimaryKey {
			if len(pk) == 0 {
				return nil, fmt.Errorf("#%d empty primary key map", i)
			}

			f, err := defaultFieldParser(pk)
			if err != nil {
				return nil, fmt.Errorf("#%d %s", i, err)
			}

			if _, ok := obj.FieldMap[f.Name.String()]; ok {
				return nil, fmt.Errorf("#%d duplicate pk defination for %s", i, f.Name)
			}

			pks = append(pks, f)
			obj.FieldMap[f.Name.String()] = f
		}
	}
	obj.PrimaryKey = pks

	for i, field := range def.Fileds {
		if len(field) == 0 {
			return nil, fmt.Errorf("#%d empty field map", i)
		}

		f, err := defaultFieldParser(field)
		if err != nil {
			return nil, fmt.Errorf("#%d %s", i, err)
		}

		if _, ok := obj.FieldMap[f.Name.String()]; ok {
			return nil, fmt.Errorf("#%d duplicate field defination for %s", i, f.Name)
		}

		obj.Fields = append(obj.Fields, f)
		obj.FieldMap[f.Name.String()] = f
	}

	for i, uIdxFields := range def.Uniques {
		if len(uIdxFields) == 0 {
			return nil, fmt.Errorf("#%d empty unique index", i)
		}

		idx := meta.Index{
			IsUnique: true,
		}

		for _, fName := range uIdxFields {
			field, ok := obj.FieldMap[fName]
			if !ok {
				return nil, fmt.Errorf("#%d field %s in unique index not found", i, fName)
			}

			idx.Fields = append(idx.Fields, field)
		}

		obj.Indexes = append(obj.Indexes, idx)
	}

	for i, nIdxFields := range def.Indexes {
		if len(nIdxFields) == 0 {
			return nil, fmt.Errorf("#%d empty normal index", i)
		}

		idx := meta.Index{}

		for _, fName := range nIdxFields {
			field, ok := obj.FieldMap[fName]
			if !ok {
				return nil, fmt.Errorf("#%d field %s in normal index not found", i, fName)
			}

			idx.Fields = append(idx.Fields, field)
		}

		obj.Indexes = append(obj.Indexes, idx)
	}

	for i, sIdxFields := range def.Sparses {
		if len(sIdxFields) == 0 {
			return nil, fmt.Errorf("#%d empty sparse index", i)
		}

		idx := meta.Index{}

		for _, fName := range sIdxFields {
			field, ok := obj.FieldMap[fName]
			if !ok {
				return nil, fmt.Errorf("#%d field %s in sparse index not found", i, fName)
			}

			idx.Fields = append(idx.Fields, field)
		}

		obj.Indexes = append(obj.Indexes, idx)
	}

	if def.Partition != nil {
		partitionColumns := make([]string, len(def.Partition.Columns))
		for i, column := range def.Partition.Columns {
			field, ok := obj.FieldMap[column]
			if !ok {
				return nil, fmt.Errorf("partition column field %s not found", column)
			}

			partitionColumns[i] = util.Camel2Snake(field.Column)
		}

		partitionDefs := make([]meta.PartitionDef, 0, len(def.Partition.Defs))
		for _, pDef := range def.Partition.Defs {
			partitionDefs = append(partitionDefs, meta.PartitionDef{
				Name:     pDef.Name,
				Op:       pDef.Op,
				Expr:     pDef.Expr,
				Values:   pDef.Values,
				MaxValue: pDef.MaxValue,
			})
		}

		obj.Partition = meta.Partition{
			Type:      def.Partition.Type,
			Linear:    def.Partition.Linear,
			Algorithm: def.Partition.Algorithm,
			Expr:      def.Partition.Expr,
			Columns:   partitionColumns,
			Num:       def.Partition.Num,
			Defs:      partitionDefs,
		}
	}

	return obj, nil
}

func defaultFieldParser(field map[string]interface{}) (meta.Field, error) {
	f := meta.Field{}

	for key, val := range field {
		switch key {
		case FiledOptionLength:
			f.Length, _ = val.(int)

		case FiledOptionNullable:
			f.Nullable, _ = val.(bool)

		case FiledOptionColumn:
			f.Column, _ = val.(string)

		case FiledOptionDefault:
			f.DefaultValue = val

		case FiledOptionAutoIncr:
			f.AutoIncr, _ = val.(bool)

		default:
			if f.Name != "" {
				return f, fmt.Errorf("multi field name: [%s, %s]", f.Name, key)
			}

			fName := meta.FieldName(key)
			if !fName.Valid() {
				return f, fmt.Errorf("invaid field name %q", key)
			}

			typStr, _ := val.(string)
			fType := meta.FieldType(typStr)
			if !fType.Valid() {
				return f, fmt.Errorf("invalid field type %q", fType)
			}

			f.Name, f.Type = fName, fType
		}
	}

	if f.Column == "" {
		f.Column = util.Camel2Snake(f.Name.String())
	}

	return f, nil
}
