package shithttpcli

import (
	"bytes"
	"context"
	"net/http"
)

func NewHttpClient(baseURL string, opts ...ClientOption) *ClientConfig {
	client := &ClientConfig{
		httpClient: &http.Client{},
		baseURL:    baseURL,
	}
	for _, opt := range opts {
		opt(client)
	}
	return client
}

func (c *ClientConfig) Get(ctx context.Context, endpoint string, reqConfig *RequestConfig) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+endpoint, nil)
	if err != nil {
		return nil, err
	}

	for key, header := range reqConfig.headers {
		req.Header.Set(key, header)
	}

	q := req.URL.Query()
	for key, value := range reqConfig.query {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	return c.httpClient.Do(req)
}

func (c *ClientConfig) Post(ctx context.Context, reqConfig *RequestConfig) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, "POST", c.baseURL, bytes.NewBuffer(reqConfig.body))
	if err != nil {
		return nil, err
	}

	for key, header := range reqConfig.headers {
		req.Header.Set(key, header)
	}

	q := req.URL.Query()
	for key, value := range reqConfig.query {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	return c.httpClient.Do(req)
}
