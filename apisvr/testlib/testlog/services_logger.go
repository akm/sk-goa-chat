package testlog

import (
	"apisvr/lib/log"
	"testing"
	"time"
)

func NewServicesLogger(t *testing.T) *log.Logger {
	w := NewLogWriter(t)
	output := log.ConsoleWriter{Out: w, TimeFormat: time.RFC3339}
	logger := log.New(output).With().Timestamp().Logger()
	return &logger
}
