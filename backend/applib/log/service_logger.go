package log

import (
	"applib/time"
	"io"
	"os"
	"regexp"

	"github.com/rs/zerolog"
)

func NewServiceLogger(w io.Writer, serviceName string) *Logger {
	var output io.Writer
	useConsoleWriter := regexp.MustCompile(`(?i)^(?:1|true|on|yes)$`).MatchString(os.Getenv("LOG_CONSOLE_WRITER"))
	if useConsoleWriter {
		output = zerolog.ConsoleWriter{Out: w, TimeFormat: time.RFC3339}
	} else {
		output = w
	}
	logger := zerolog.New(output).With().Timestamp().Str("service", serviceName).Logger()
	return &logger
}
