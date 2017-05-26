package db

import (
	"github.com/gocraft/dbr"
)

type Where struct {
	Where dbr.Builder
	Query string
	Conds []interface{}
}

func (w *Where) Builder() dbr.Builder {
	if w.Query != "" {
		return dbr.Expr(w.Query, w.Conds...)
	}

	return w.Where
}
