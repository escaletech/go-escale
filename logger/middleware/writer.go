package middleware

import "net/http"

type ResponseWriter struct {
	rw     http.ResponseWriter
	status int
}

func NewLoggerResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{rw: w, status: http.StatusOK}
}

func (lrw *ResponseWriter) Write(body []byte) (int, error) {
	return lrw.rw.Write(body)
}

func (lrw *ResponseWriter) Header() http.Header {
	return lrw.rw.Header()
}

func (lrw *ResponseWriter) WriteHeader(code int) {
	lrw.status = code
	lrw.rw.WriteHeader(code)
}
