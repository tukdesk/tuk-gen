package types

import (
	"fmt"
	"strconv"
	"time"

	"database/sql"
	"database/sql/driver"
)

var (
	sec = int64(time.Second)

	_timestamp Timestamp
	_          sql.Scanner   = &_timestamp
	_          driver.Valuer = _timestamp
)

func NewTimestamp(t time.Time) Timestamp {
	return Timestamp(t.Unix())
}

type Timestamp int64

func (this *Timestamp) Scan(val interface{}) error {
	if val == nil {
		*this = 0
		return nil
	}

	switch val := val.(type) {
	case int64:
		*this = Timestamp(val)
		return nil

	case time.Time:
		*this = Timestamp(val.Unix())
		return nil

	case []byte:
		i, err := strconv.Atoi(string(val))
		if err != nil {
			return fmt.Errorf("fail to conv %s to int", string(val))
		}

		*this = Timestamp(i)
		return nil

	case nil:
		*this = 0
		return nil
	}

	return fmt.Errorf("invalid value type %T", val)
}

func (this Timestamp) Value() (driver.Value, error) {
	return int64(this), nil
}

func (this Timestamp) Time() time.Time {
	v := int64(this)
	return time.Unix(v, 0)
}

func (this Timestamp) Int64() int64 {
	return int64(this)
}
