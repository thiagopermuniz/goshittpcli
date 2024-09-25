package shithttpcli

import (
	"net/http"
	"time"
)

func NewRequestConfig() *RequestConfig {
	return &RequestConfig{
		headers: make(map[string]string),
		query:   make(map[string]string),
	}
}

func WithCustomTransport(ht *http.Client) ClientOption {
	return func(c *ClientConfig) {
		c.httpClient.Transport = ht.Transport
	}
}

// WithHeader adiciona um cabeçalho à configuração da requisição
func (r *RequestConfig) WithHeader(key, value string) *RequestConfig {
	r.headers[key] = value
	return r
}

// WithQuery adiciona um parâmetro de query à configuração da requisição
func (r *RequestConfig) WithQuery(key, value string) *RequestConfig {
	r.query[key] = value
	return r
}

// WithBody define o corpo da requisição
func (r *RequestConfig) WithBody(body []byte) *RequestConfig {
	r.body = body
	return r
}

func WithRetry(r *RetryConfig) ClientOption {
	return func(c *ClientConfig) {
		c.retry.enabled = r.enabled
		c.retry.attempts = r.attempts
		c.retry.delay = r.delay
	}
}

func WithTimeout(t int) ClientOption {
	return func(c *ClientConfig) {
		c.timeout = time.Duration(t) * time.Millisecond
	}
}
