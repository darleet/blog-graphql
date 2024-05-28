package recoverer

import (
	"go.uber.org/zap"
	"net/http"
)

func Middleware(log *zap.SugaredLogger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					log.Error("Panic recovered", zap.Any("err", err))
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}
