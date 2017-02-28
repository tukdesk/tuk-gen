package meta

type Index struct {
	IsSparse bool
	IsUnique bool
	IsDesc   bool
	Fields   []Field
}
