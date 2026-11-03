package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"training/handler"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	mux := http.NewServeMux()
	//Health Handler API
	mux.HandleFunc("GET /p1/training/health", handler.HealthHandler)

	//Hello Handler API
	mux.HandleFunc("GET /p1/training/hello", handler.HelloHandler)

	//logger
	LoggerMux := handler.LoggingMiddleware(logger, mux)
	log.Println("logger", LoggerMux)
	log.Println("Server started on port 8080")
	err := http.ListenAndServe(":8080", LoggerMux)
	if err != nil {
		log.Fatal(err)
	}
}
