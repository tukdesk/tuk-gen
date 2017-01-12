package meta

type Index struct {
	IsSparse bool
	IsUnique bool
	Fields   []Field
}
