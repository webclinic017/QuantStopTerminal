package engine

import (
	"encoding/binary"
	"fmt"
	"github.com/quantstop/quantstopterminal/internal/log"
	"github.com/quantstop/quantstopterminal/internal/ntpmonitor"
	"net"
	"sync"
	"time"
)

type NTPCheckerSubsystem struct {
	Subsystem
	wg                        sync.WaitGroup
	level                     int64
	allowedDifference         time.Duration
	allowedNegativeDifference time.Duration
	pools                     []string
	checkInterval             time.Duration
	retryLimit                int
}

func (s *NTPCheckerSubsystem) init(bot *Engine, name string) error {
	if err := s.Subsystem.init(bot, name, bot.Config.NTP.Enabled); err != nil {
		return err
	}

	if s.bot.Config.NTP.AllowedNegativeDifference == nil ||
		s.bot.Config.NTP.AllowedDifference == nil {
		return ntpmonitor.ErrNilNTPConfigValues
	}

	s.level = int64(s.bot.Config.NTP.Level)
	s.allowedDifference = *s.bot.Config.NTP.AllowedDifference
	s.allowedNegativeDifference = *s.bot.Config.NTP.AllowedNegativeDifference
	s.pools = s.bot.Config.NTP.Pool
	s.checkInterval = ntpmonitor.DefaultNTPCheckInterval
	s.retryLimit = ntpmonitor.DefaultRetryLimit
	s.initialized = true
	log.Debugln(log.NTPLogger, s.name+MsgSubsystemInitialized)
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

	log.Debugln(log.NTPLogger, s.name+MsgSubsystemStarted)
	//wg.Add(1)
	//s.wg.Add(1)
	go s.run(wg)
	return nil
}

// stop attempts to shut down the subsystem
func (s *NTPCheckerSubsystem) stop() error {
	if err := s.Subsystem.stop(); err != nil {
		return err
	}

	close(s.shutdown)

	//s.wg.Wait()
	s.started = false
	log.Debugln(log.NTPLogger, s.name+MsgSubsystemShutdown)

	return nil
}

// continuously checks the internet connection at intervals
func (s *NTPCheckerSubsystem) run(wg *sync.WaitGroup) {
	log.Debugln(log.NTPLogger, s.name+" goroutine started.")
	t := time.NewTicker(s.checkInterval)
	defer func() {
		t.Stop()
		//s.wg.Done()
		//wg.Done()
		log.Debugln(log.NTPLogger, s.name+" goroutine stopped.")
	}()

	for {
		select {
		case <-s.shutdown:
			log.Debugln(log.NTPLogger, s.name+" goroutine stopping.")
			return
		case <-t.C:
			err := s.processTime()
			if err != nil {
				log.Error(log.NTPLogger, err)
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
		log.Warnf(log.NTPLogger, "NTP manager: Time out of sync (NTP): %v | (time.Now()): %v | (Difference): %v | (Allowed): +%v / %v\n",
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
			log.Warnf(log.NTPLogger, "Unable to connect to hosts %v attempting next", s.pools[i])
			continue
		}

		if err = con.SetDeadline(time.Now().Add(5 * time.Second)); err != nil {
			log.Warnf(log.NTPLogger, "Unable to SetDeadline. Error: %s\n", err)
			err = con.Close()
			if err != nil {
				log.Error(log.NTPLogger, err)
			}
			continue
		}

		req := &ntpmonitor.NTPPacket{Settings: 0x1B}
		if err = binary.Write(con, binary.BigEndian, req); err != nil {
			log.Warnf(log.NTPLogger, "Unable to write. Error: %s\n", err)
			err = con.Close()
			if err != nil {
				log.Error(log.NTPLogger, err)
			}
			continue
		}

		rsp := &ntpmonitor.NTPPacket{}
		if err = binary.Read(con, binary.BigEndian, rsp); err != nil {
			log.Warnf(log.NTPLogger, "Unable to read. Error: %s\n", err)
			err = con.Close()
			if err != nil {
				log.Error(log.NTPLogger, err)
			}
			continue
		}

		secs := float64(rsp.TxTimeSec) - 2208988800
		nanos := (int64(rsp.TxTimeFrac) * 1e9) >> 32

		err = con.Close()
		if err != nil {
			log.Error(log.NTPLogger, err)
		}
		return time.Unix(int64(secs), nanos)
	}
	log.Warnln(log.NTPLogger, "No valid NTP servers found, using current system time")
	return time.Now().UTC()
}
