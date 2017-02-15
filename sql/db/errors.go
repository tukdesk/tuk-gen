package db

import (
	"database/sql"
)

func IsErrNoRows(err error) bool {
	return err == sql.ErrNoRows
}

func IsErrTxDone(err error) bool {
	return err == sql.ErrTxDone
}
