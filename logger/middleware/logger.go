package accesslog

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/escaletech/go-escale/logger"
	"github.com/escaletech/go-escale/requestid"
)

type RequestLogger struct {
	logger         *logger.Logger
	start          time.Time
	ResponseWriter *ResponseWriter
}

func NewRequestLogger(w http.ResponseWriter, log *logger.Logger) *RequestLogger {
	tracking := requestid.GenerateTrackingId()
	log.SetTrackingId(tracking)
	return &RequestLogger{
		logger:         log,
		start:          time.Now(),
		ResponseWriter: NewLoggerResponseWriter(w),
	}
}

func (rl *RequestLogger) WriteRequestLog(r *http.Request) {
	ip, _ := rl.getIP(r)

	msg := fmt.Sprintf(
		"[ip:%v] [iss:%v] %v %v - %v duration: %vms",
		ip,
		r.Host,
		r.Method,
		r.URL.Path,
		rl.ResponseWriter.status,
		time.Since(rl.start).Milliseconds(),
	)

	rl.logger.Access(msg)

	requestInfo := rl.getRequestInfo(r)
	msgLogInfoRequest := fmt.Sprintf(
		"[header:%v] [request:%v]",
		requestInfo["header"],
		requestInfo["payload"],
	)

	if rl.ResponseWriter.status < http.StatusBadRequest {
		rl.logger.Debug(msgLogInfoRequest)
		return
	}

	if rl.ResponseWriter.status >= http.StatusInternalServerError {
		rl.logger.Error(msgLogInfoRequest)
		return
	}

	if rl.ResponseWriter.status < http.StatusInternalServerError {
		rl.logger.Warn(msgLogInfoRequest)
		return
	}
}

func (rl *RequestLogger) getIP(r *http.Request) (string, error) {
	ips := r.Header.Get("X-Forwarded-For")
	splitIps := strings.Split(ips, ",")

	if len(splitIps) > 0 {
		netIP := net.ParseIP(splitIps[len(splitIps)-1])
		if netIP != nil {
			return netIP.String(), nil
		}
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}

	netIP := net.ParseIP(ip)
	if netIP != nil {
		ip := netIP.String()
		if ip == "::1" {
			return "127.0.0.1", nil
		}
		return ip, nil
	}

	return "", errors.New("IP not found")
}

func (rl *RequestLogger) getRequestInfo(r *http.Request) logger.Fields {
	headers := r.Header

	if headers.Get("Authorization") != "" {
		headers.Set("Authorization", "*****")
	}
	return logger.Fields{
		"header":  r.Header,
		"payload": r.Body,
	}
}
