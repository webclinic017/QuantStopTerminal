package engine

import (
	"encoding/binary"
	"fmt"
	"github.com/quantstop/quantstopterminal/internal/config"
	"github.com/quantstop/quantstopterminal/internal/ntpmonitor"
	"github.com/quantstop/quantstopterminal/internal/qstlog"
	"net"
	"sync"
	"time"
)

type NTPCheckerSubsystem struct {
	Subsystem
	level                     int64
	allowedDifference         time.Duration
	allowedNegativeDifference time.Duration
	pools                     []string
	checkInterval             time.Duration
	retryLimit                int
}

func (s *NTPCheckerSubsystem) init(config *config.Config, name string) error {
	if err := s.Subsystem.init(config, name); err != nil {
		return err
	}

	if s.config.NTP.AllowedNegativeDifference == nil ||
		s.config.NTP.AllowedDifference == nil {
		return ntpmonitor.ErrNilNTPConfigValues
	}

	s.level = int64(s.config.NTP.Level)
	s.allowedDifference = *s.config.NTP.AllowedDifference
	s.allowedNegativeDifference = *s.config.NTP.AllowedNegativeDifference
	s.pools = s.config.NTP.Pool
	s.checkInterval = ntpmonitor.DefaultNTPCheckInterval
	s.retryLimit = ntpmonitor.DefaultRetryLimit
	s.enabled = config.NTP.Enabled
	s.initialized = true
	qstlog.Debugln(qstlog.NTPLogger, s.name+MsgSubsystemInitialized)
	return nil
}

// start runs the subsystem
func (s *NTPCheckerSubsystem) start(wg *sync.WaitGroup) (err error) {
	if err = s.Subsystem.start(wg); err != nil {
		return err
	}

	s.started = true
	if s.enabled && s.initialized {
		// Sometimes the NTP client can have transient issues due to UDP, try
		// the default retry limits before giving up
	check:
		for i := 0; i < s.retryLimit; i++ {
			err := s.processTime()
			switch err {
			case nil:
				break check
			case ErrSubsystemNotStarted:
				s.started = false
				return fmt.Errorf("NTP subsystem start() check %w", ErrSubsystemNotStarted)
			default:
				if i == s.retryLimit-1 {
					return err
				}
			}
		}
	}
	if !s.enabled {
		s.started = false
		return ntpmonitor.ErrNTPSubsystemDisabled
	}
	qstlog.Debugln(qstlog.NTPLogger, s.name+MsgSubsystemStarted)
	go s.run()
	//logger.Debugf(logger.NTPLogger, "NTP subsystem %s", MsgSubSystemStarted)
	return nil
}

// stop attempts to shut down the subsystem
func (s *NTPCheckerSubsystem) stop() error {
	if err := s.Subsystem.stop(); err != nil {
		return err
	}

	s.started = false

	//logger.Debugf(logger.NTPLogger, "NTP manager %s", MsgSubSystemShuttingDown)
	qstlog.Debugln(qstlog.NTPLogger, s.name+MsgSubsystemShutdown)
	close(s.shutdown)
	return nil
}

// continuously checks the internet connection at intervals
func (s *NTPCheckerSubsystem) run() {
	t := time.NewTicker(s.checkInterval)
	defer func() {
		t.Stop()
	}()

	for {
		select {
		case <-s.shutdown:
			return
		case <-t.C:
			err := s.processTime()
			if err != nil {
				qstlog.Error(qstlog.NTPLogger, err)
			}
		}
	}
}

// FetchNTPTime returns the time from defined NTP pools
func (s *NTPCheckerSubsystem) FetchNTPTime() (time.Time, error) {
	if s == nil {
		return time.Time{}, fmt.Errorf("ntp manager %w", ErrNilSubsystem)
	}
	if s.started == false {
		return time.Time{}, fmt.Errorf("NTP manager %w", ErrSubsystemNotStarted)
	}
	return s.checkTimeInPools(), nil
}

// processTime determines the difference between system time and NTP time
// to discover discrepancies
func (s *NTPCheckerSubsystem) processTime() error {
	if s.started == false {
		return fmt.Errorf("NTP manager %w", ErrSubsystemNotStarted)
	}
	NTPTime, err := s.FetchNTPTime()
	if err != nil {
		return err
	}
	currentTime := time.Now()
	diff := NTPTime.Sub(currentTime)
	configNTPTime := s.allowedDifference
	negDiff := s.allowedNegativeDifference
	configNTPNegativeTime := -negDiff
	if diff > configNTPTime || diff < configNTPNegativeTime {
		qstlog.Warnf(qstlog.NTPLogger, "NTP manager: Time out of sync (NTP): %v | (time.Now()): %v | (Difference): %v | (Allowed): +%v / %v\n",
			NTPTime,
			currentTime,
			diff,
			configNTPTime,
			configNTPNegativeTime)
	}
	return nil
}

// checkTimeInPools returns local based on ntp servers provided timestamp
// if no server can be reached will return local time in UTC()
func (s *NTPCheckerSubsystem) checkTimeInPools() time.Time {
	for i := range s.pools {
		con, err := net.DialTimeout("udp", s.pools[i], 5*time.Second)
		if err != nil {
			qstlog.Warnf(qstlog.NTPLogger, "Unable to connect to hosts %v attempting next", s.pools[i])
			continue
		}

		if err = con.SetDeadline(time.Now().Add(5 * time.Second)); err != nil {
			qstlog.Warnf(qstlog.NTPLogger, "Unable to SetDeadline. Error: %s\n", err)
			err = con.Close()
			if err != nil {
				qstlog.Error(qstlog.NTPLogger, err)
			}
			continue
		}

		req := &ntpmonitor.NTPPacket{Settings: 0x1B}
		if err = binary.Write(con, binary.BigEndian, req); err != nil {
			qstlog.Warnf(qstlog.NTPLogger, "Unable to write. Error: %s\n", err)
			err = con.Close()
			if err != nil {
				qstlog.Error(qstlog.NTPLogger, err)
			}
			continue
		}

		rsp := &ntpmonitor.NTPPacket{}
		if err = binary.Read(con, binary.BigEndian, rsp); err != nil {
			qstlog.Warnf(qstlog.NTPLogger, "Unable to read. Error: %s\n", err)
			err = con.Close()
			if err != nil {
				qstlog.Error(qstlog.NTPLogger, err)
			}
			continue
		}

		secs := float64(rsp.TxTimeSec) - 2208988800
		nanos := (int64(rsp.TxTimeFrac) * 1e9) >> 32

		err = con.Close()
		if err != nil {
			qstlog.Error(qstlog.NTPLogger, err)
		}
		return time.Unix(int64(secs), nanos)
	}
	qstlog.Warnln(qstlog.NTPLogger, "No valid NTP servers found, using current system time")
	return time.Now().UTC()
}
