package types

import (
	"fmt"

	"database/sql"
	"database/sql/driver"
)

var (
	_id Id
	_   sql.Scanner   = &_id
	_   driver.Valuer = _id
)

type Id int64

func (this *Id) Scan(val interface{}) error {
	switch val := val.(type) {
	case int64:
		*this = Id(val)
		return nil
	}

	return fmt.Errorf("invalid value type %T", val)
}

func (this Id) Value() (driver.Value, error) {
	return int64(this), nil
}
