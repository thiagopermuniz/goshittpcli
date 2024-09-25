package shithttpcli

import (
	"context"
	"io"
	"net/http"

	"github.com/avast/retry-go"
)

func NewHttpClient(c *ClientConfig, opts ...ClientOption) *ClientConfig {
	client := &ClientConfig{
		httpClient: &http.Client{},
		BaseUrl:    c.BaseUrl,
	}

	for _, opt := range opts {
		opt(client)
	}
	return client
}

func (c *ClientConfig) initOps(opts ...Option) func() {
	for _, opt := range opts {
		opt(c)
	}
	return func() {
		c.query = make(Query)
		c.body = nil
	}
}

func (c *ClientConfig) prepareReq(req *http.Request) *http.Request {
	for k, v := range c.headers {
		req.Header.Set(k, v.Value)
	}
	q := req.URL.Query()
	for k, v := range c.query {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
	return req
}

func (c *ClientConfig) prepareRes(resp *http.Response, err error) (*Response, error) {
	r := &Response{resp: resp}
	if err != nil {
		return r, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	r.body = body
	if err != nil {
		return r, err
	}

	return r, nil
}

func (c *ClientConfig) sendReq(ctx context.Context, req *http.Request) (*Response, error) {
	reqCtx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	res, err := c.prepareRes(c.httpClient.Do(req.WithContext(reqCtx)))
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *ClientConfig) sendReqWithRetry(ctx context.Context, req *http.Request) (*Response, error) {
	reqCtx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	r := &Response{}
	err := retry.Do(
		func() error {
			res, err := c.prepareRes(c.httpClient.Do(req))
			if err != nil {
				return err
			}
			r = res
			return nil
		},
		retry.Context(reqCtx),
		retry.RetryIf(func(err error) bool {
			return err != nil
		}),
		retry.Attempts(c.retry.RetryAttempts),
		retry.Delay(c.retry.RetryDelay),
		retry.LastErrorOnly(true),
	)
	if err != nil {
		return r, err
	}
	return r, nil
}

func (c *ClientConfig) Get(ctx context.Context, endpoint string, opts ...Option) (*Response, error) {
	cls := c.initOps(opts...)
	defer cls()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.BaseUrl+endpoint, nil)
	if err != nil {
		return nil, err
	}
	prepReq := c.prepareReq(req)
	if c.retry.hasRetry {
		return c.sendReqWithRetry(ctx, prepReq)
	}
	return c.sendReq(ctx, prepReq)
}
