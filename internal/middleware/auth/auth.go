package auth

import (
	"github.com/darleet/blog-graphql/pkg/utils"
	"net/http"
)

// Middleware is a simple middleware that puts the userID from cookies in the context
func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c, err := r.Cookie("user-id")

			// Allow unauthenticated users in
			if err != nil || c == nil {
				next.ServeHTTP(w, r)
				return
			}

			// put userID in context
			ctx := utils.SetUserID(r.Context(), c.Value)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
