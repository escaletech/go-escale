package httpserver_test

import (
	"testing"

	"github.com/escaletech/go-escale/httpserver"
	"github.com/stretchr/testify/assert"
)

type cases struct {
	env              string
	port             string
	expectedResponse string
}

func TestServerAddr(t *testing.T) {
	testCases := []cases{
		{"dev", "3000", "localhost:3000"},
		{"staging", "666", ":666"},
		{"production", "3000", ":3000"},
	}

	for _, testCase := range testCases {
		t.Run("env is "+testCase.env+" and port is "+testCase.port, func(t *testing.T) {
			t.Run("returns \""+testCase.expectedResponse+"\"", func(t *testing.T) {
				addr := httpserver.SetAddr(testCase.env, testCase.port)
				assert.Equal(t, testCase.expectedResponse, addr)
			})
		})
	}
}
