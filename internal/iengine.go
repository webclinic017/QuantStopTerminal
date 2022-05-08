package internal

import (
	"database/sql"
	"github.com/quantstop/qsx/core"
)

type IEngine interface {
	GetUptime() string
	GetSubsystemsStatus() map[string]bool
	SetSubsystem(subSystemName string, enable bool) error
	GetVersion() map[string]string
	GetSQL() (*sql.DB, error)
	SetSystemConfig(string, string) error
	GetExchange(string) core.Qsx
	GetSupportedExchangesList() []string
}
