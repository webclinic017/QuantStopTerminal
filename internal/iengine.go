package internal

import "database/sql"

type IEngine interface {
	GetUptime() string
	GetSubsystemsStatus() map[string]bool
	SetSubsystem(subSystemName string, enable bool) error
	GetVersion() map[string]string
	GetSQL() (*sql.DB, error)
}
