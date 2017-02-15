package types

import (
	"fmt"
	"strconv"

	"database/sql"
	"database/sql/driver"
)

var (
	_id Id
	_   sql.Scanner   = &_id
	_   driver.Valuer = _id
)

type Id int64

func (this Id) Int64() int64 {
	return int64(this)
}

func (this *Id) Scan(val interface{}) error {
	switch val := val.(type) {
	case int64:
		*this = Id(val)
		return nil

	case []byte:
		i, err := strconv.Atoi(string(val))
		if err != nil {
			return fmt.Errorf("fail to conv %s to int", string(val))
		}

		*this = Id(i)
	}

	return fmt.Errorf("invalid value type %T", val)
}

func (this Id) Value() (driver.Value, error) {
	return int64(this), nil
}
