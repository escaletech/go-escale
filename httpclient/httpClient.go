package httpclient

import (
	"crypto/tls"
	"net/http"
	"time"
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

// Make an HTTP request using default http.Client.Do function
func (h *HTTPClient) DoRequest(params Request) (*http.Response, error) {
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
