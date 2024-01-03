package goaext

import (
	"fmt"
	"os"
)

type Logger interface {
	Log(keyvals ...interface{}) error
}

func LoggerErrorHandlerFunc(logger Logger) func(error) error {
	return func(err error) error {
		if err != nil {
			if err := logger.Log("error", err); err != nil {
				fmt.Fprintf(os.Stderr, "failed to logger.Log %+v\n", err)
			}
		}
		return err
	}
}

func StderrErrorHandler(err error) error {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
	}
	return err
}
