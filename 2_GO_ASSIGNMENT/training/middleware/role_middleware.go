package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

func RequireRole(allowedRoles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			//Get authenticated user
			claims, ok := GetClaims(r.Context())

			if !ok || claims == nil {
				writeError(
					w,
					http.StatusUnauthorized,
					"authentication required",
				)
				return
			}
			userRole := strings.TrimSpace(claims.Role)

			for _, role := range allowedRoles {
				fmt.Println(role)
				if strings.EqualFold(
					userRole,
					strings.TrimSpace(role),
				) {
					next.ServeHTTP(w, r)
					return
				}
			}
			//Access denied
			writeError(
				w,
				http.StatusForbidden,
				"forbidden: not permission",
			)
		})
	}
}
