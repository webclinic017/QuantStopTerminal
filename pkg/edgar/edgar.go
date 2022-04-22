package edgar

import (
	"context"
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

/*
For official documentation see:  https://www.sec.gov/edgar/sec-api-documentation
*/

const DefaultClientTimeout = 10 * time.Second

var DefaultRateLimiter = rate.NewLimiter(rate.Every(10*time.Second), 50) // 50 request every 10 seconds

type Edgar struct {
	client      *http.Client
	Ratelimiter *rate.Limiter
}

func New() *Edgar {
	return &Edgar{
		client:      &http.Client{Timeout: DefaultClientTimeout},
		Ratelimiter: DefaultRateLimiter,
	}
}

//Do dispatch the HTTP request to the network
func (edgar *Edgar) Do(req *http.Request) (*http.Response, error) {

	/*
		Per the documentation at https://www.sec.gov/os/accessing-edgar-data
		Current max request rate: 10 requests/second.
		Sample Declared Bot Request Headers:
			User-Agent:  Sample Company Name AdminContact@<sample company domain>.com
			Accept-Encoding:  gzip, deflate (note: throws error with this header)
			Host:  www.sec.gov
	*/
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	// Rate Limiter
	ctx := context.Background()
	err := edgar.Ratelimiter.Wait(ctx) // This is a blocking call. Honors the rate limit
	if err != nil {
		return nil, err
	}

	// Execute the request
	resp, err := edgar.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
