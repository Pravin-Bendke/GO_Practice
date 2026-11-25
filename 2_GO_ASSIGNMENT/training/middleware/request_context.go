package middleware

import (
	"crypto/rand"
	"fmt"
	"net/http"

	reqctx "training/request"
)

// Generate unique request ID.
func generateRequestID() (string, error) {

	bytes := make([]byte, 16)

	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	// UUID version 4.
	bytes[6] = (bytes[6] & 0x0f) | 0x40
	bytes[8] = (bytes[8] & 0x3f) | 0x80

	return fmt.Sprintf(
		"%08x-%04x-%04x-%04x-%012x",
		bytes[0:4],
		bytes[4:6],
		bytes[6:8],
		bytes[8:10],
		bytes[10:16],
	), nil
}

// Request Context Middleware.
func RequestContextMiddleware(
	next http.Handler,
) http.Handler {

	return http.HandlerFunc(func(
		w http.ResponseWriter,
		r *http.Request,
	) {

		// Generate unique request ID.
		requestID, err := generateRequestID()

		if err != nil {
			http.Error(
				w,
				"could not generate request ID",
				http.StatusInternalServerError,
			)
			return
		}

		// Store Request ID in context.
		ctx := reqctx.WithRequestID(
			r.Context(),
			requestID,
		)

		// Add Request ID to response header.
		w.Header().Set(
			"X-Request-ID",
			requestID,
		)

		// Pass updated context to next middleware.
		next.ServeHTTP(
			w,
			r.WithContext(ctx),
		)
	})
}
