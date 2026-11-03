package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"
)

// w http.ResponseWriter := we use to send response to client
// r *http.Request := contain information about the incomng http requests. Like HTTP Method,url path,Headers,Request body

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]string{
		"status": "ok",
	}
	json.NewEncoder(w).Encode(response)
}
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]string{
		"message": "Welcome to GO Training API",
	}
	json.NewEncoder(w).Encode(response)
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func LoggingMiddleware(logger *slog.Logger, next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		logger.Info(
			"request received",
			"method", r.Method,
			"path", r.URL.Path,
		)

		rw := &responseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(rw, r)

		duration := time.Since(start)

		logger.Info(
			"request completed",
			"method", r.Method,
			"path", r.URL.Path,
			"status", rw.statusCode,
			"duration", duration,
		)
	})
}
