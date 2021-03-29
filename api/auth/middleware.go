package auth

import (
	"context"
	"net/http"

	"github.com/mhd53/quanta-fitness-server/pkg/crypto"
)

var userCtxKey = &contextKey{"uname"}

type contextKey struct {
	name string
}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			// Allow unauthorized users in.
			if header == "" {
				next.ServeHTTP(w, r)
				return

			}

			tokenStr := header
			uname, err := crypto.ParseToken(tokenStr)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return

			}

			ctx := context.WithValue(r.Context(), userCtxKey, uname)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)

		})

	}

}

// Find the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) string {
	raw, _ := ctx.Value(userCtxKey).(string)
	return raw

}
