package httpclient

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

type HttpClientMock struct {
	mock.Mock
}

func (mock *HttpClientMock) DoRequest(params Request) (*http.Response, error) {
	args := mock.Called()
	result := args.Get(0)

	return result.(*http.Response), args.Error(1)
}
