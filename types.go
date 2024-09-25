package shithttpcli

import (
	"net/http"
	"time"
)

// RetryConfig encapsula as configurações de retry para o cliente HTTP.
type RetryConfig struct {
	enabled  bool          // Indica se o retry está habilitado.
	attempts uint          // Número de tentativas de retry.
	delay    time.Duration // Intervalo entre as tentativas de retry.
}

// ClientConfig contém as configurações principais do cliente HTTP.
type ClientConfig struct {
	httpClient *http.Client  // Cliente HTTP subjacente.
	baseURL    string        // URL base para as requisições.
	timeout    time.Duration // Timeout para as requisições.
	retry      RetryConfig   // Configurações de retry.
}

// RequestConfig encapsula as configurações específicas de uma requisição HTTP.
type RequestConfig struct {
	headers map[string]string // Cabeçalhos HTTP.
	query   map[string]string // Parâmetros de query string.
	body    []byte            // Corpo da requisição.
}

// Option define uma função que modifica a configuração do cliente HTTP.
type Option func(*ClientConfig)

// ClientOption é um alias para Option, usado para configurar o cliente HTTP.
type ClientOption Option

// Response encapsula a resposta HTTP recebida.
type Response struct {
	response *http.Response // Resposta HTTP original.
	body     []byte         // Corpo da resposta.
}
