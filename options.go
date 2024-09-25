package shithttpcli

import (
	"net/http"
	"time"
)

func WithCustomTransport(ht *http.Client) ClientOption {
	return func(c *ClientConfig) {
		c.httpClient.Transport = ht.Transport
	}
}

func WithHeader(key, value string) Option {
	return func(c *ClientConfig) {
		if c.headers == nil {
			c.headers = make(map[string]Header)
		}
		c.headers[key] = Header{Value: value}
	}
}

func WithQuery(key, value string) Option {
	return func(c *ClientConfig) {
		if c.query == nil {
			c.query = make(map[string]string)
		}
		c.query[key] = value
	}
}

func WithBody(body []byte) Option {
	return func(c *ClientConfig) {
		c.body = body
	}
}

func WithRetry(r *Retry) ClientOption {
	return func(c *ClientConfig) {
		c.retry.hasRetry = true
		c.retry.RetryAttempts = r.RetryAttempts
		c.retry.RetryDelay = r.RetryDelay
	}
}

func WithTimeout(t int) ClientOption {
	return func(c *ClientConfig) {
		c.timeout = time.Duration(t) * time.Millisecond
	}
}
