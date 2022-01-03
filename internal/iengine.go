package internal

type IEngine interface {
	GetUptime() string
	GetSubsystemsStatus() map[string]bool
	SetSubsystem(subSystemName string, enable bool) error
	GetVersionString(short bool) string
}
