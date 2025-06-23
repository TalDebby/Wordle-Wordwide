package middleware

import (
	"log"
	"net/http"
	"time"
)

type Middleware func(next http.Handler) http.Handler

type wrappedResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func createStack(middlewares ...Middleware) Middleware {
	stackedMiddleware := func(next http.Handler) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}
		return next
	}

	return stackedMiddleware
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		log.Printf("Start request: %s %s", r.Method, r.URL.Path)
		wrappedWriter := &wrappedResponseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrappedWriter, r)

		log.Printf("Finish request: %d %s %s %s", wrappedWriter.statusCode, r.Method, r.URL.Path, time.Since(startTime))
	})
}
