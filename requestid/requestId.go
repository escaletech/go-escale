package requestid

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type requestIDKey struct{}

const requestIDHeader = "X-Request-Id"

func Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, set(r))
	})
}

func set(r *http.Request) *http.Request {
	current := r.Header.Get(requestIDHeader)
	if current == "" {
		current = GenerateTrackingId()
	}

	return r.WithContext(context.WithValue(r.Context(), requestIDKey{}, current))
}

func Get(r *http.Request) string {
	return r.Context().Value(requestIDKey{}).(string)
}

func GenerateTrackingId() string {
	id, _ := uuid.NewRandom()
	return id.String()
}
