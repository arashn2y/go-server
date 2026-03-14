package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/arashn0uri/go-server/internal/constants"
	"github.com/arashn0uri/go-server/internal/json"
	"github.com/arashn0uri/go-server/internal/utils"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			json.WriteError(w, http.StatusUnauthorized, "missing or invalid authorization header")
			return
		}
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := utils.ValidateToken(tokenStr)
		if err != nil {
			json.WriteError(w, http.StatusUnauthorized, "invalid token: "+err.Error())
			return
		}

		ctx := context.WithValue(r.Context(), constants.ContextKeyUserID, claims.UserID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
