package logtest

import (
	"testing"
)

type LogWriter struct {
	t *testing.T
}

func NewLogWriter(t *testing.T) *LogWriter {
	return &LogWriter{t: t}
}

func (w *LogWriter) Write(p []byte) (n int, err error) {
	w.t.Logf("%s", p)
	return 0, nil
}
