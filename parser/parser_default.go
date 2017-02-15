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

	pk := meta.FieldName(def.PrimaryKey)
	if pk == "" {
		pk = defaultPrimaryKey
	}

	obj.PrimaryKey = meta.Field{
		Name:         pk,
		Type:         meta.FieldTypeId,
		IsPrimaryKey: true,
		Column:       util.Camel2Snake(pk.String()),
	}

	obj.FieldMap[obj.PrimaryKey.Name.String()] = obj.PrimaryKey

	for i, field := range def.Fileds {
		if len(field) == 0 {
			return nil, fmt.Errorf("#%d empty field map")
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
