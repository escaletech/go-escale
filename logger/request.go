package logger

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

type contextKey struct{}

var noop *logrus.Entry

func init() {
	logger := logrus.New()
	logger.Out = ioutil.Discard
	noop = logger.WithField("noop", true)
}

func SetInRequest(r *http.Request, l *logrus.Entry) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), contextKey{}, l))
}

func Get(r *http.Request) *logrus.Entry {
	log, ok := r.Context().Value(contextKey{}).(*logrus.Entry)
	if !ok {
		return noop
	}

	return log
}
