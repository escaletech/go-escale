package httpclient

import (
	"io"
	"net/http"
)

type HTTPClient struct {
	Client *http.Client
}

type Config struct {
	TimeoutInSeconds   int
	InsecureSkipVerify bool
}

type Request struct {
	Method  string
	URL     string
	Headers map[string]string
	Body    io.Reader
	Config  Config
}
