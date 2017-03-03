package parser

const (
	FiledOptionLength   = "length"
	FiledOptionNullable = "nullable"
	FiledOptionDefault  = "default"
	FiledOptionColumn   = "column"
)

type Def struct {
	Engine     string                   `yaml:"engine,omitempty"`
	DB         string                   `yaml:"db,omitempty"`
	Table      string                   `yaml:"table,omitempty"`
	PrimaryKey string                   `yaml:"pk,omitempty"`
	Fileds     []map[string]interface{} `yaml:"fields"`
	Indexes    [][]string               `yaml:"indexes,flow"`
	Uniques    [][]string               `yaml:"uniques,flow"`
	Sparses    [][]string               `yaml:"sparses,flow"`
	Partition  *Partition               `yaml:"partition,omitempty"`
}

type Partition struct {
	Type      string         `yaml:"type"`
	Linear    bool           `yaml:"linear"`
	Algorithm int            `yaml:"algorithm"`
	Expr      string         `yaml:"expr"`
	Columns   []string       `yaml:"columns"`
	Num       int            `yaml:"num"`
	Defs      []PartitionDef `yaml:"defs"`
}

type PartitionDef struct {
	Name     string   `yaml:"name"`
	Op       string   `yaml:"op"`
	Expr     string   `yaml:"expr"`
	Values   []string `yaml:"values"`
	MaxValue bool     `yaml:"max_value"`
}
