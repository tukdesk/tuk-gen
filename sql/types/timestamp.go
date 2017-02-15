package types

import (
	"fmt"
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
	return Timestamp(t.UnixNano())
}

type Timestamp int64

func (this *Timestamp) Scan(val interface{}) error {
	switch val := val.(type) {
	case int64:
		*this = Timestamp(val)
		return nil

	case time.Time:
		*this = Timestamp(val.UnixNano())
		return nil
	}

	return fmt.Errorf("invalid value type %T", val)
}

func (this Timestamp) Value() (driver.Value, error) {
	return int64(this), nil
}

func (this Timestamp) Time() time.Time {
	v := int64(this)
	return time.Unix(v/sec, v%sec)
}
