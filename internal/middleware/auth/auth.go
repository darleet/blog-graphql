package auth

import (
	"github.com/darleet/blog-graphql/pkg/utils"
	"net/http"
)

// Middleware is a simple middleware that puts the userID from cookies in the context
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := r.Header.Get("User-ID")

		// Allow unauthenticated users in
		if c == "" {
			next.ServeHTTP(w, r)
			return
		}

		// put userID in context
		ctx := utils.SetUserID(r.Context(), c)

		// and call the next with our new context
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
