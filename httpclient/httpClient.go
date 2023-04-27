package httpclient

import (
	"crypto/tls"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/escaletech/go-escale/messages"
	"github.com/escaletech/go-escale/slicer"
)

var Client *http.Client

func init() {
	Client = &http.Client{}
}

func New() *HTTPClient {
	return &HTTPClient{
		Client: Client,
	}
}

// Define a configuração do certificado TLS
func (h *HTTPClient) SetTLSConfig(config *tls.Config) {
	h.tlsConfig = config
}

// Faça uma solicitação HTTP usando a função padrão http.Client.Do
func (h *HTTPClient) DoRequest(params Request) (*http.Response, error) {
	if err := validate(params); err != nil {
		return nil, err
	}

	url := params.URL
	if params.QueryParams != nil {
		url += castQueryParamsToString(params.QueryParams)
	}

	req, err := http.NewRequest(params.Method, url, params.Body)
	if err != nil {
		return nil, err
	}
	req = h.setHeaders(req, params.Headers)

	h.configClient(params.Config)
	return h.Client.Do(req)
}

func (h *HTTPClient) configClient(requestConfig Config) {
	if requestConfig.TimeoutInSeconds > 0 {
		Client.Timeout = time.Duration(requestConfig.TimeoutInSeconds) * time.Second
	}

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: requestConfig.InsecureSkipVerify},
	}

	// adiciona a configuração do certificado TLS ao transporte, se definida
	if h.tlsConfig != nil {
		transport.TLSClientConfig = h.tlsConfig
	}

	h.Client.Transport = transport
	h.Client = Client
}

func (h *HTTPClient) setHeaders(req *http.Request, headers map[string]string) *http.Request {
	for key, value := range headers {
		req.Header[key] = []string{value}
	}

	return req
}

func validate(params Request) error {
	var errMessage string

	allowedMethods := []string{"GET", "POST", "PATCH", "PUT", "DELETE"}
	methodAllowed, _ := slicer.Includes(strings.ToUpper(params.Method), allowedMethods)

	if !*methodAllowed {
		errMessage = messages.RequestMethodNotAllowed
	}

	if params.URL == "" {
		errMessage = messages.MissingRequestURL
	}

	if errMessage != "" {
		return errors.New(errMessage)
	}

	return nil
}

func castQueryParamsToString(pathParams map[string]string) string {
	if pathParams == nil {
		return ""
	}
	b := strings.Builder{}
	b.WriteString("?")
	for key, value := range pathParams {
		b.WriteString(key)
		b.WriteString("=")
		b.WriteString(value)
		b.WriteString("&")
	}
	l := b.Len()
	return b.String()[:l-1]
}
