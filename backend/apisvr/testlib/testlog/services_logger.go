package testlog

import (
	"apisvr/lib/log"
	"apisvr/lib/time"
	"testing"
)

func NewServicesLogger(t *testing.T) *log.Logger {
	w := NewLogWriter(t)
	output := log.ConsoleWriter{Out: w, TimeFormat: time.RFC3339}
	logger := log.New(output).With().Timestamp().Logger()
	return &logger
}
