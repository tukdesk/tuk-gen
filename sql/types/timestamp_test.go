package types

import (
	"testing"
	"time"
)

func TestTimestamp(t *testing.T) {
	for i := 0; i < 10; i++ {
		now := time.Now()
		ts := Timestamp(now.UnixNano())

		if got := ts.Time(); got != now {
			t.Errorf("expected %s, got %s", now, got)
		}

		time.Sleep(time.Millisecond)
	}
}
