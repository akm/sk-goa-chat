package testlog

import (
	"apisvr/services/gen/log"
	"testing"
	"time"

	"github.com/rs/zerolog"
)

func NewServicesLogger(t *testing.T) *log.Logger {
	w := NewLogWriter(t)
	output := zerolog.ConsoleWriter{Out: w, TimeFormat: time.RFC3339}
	logger := zerolog.New(output).With().Timestamp().Logger()
	return &log.Logger{Logger: &logger}
}
