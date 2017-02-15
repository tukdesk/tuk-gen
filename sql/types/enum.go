package types

import (
	"fmt"

	"database/sql"
	"database/sql/driver"
)

var (
	_enum Enum
	_     sql.Scanner   = &_enum
	_     driver.Valuer = _enum
)

type Enum int32

func (this *Enum) Scan(val interface{}) error {
	switch val := val.(type) {
	case int64:
		*this = Enum(val)
		return nil
	}

	return fmt.Errorf("invalid value type %T", val)
}

func (this Enum) Value() (driver.Value, error) {
	return int32(this), nil
}
