package shithttpcli

import (
	"net/http"
	"time"
)

type Retry struct {
	hasRetry      bool
	RetryAttempts uint
	RetryDelay    time.Duration
}

type ClientConfig struct {
	httpClient *http.Client
	headers    map[string]Header
	BaseUrl    string
	timeout    time.Duration
	retry      Retry
	query      Query
	body       Body
}

type Option func(*ClientConfig)
type ClientOption Option

type (
	Body   []byte
	Query  map[string]string
	Header struct {
		Value string
	}
)

type Response struct {
	resp *http.Response
	body Body
}
