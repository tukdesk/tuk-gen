package types

import (
	"fmt"
	"strconv"

	"database/sql"
	"database/sql/driver"
)

var (
	_enum Enum
	_     sql.Scanner   = &_enum
	_     driver.Valuer = _enum
)

type Enum int32

func (this Enum) Int32() int32 {
	return int32(this)
}

func (this *Enum) Scan(val interface{}) error {
	switch val := val.(type) {
	case int64:
		*this = Enum(val)
		return nil

	case []byte:
		i, err := strconv.Atoi(string(val))
		if err != nil {
			return fmt.Errorf("fail to conv %s to int", string(val))
		}

		*this = Enum(i)
	}

	return fmt.Errorf("invalid value type %T", val)
}

func (this Enum) Value() (driver.Value, error) {
	return int32(this), nil
}
