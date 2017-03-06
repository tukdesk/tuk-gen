package meta

type Partition struct {
	Type      string
	Linear    bool
	Algorithm int
	Expr      string
	Columns   []string
	Num       int
	Defs      []PartitionDef
}

type PartitionDef struct {
	Name     string
	Op       string
	Expr     string
	Values   []string
	MaxValue bool
}
