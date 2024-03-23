package log

import (
	"github.com/rs/zerolog"
)

type (
	Logger        = zerolog.Logger
	ConsoleWriter = zerolog.ConsoleWriter

	Level = zerolog.Level
)

var (
	New = zerolog.New

	GlobalLevel    = zerolog.GlobalLevel
	SetGlobalLevel = zerolog.SetGlobalLevel
)

const (
	DebugLevel = zerolog.DebugLevel
	InfoLevel  = zerolog.InfoLevel
	WarnLevel  = zerolog.WarnLevel
	ErrorLevel = zerolog.ErrorLevel
	FatalLevel = zerolog.FatalLevel
	PanicLevel = zerolog.PanicLevel
	NoLevel    = zerolog.NoLevel
	Disabled   = zerolog.Disabled
	TraceLevel = zerolog.TraceLevel
)
