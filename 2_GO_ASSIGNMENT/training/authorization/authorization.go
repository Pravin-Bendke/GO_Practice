package authorization

import (
	"encoding/json"
	"net/http"
	"strings"
	reqctx "training/request"
)

func RequireRole(
	allowedRoles ...string,
) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userRole := reqctx.GetRole(r.Context())

			if userRole == "" {
				writeError(
					w,
					http.StatusUnauthorized,
					"authentication required",
				)
				return
			}
			for _, role := range allowedRoles {
				if strings.EqualFold(
					strings.TrimSpace(userRole),
					strings.TrimSpace(role),
				) {
					next.ServeHTTP(w, r)
					return
				}
			}
			writeError(
				w,
				http.StatusForbidden,
				"forbidden: insufficient permission",
			)

		})
	}
}

func writeError(
	w http.ResponseWriter,
	status int,
	message string,
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(
		map[string]string{
			"error": message,
		},
	)
}
