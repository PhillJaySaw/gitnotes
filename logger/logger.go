package logger

import (
	"os"

	"github.com/charmbracelet/log"
)

func New() *log.Logger {
	logger := log.NewWithOptions(os.Stderr, log.Options{
		ReportTimestamp: false,
	})

	return logger
}
