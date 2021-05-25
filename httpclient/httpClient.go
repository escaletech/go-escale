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

// Make an HTTP request using default http.Client.Do function
func DoRequest(params Request) (*http.Response, error) {
	client := configClient(params.Config)

	req, err := http.NewRequest(params.Method, params.URL, params.Body)
	if err != nil {
		return nil, err
	}
	req = setHeaders(req, params.Headers)

	return client.Do(req)
}

func configClient(requestConfig Config) *http.Client {
	if requestConfig.TimeoutInSeconds > 0 {
		Client.Timeout = time.Duration(requestConfig.TimeoutInSeconds) * time.Second
	}

	Client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: requestConfig.InsecureSkipVerify},
	}

	return Client
}

func setHeaders(req *http.Request, headers map[string]string) *http.Request {
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	return req
}
