package logger

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/escaletech/go-escale/requestid"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func Middleware(env string, sensitiveFields []string) mux.MiddlewareFunc {
	logger := New(env)
	return func(next http.Handler) http.Handler {
		return &middleware{next, logger, sensitiveFields}
	}
}

func (h *middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rid := requestid.Get(r)
	log := h.logger.WithField("request_id", rid)
	r = SetInRequest(r, log)

	start := time.Now()
	cleanURL := h.suppressSensitiveQueryStrings(r.RequestURI)

	httpFields := logrus.Fields{
		"method": r.Method,
		"url":    cleanURL,
	}

	fields := logrus.Fields{
		"action": r.URL.Path,
		"http":   httpFields,
	}

	lw := newLoggerReponseWriter(w)
	h.next.ServeHTTP(lw, r)

	latency := time.Since(start).Milliseconds()
	message := fmt.Sprintf("%v %v | %vms | %v \"%v\" | Source %v", r.Method, cleanURL, latency, lw.status, http.StatusText(lw.status), r.RemoteAddr)

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

func (h *middleware) suppressSensitiveQueryStrings(urlStr string) string {
	u, _ := url.Parse(urlStr)
	values, _ := url.ParseQuery(u.RawQuery)

	for _, field := range h.sensitiveFields {
		values.Set(field, "")
	}

	u.RawQuery = values.Encode()
	return u.String()
}
