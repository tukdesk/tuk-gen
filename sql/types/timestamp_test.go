package types

import (
	"testing"
	"time"
)

func TestTimestamp(t *testing.T) {
	for i := 0; i < 10; i++ {
		now := time.Now()
		ts := Timestamp(now.Unix())

		if got := ts.Time(); got.Unix() != now.Unix() {
			t.Errorf("expected %s, got %s", now, got)
		}

		time.Sleep(time.Millisecond)
	}
}
