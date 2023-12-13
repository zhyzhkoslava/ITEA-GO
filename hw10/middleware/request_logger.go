package middleware

import (
	"log"
	"net/http"
)

type RequestLoggerMiddleware struct {
	logger *log.Logger
}

func NewRequestLogger(logger *log.Logger) *RequestLoggerMiddleware {
	return &RequestLoggerMiddleware{logger: logger}
}

func (m *RequestLoggerMiddleware) Wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.logger.Printf("incoming request to %s", r.URL.Path)

		next.ServeHTTP(w, r)
	})
}
