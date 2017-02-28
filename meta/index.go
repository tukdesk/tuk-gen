package meta

type Index struct {
	IsSparse bool
	IsUnique bool
	Fields   []IndexField
}

type IndexField struct {
	Field
	Desc bool
}
