package logger

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

type middleware struct {
	next            http.Handler
	logger          *logrus.Logger
	sensitiveFields []string
}

type loggerReponseWriter struct {
	http.Flusher
	http.ResponseWriter
	http.CloseNotifier
	status int
}
