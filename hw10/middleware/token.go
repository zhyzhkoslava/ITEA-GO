package middleware

import (
	"net/http"
)

type TokenMiddleware struct {
	Token string
}

func NewTokenMiddleware(token string) *TokenMiddleware {
	return &TokenMiddleware{
		Token: token,
	}
}

func (tm *TokenMiddleware) Wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost || r.Method == http.MethodPatch || r.Method == http.MethodDelete {
			requestToken := r.Header.Get("Authorization")

			if tm.isValidToken(requestToken) {
				next.ServeHTTP(w, r)
				return
			} else {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

func (tm *TokenMiddleware) isValidToken(requestToken string) bool {
	return requestToken == tm.Token
}
