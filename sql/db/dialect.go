package db

import (
	"github.com/gocraft/dbr"
)

var dialects = map[string]dbr.Dialect{}

func RegisterDialect(engine string, d dbr.Dialect) {
	dialects[engine] = d
}
