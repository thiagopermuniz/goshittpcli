package shithttpcli

import (
	"net/http"
	"time"
)

func NewRequestConfig() *RequestConfig {
	return &RequestConfig{
		Headers: make(map[string]string),
		Query:   make(map[string]string),
	}
}

func WithCustomTransport(ht *http.Client) ClientOption {
	return func(c *ClientConfig) {
		c.httpClient.Transport = ht.Transport
	}
}

// WithHeader adiciona um cabeçalho à configuração da requisição
func (r *RequestConfig) WithHeader(key, value string) *RequestConfig {
	r.Headers[key] = value
	return r
}

// WithQuery adiciona um parâmetro de query à configuração da requisição
func (r *RequestConfig) WithQuery(key, value string) *RequestConfig {
	r.Query[key] = value
	return r
}

// WithBody define o corpo da requisição
func (r *RequestConfig) WithBody(body []byte) *RequestConfig {
	r.Body = body
	return r
}

func WithRetry(r *RetryConfig) ClientOption {
	return func(c *ClientConfig) {
		c.retry.enabled = r.enabled
		c.retry.attempts = r.attempts
		c.retry.delay = r.delay
	}
}

func WithTimeout(t time.Duration) ClientOption {
	return func(c *ClientConfig) {
		c.timeout = t
	}
}
