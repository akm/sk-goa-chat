package testlog

import (
	"log"
	"testing"
)

func New(t *testing.T) *log.Logger {
	w := NewLogWriter(t)
	return log.New(w, "TEST ", log.LstdFlags)
}
