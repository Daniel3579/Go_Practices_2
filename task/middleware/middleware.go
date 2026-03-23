package middleware

import (
	"context"
	"net/http"
	"separation/task/auth"
)

func ValidateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var token string = r.Header.Get("AuthorizationAccess")
		username, err, status := auth.RequestValidate(token)

		if err != nil {
			if status == http.StatusUnauthorized {
				var refreshToken string = r.Header.Get("AuthorizationRefresh")
				accessToken, err, status := auth.RequestRefreshToken(refreshToken)

				if err != nil {
					http.Error(w, err.Error(), status)
					return
				}

				username, err, status = auth.RequestValidate(accessToken)

				if err != nil {
					http.Error(w, err.Error(), status)
					return
				}
			} else {
				http.Error(w, err.Error(), status)
				return
			}
		}

		// Set the username in the context for use in the next handler
		ctx := context.WithValue(r.Context(), "username", username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
