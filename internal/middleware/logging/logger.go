package logging

import (
	"go.uber.org/zap"
	"net/http"
)

func NewMiddleware(log *zap.SugaredLogger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// log request
			log.Infow("Got request", "method", r.Method, "path", r.URL.Path,
				"user-agent", r.UserAgent())
			// call next
			next.ServeHTTP(w, r)
		})
	}
}
