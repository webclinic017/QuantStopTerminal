package database

import (
	"github.com/quantstop/quantstopterminal/internal/qstlog"
)

// Logger implements io.Writer interface to redirect SQLBoiler debug output to GCT logger
type Logger struct{}

// Write takes input and sends to GCT logger
func (l Logger) Write(p []byte) (n int, err error) {
	qstlog.Debugf(qstlog.DatabaseLogger, "SQL: %s", p)
	return 0, nil
}
