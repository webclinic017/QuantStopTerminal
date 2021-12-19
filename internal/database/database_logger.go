package database

import (
	"github.com/quantstop/quantstopterminal/pkg/logger"
)

// Logger implements io.Writer interface to redirect SQLBoiler debug output to GCT logger
type Logger struct{}

// Write takes input and sends to GCT logger
func (l Logger) Write(p []byte) (n int, err error) {
	logger.Debugf(logger.DatabaseLogger, "SQL: %s", p)
	return 0, nil
}
