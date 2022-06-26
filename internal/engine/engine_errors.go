package engine

import "errors"

var (
	ErrNilCoreSQL         = errors.New("engine: GetCoreSQL cannot return nil database")
	ErrNilCoinbaseSQL     = errors.New("engine: GetCoinbaseSQL cannot return nil database")
	ErrNilTDAmeritradeSQL = errors.New("engine: GetTDAmeritradeSQL cannot return nil database")
)
