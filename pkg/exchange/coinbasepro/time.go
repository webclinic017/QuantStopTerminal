package coinbasepro

import (
	"context"
	"fmt"
	"strings"
	"time"
)

// Timestamps
//
// The docs read:
// Unless otherwise specified, all timestamps from API are returned as ISO 8601 with microseconds.
// Make sure you can parse the following ISO 8601 format.
// Most modern languages and libraries will handle this without issues.
//
//     `2014-11-06T10:34:47.123456Z`
//
// As far as I can tell, this is misleading.
// There are several flavors of time/timestamp and little documentation of when which one appears.

type Time time.Time

func (t *Time) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		var ptr Time
		*t = ptr
		return nil
	}
	layouts := []string{
		time.RFC3339Nano,
		time.RFC3339,
		"2006-01-02 15:04:05+00",
		"2006-01-02 15:04:05.999999",
		"2006-01-02 15:04:05.999999+00",
	}
	var parsedTime time.Time
	var err error
	for _, layout := range layouts {
		parsedTime, err = time.Parse(layout, strings.ReplaceAll(string(data), "\"", ""))
		if err != nil {
			continue
		}
		break
	}
	if err != nil {
		return fmt.Errorf("time %s in unhandled format", data)
	}
	*t = Time(parsedTime)
	return nil
}

// MarshalJSON marshal time back to time.Time for json encoding
func (t Time) MarshalJSON() ([]byte, error) {
	return t.Time().MarshalJSON()
}

func (t *Time) Time() time.Time {
	return time.Time(*t)
}

type ServerTime struct {
	ISO   time.Time `json:"iso"`
	Epoch float64   `json:"epoch"`
}

// GetServerTime retrieves the Coinbase Pro API server time.
func (c *Client) GetServerTime(ctx context.Context) (ServerTime, error) {
	var serverTime ServerTime
	return serverTime, c.Get(ctx, "/time", &serverTime)
}
