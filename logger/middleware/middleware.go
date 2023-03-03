package middleware

import (
	"net/http"

	"github.com/escaletech/go-escale/logger"
	"github.com/gorilla/mux"
)

type middleware struct {
	next   http.Handler
	logger *logger.Logger
}

func Middleware(logger *logger.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return &middleware{next, logger}
	}
}

func (m *middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	reqLog := NewRequestLogger(w, m.logger)
	m.next.ServeHTTP(w, r)
	reqLog.WriteRequestLog(r)
}
