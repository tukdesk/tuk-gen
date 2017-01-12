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
}
