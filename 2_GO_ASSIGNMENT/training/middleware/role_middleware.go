package middleware

import (
	"fmt"
	"net/http"
)

func RequireRole(allowedRoles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			claims, ok := GetClaims(r.Context())

			if !ok || claims == nil {
				writeError(
					w,
					http.StatusUnauthorized,
					"authentication required",
				)
				return
			}
			allowed := false
			for _, role := range allowedRoles {
				fmt.Println(role)
				if claims.Role == role {
					allowed = true
					break
				}
			}
			if !allowed {
				writeError(
					w,
					http.StatusForbidden,
					"permission not allowed",
				)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
