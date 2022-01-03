package coinbasepro

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type APIClient struct {
	auth       *Auth
	baseURL    *url.URL
	feedURL    *url.URL
	httpClient *http.Client
	timestamp  func() string
}

func NewAPIClient(auth *Auth) (*APIClient, error) {

	baseURL, _ := url.Parse(coinbaseproAPIURL)
	feedURL, _ := url.Parse(coinbaseproWebsocketURL)

	apiClient := APIClient{
		auth:       auth,
		baseURL:    baseURL,
		feedURL:    feedURL,
		httpClient: http.DefaultClient,
		timestamp: func() string {
			return strconv.FormatInt(time.Now().Unix(), 10)
		},
	}
	return &apiClient, nil
}

func NewSandboxAPIClient(auth *Auth) (*APIClient, error) {

	baseURL, _ := url.Parse(coinbaseproSandboxRestAPIURL)
	feedURL, _ := url.Parse(coinbaseproSandboxWebsocketURL)

	apiClient := APIClient{
		auth:       auth,
		baseURL:    baseURL,
		feedURL:    feedURL,
		httpClient: http.DefaultClient,
		timestamp: func() string {
			return strconv.FormatInt(time.Now().Unix(), 10)
		},
	}
	return &apiClient, nil
}

func (a *APIClient) Get(ctx context.Context, relativePath string, result interface{}) error {
	return a.Do(ctx, "GET", relativePath, nil, result)
}

func (a *APIClient) Post(ctx context.Context, relativePath string, content interface{}, result interface{}) error {
	return a.Do(ctx, "POST", relativePath, content, result)
}

func (a *APIClient) Do(ctx context.Context, method string, relativePath string, content interface{}, result interface{}) (capture error) {
	resp, err := a.do(ctx, method, relativePath, content, result)
	if err != nil {
		return err
	}
	if isPaged(resp) {
		err = paginate(resp, result)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *APIClient) do(ctx context.Context, method string, relativePath string, content interface{}, result interface{}) (resp *http.Response, capture error) {
	//println("Base URL:" + a.baseURL.String())
	uri, err := a.baseURL.Parse(relativePath)
	//println("Request URL:" + uri.String())
	if err != nil {
		return nil, err
	}
	//fmt.Printf("%s %s\n", method, relativePath)
	var b bytes.Buffer
	if content != nil {
		err = json.NewEncoder(&b).Encode(content)
		if err != nil {
			return nil, err
		}
	}
	timestamp := a.timestamp()
	msg := fmt.Sprintf("%s%s%s%s", timestamp, method, relativePath, b.Bytes())
	signature, err := a.auth.Sign(msg)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, method, uri.String(), &b)
	if err != nil {
		return nil, err
	}
	a.addHeaders(req, timestamp, signature)
	resp, err = a.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 300 {
		coinbaseErr := Error{StatusCode: resp.StatusCode}
		decoder := json.NewDecoder(resp.Body)
		if err = decoder.Decode(&coinbaseErr); err != nil {
			return nil, err
		}
		return nil, coinbaseErr
	}
	defer func() { Capture(&capture, resp.Body.Close()) }()
	if result != nil {
		decoder := json.NewDecoder(resp.Body)
		if err = decoder.Decode(result); err != nil {
			return nil, err
		}
	}
	return resp, err
}

func (a *APIClient) Close() error { return nil }

func (a *APIClient) addHeaders(req *http.Request, timestamp string, signature string) {
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "Golang Reticule v0.1")
	req.Header.Add("CB-ACCESS-KEY", a.auth.Key)
	req.Header.Add("CB-ACCESS-PASSPHRASE", a.auth.Passphrase)
	req.Header.Add("CB-ACCESS-TIMESTAMP", timestamp)
	req.Header.Add("CB-ACCESS-SIGN", signature)
}

func isPaged(resp *http.Response) bool {
	return resp.Header.Get("CB-BEFORE") != "" && resp.Header.Get("CB-AFTER") != ""
}

func paginate(resp *http.Response, result interface{}) error {
	paginated := struct {
		Page *Pagination
	}{
		&Pagination{
			Before: resp.Header.Get("CB-BEFORE"),
			After:  resp.Header.Get("CB-AFTER"),
		},
	}
	if _, ok := result.(*json.RawMessage); ok {
		// pagination is never present in raw result, just skip
		return nil
	}
	return mapstructure.Decode(paginated, result)
}
