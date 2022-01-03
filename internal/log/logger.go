package log

import (
	"errors"
	"fmt"
	"io"
	"time"
)

var (
	errEmptyLoggerName            = errors.New("cannot have empty logger name")
	errSubLoggerAlreadyregistered = errors.New("sub logger already registered")
)

func newLogger(c *Config) Logger {
	return Logger{
		Timestamp:         c.AdvancedSettings.TimeStampFormat,
		Spacer:            c.AdvancedSettings.Spacer,
		ErrorHeader:       c.AdvancedSettings.Headers.Error,
		InfoHeader:        c.AdvancedSettings.Headers.Info,
		WarnHeader:        c.AdvancedSettings.Headers.Warn,
		DebugHeader:       c.AdvancedSettings.Headers.Debug,
		ShowLogSystemName: *c.AdvancedSettings.ShowLogSystemName,
	}
}

func (l *Logger) newLogEvent(data, header, slName string, w io.Writer) error {
	if w == nil {
		return errors.New("io.Writer not set")
	}

	pool, ok := eventPool.Get().(*[]byte)
	if !ok {
		return errors.New("unable to type assert slice of bytes pointer")
	}

	// append header
	*pool = append(*pool, header...)

	// append spacer
	*pool = append(*pool, l.Spacer...)

	// append timestamp
	if l.Timestamp != "" {
		*pool = time.Now().AppendFormat(*pool, l.Timestamp)
	}

	// append log system name (sublogger name)
	if l.ShowLogSystemName {
		*pool = append(*pool, l.Spacer...)
		*pool = append(*pool, slName...)
	}

	// append spacer
	*pool = append(*pool, l.Spacer...)

	// append log data
	*pool = append(*pool, data...)
	if data == "" || data[len(data)-1] != '\n' {
		*pool = append(*pool, '\n')
	}

	// write to the pool
	_, err := w.Write(*pool)
	*pool = (*pool)[:0]
	eventPool.Put(pool)

	return err
}

// CloseLogger is called on shutdown of application
func CloseLogger() error {
	return GlobalLogFile.Close()
}

// Level retries the current sublogger levels
func Level(name string) (Levels, error) {
	RWM.RLock()
	defer RWM.RUnlock()
	subLogger, found := SubLoggers[name]
	if !found {
		return Levels{}, fmt.Errorf("logger %s not found", name)
	}
	return subLogger.levels, nil
}

// SetLevel sets sublogger levels
func SetLevel(s, level string) (Levels, error) {
	RWM.Lock()
	defer RWM.Unlock()
	subLogger, found := SubLoggers[s]
	if !found {
		return Levels{}, fmt.Errorf("sub logger %v not found", s)
	}
	subLogger.SetLevels(splitLevel(level))
	return subLogger.levels, nil
}
