package internal

import (
	"database/sql"
	"github.com/quantstop/qsx/core"
)

type IEngine interface {
	GetUptime() string
	//GetCoreConfig() map[string]string
	SetConfig(string, string) error
	GetSubsystemsStatus() map[string]bool
	SetSubsystem(subSystemName string, enable bool) error
	GetVersion() map[string]string
	GetCoreSQL() (*sql.DB, error)
	GetCoinbaseSQL() (*sql.DB, error)
	GetTDAmeritradeSQL() (*sql.DB, error)

	GetExchange(string) core.Qsx
	GetSupportedExchangesList() []string
}
