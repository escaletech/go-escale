package logger

import (
	"fmt"
	"net/http"
	"time"

	"github.com/escaletech/go-escale/requestid"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func Middleware(env string) mux.MiddlewareFunc {
	logger := New(env)
	return func(next http.Handler) http.Handler {
		return &middleware{next, logger}
	}
}

type middleware struct {
	next   http.Handler
	logger *logrus.Logger
}

func (h *middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rid := requestid.Get(r)
	log := h.logger.WithField("request_id", rid)
	r = SetInRequest(r, log)

	start := time.Now()

	httpFields := logrus.Fields{
		"method": r.Method,
		"url":    r.URL.String(),
	}

	fields := logrus.Fields{
		"action": r.URL.Path,
		"http":   httpFields,
	}

	lw := newLoggerReponseWriter(w)
	h.next.ServeHTTP(lw, r)

	latency := time.Since(start).Milliseconds()
	message := fmt.Sprintf("%v %v | %vms | %v \"%v\"", r.Method, r.RequestURI, latency, lw.status, http.StatusText(lw.status))

	fields["duration"] = latency
	httpFields["status_code"] = lw.status

	if lw.status >= http.StatusBadRequest && lw.status < http.StatusInternalServerError {
		log.WithFields(fields).Warn(message)
		return
	}

	if lw.status >= http.StatusInternalServerError {
		log.WithFields(fields).Error(message)
		return
	}

	log.WithFields(fields).Info(message)
}

// loggerReponseWriter - wrapper to ResponseWriter
type loggerReponseWriter struct {
	http.Flusher
	http.ResponseWriter
	http.CloseNotifier
	status int
}

func newLoggerReponseWriter(w http.ResponseWriter) *loggerReponseWriter {
	var flusher http.Flusher
	var cNotifier http.CloseNotifier
	var ok bool
	if flusher, ok = w.(http.Flusher); !ok {
		flusher = nil
	}

	if cNotifier, ok = w.(http.CloseNotifier); !ok {
		cNotifier = nil
	}

	return &loggerReponseWriter{flusher, w, cNotifier, http.StatusOK}
}

func (lrw *loggerReponseWriter) Write(body []byte) (int, error) {
	return lrw.ResponseWriter.Write(body)
}

func (lrw *loggerReponseWriter) WriteHeader(code int) {
	lrw.status = code
	lrw.ResponseWriter.WriteHeader(code)
}
