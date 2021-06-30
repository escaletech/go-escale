package httpclient

import (
	"crypto/tls"
	"errors"
	"net/http"
	"time"

	"github.com/escaletech/go-escale/messages"
	"github.com/escaletech/go-escale/slicer"
)

var Client *http.Client

func init() {
	Client = &http.Client{}
}

// Create a new HTTPClient
func New() *HTTPClient {
	return &HTTPClient{
		Client: Client,
	}
}

// Make a HTTP request using default http.Client.Do function
func (h *HTTPClient) DoRequest(params Request) (*http.Response, error) {
	if err := validate(params); err != nil {
		return nil, err
	}

	req, err := http.NewRequest(params.Method, params.URL, params.Body)
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

	Client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: requestConfig.InsecureSkipVerify},
	}

	h.Client = Client
}

func (h *HTTPClient) setHeaders(req *http.Request, headers map[string]string) *http.Request {
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	return req
}

func validate(params Request) error {
	var errMessage string

	allowedMethods := []string{"GET", "POST", "PATCH", "PUT", "DELETE"}
	methodAllowed, _ := slicer.Includes(params.Method, allowedMethods)

	if !*methodAllowed {
		errMessage = messages.RequestMethodNotAllowed
	}

	if params.URL == "" {
		errMessage = messages.MissingRequestURL
	}

	return errors.New(errMessage)
}
