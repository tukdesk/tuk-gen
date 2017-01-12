package parser

import (
	"fmt"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestDefUnmarshal(t *testing.T) {
	engine := "postgres"
	db := "base"
	table := "table"

	data := fmt.Sprintf(`
Account:
  engine: %s
  db: %s
  table: %s
  fields:
    - Username: string
      length: 30
    - Password: string
      length: 64
    - Email: string
      length: 100
      nullable: true
    - CreatedAt: timestamp
    - UpdatedAt: timestamp
  uniques: [[Username]]
  indexes: [[CreatedAt], [UpdatedAt]]
  sparses: [[Email]]
    `, engine, db, table)

	out := map[string]Def{}

	err := yaml.Unmarshal([]byte(data), out)
	if err != nil {
		t.Fatal(err)
	}

	modelAccount, ok := out["Account"]
	if !ok {
		t.Fatal("no model named Account")
	}

	if modelAccount.Engine != engine {
		t.Errorf("expected engine %s, got %s", engine, modelAccount.Engine)
	}

	if modelAccount.DB != db {
		t.Errorf("expected db %s, got %s", db, modelAccount.DB)
	}

	if modelAccount.Table != table {
		t.Errorf("expected table %s, got %s", table, modelAccount.Table)
	}

	if len(modelAccount.Fileds) != 5 {
		t.Errorf("expected 5 fields, got %d", len(modelAccount.Fileds))
	}

	if len(modelAccount.Uniques) != 1 {
		t.Errorf("expected 1 unique index, got %d", len(modelAccount.Uniques))
	}

	if len(modelAccount.Indexes) != 2 {
		t.Errorf("expected 2 normal index, got %d", len(modelAccount.Indexes))
	}

	if len(modelAccount.Sparses) != 1 {
		t.Errorf("expected 1 sparse index, got %d", len(modelAccount.Sparses))
	}
}
