package ntpmonitor

import (
	"errors"
	"time"
)

const (
	DefaultNTPCheckInterval = time.Second * 30
	DefaultRetryLimit       = 3
)

var (
	ErrNilNTPConfigValues   = errors.New("nil allowed time differences received")
	ErrNTPSubsystemDisabled = errors.New("NTP subsystem disabled")
)

type Config struct {
	Enabled                   bool
	Verbose                   bool
	Level                     int
	Pool                      []string
	AllowedDifference         *time.Duration
	AllowedNegativeDifference *time.Duration
}

type NTPPacket struct {
	Settings       uint8  // leap yr indicator, ver number, and mode
	Stratum        uint8  // stratum of local clock
	Poll           int8   // poll exponent
	Precision      int8   // precision exponent
	RootDelay      uint32 // root delay
	RootDispersion uint32 // root dispersion
	ReferenceID    uint32 // reference id
	RefTimeSec     uint32 // reference timestamp sec
	RefTimeFrac    uint32 // reference timestamp fractional
	OrigTimeSec    uint32 // origin time secs
	OrigTimeFrac   uint32 // origin time fractional
	RxTimeSec      uint32 // receive time secs
	RxTimeFrac     uint32 // receive time frac
	TxTimeSec      uint32 // transmit time secs
	TxTimeFrac     uint32 // transmit time frac
}
