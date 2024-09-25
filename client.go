package shithttpcli

import (
	"bytes"
	"context"
	"net/http"
	"net/url"
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

// prepareRequest prepara uma requisição HTTP com base nas configurações fornecidas.
func (c *ClientConfig) prepareRequest(ctx context.Context, method, endpoint string, reqConfig *RequestConfig) (*http.Request, error) {
	// Cria a URL completa
	fullURL, err := url.Parse(c.baseURL + endpoint)
	if err != nil {
		return nil, err
	}

	// Cria a requisição com o método e o corpo (se houver)
	var body *bytes.Buffer
	if reqConfig.Body != nil {
		body = bytes.NewBuffer(reqConfig.Body)
	} else {
		body = nil
	}
	req, err := http.NewRequestWithContext(ctx, method, fullURL.String(), body)
	if err != nil {
		return nil, err
	}

	// Adiciona os cabeçalhos
	for key, header := range reqConfig.Headers {
		req.Header.Set(key, header)
	}

	// Adiciona os parâmetros de query string
	q := req.URL.Query()
	for key, value := range reqConfig.Query {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	return req, nil
}

// Get faz uma requisição HTTP GET para o endpoint especificado.
func (c *ClientConfig) Get(ctx context.Context, endpoint string, reqConfig *RequestConfig) (*http.Response, error) {
	req, err := c.prepareRequest(ctx, http.MethodGet, endpoint, reqConfig)
	if err != nil {
		return nil, err
	}
	return c.httpClient.Do(req)
}

// Post faz uma requisição HTTP POST para o endpoint especificado.
func (c *ClientConfig) Post(ctx context.Context, endpoint string, reqConfig *RequestConfig) (*http.Response, error) {
	req, err := c.prepareRequest(ctx, http.MethodPost, endpoint, reqConfig)
	if err != nil {
		return nil, err
	}
	return c.httpClient.Do(req)
}
