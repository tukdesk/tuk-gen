package db

import (
	"database/sql"
)

var (
	_ RowScanner = &sql.Row{}
	_ RowScanner = &sql.Rows{}
)

type RowScanner interface {
	Scan(...interface{}) error
}
