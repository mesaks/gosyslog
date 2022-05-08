package gosyslog

import (
	"testing"
	"time"
)

func TestNewLogger(t *testing.T) {
	log, e := New(LOG_WARNING|LOG_DAEMON, "gosyslog", 1000)

	if e != nil {
		t.FailNow()
	}

	defer log.Close()

	log.Write([]byte("Hello"))

	log.Write([]byte("Word"))

	time.Sleep(5 * time.Second)
}
