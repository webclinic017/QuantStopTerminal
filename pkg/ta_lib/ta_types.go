package ta_lib

// MaType - Moving average type
type MaType int

type moneyFlow struct {
	positive float64
	negative float64
}

// Kinds of moving averages
const (
	SMA MaType = iota
	EMA
	WMA
	DEMA
	TEMA
	TRIMA
	KAMA
	MAMA
	T3MA
)
