package middleware

import (
	"context"
	"net/http"

	"github.com/mhdianrush/go-json-web-token/helper"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessToken := r.Header.Get("Authorization")
		if accessToken == "" {
			// Unauthorize
			helper.Response(w, 401, "Unauthorized", nil)
			return
		}

		user, err := helper.ValidateToken(accessToken)
		if err != nil {
			helper.Response(w, 500, err.Error(), nil)
			return
		}

		ctx := context.WithValue(r.Context(), "userinfo", user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
