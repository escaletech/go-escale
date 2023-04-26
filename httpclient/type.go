package httpclient

import (
	"crypto/tls"
	"io"
	"net/http"
)

type HTTPClientInterface interface {
	DoRequest(params Request) (*http.Response, error)
}

type HTTPClient struct {
	Client    *http.Client
	tlsConfig *tls.Config // novo campo para armazenar a configuração do certificado
}

type Config struct {
	TimeoutInSeconds   int
	InsecureSkipVerify bool
	Certificates       []tls.Certificate
}

type Request struct {
	Method      string
	URL         string
	Headers     map[string]string
	Body        io.Reader
	Config      Config
	QueryParams map[string]string
}
