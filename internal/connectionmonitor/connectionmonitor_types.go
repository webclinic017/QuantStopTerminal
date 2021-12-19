package connectionmonitor

import "time"

// ConnectionMonitorConfig defines the connection monitor variables to ensure
// that there is internet connectivity
type Config struct {
	Enabled          bool
	Initialized      bool
	DNSList          []string      `json:"preferredDNSList"`
	PublicDomainList []string      `json:"preferredDomainList"`
	CheckInterval    time.Duration `json:"checkInterval"`
}

// DefaultCheckInterval is a const that defines the amount of time between
// checking if the connection is lost
const (
	DefaultCheckInterval = time.Second

	ConnRe       = "Internet connectivity re-established"
	ConnLost     = "Internet connectivity lost"
	ConnFound    = "Internet connectivity found"
	ConnNotFound = "No internet connectivity"
)

// Default check lists
var (
	DefaultDNSList    = []string{"8.8.8.8", "8.8.4.4", "1.1.1.1", "1.0.0.1"}
	DefaultDomainList = []string{"www.google.com", "www.cloudflare.com", "www.facebook.com"}
)
