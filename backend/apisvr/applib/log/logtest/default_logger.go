package logtest

import (
	"log"
	"testing"
)

func NewDefaultLogger(t *testing.T) *log.Logger {
	w := NewLogWriter(t)
	return log.New(w, "TEST ", log.LstdFlags)
}
