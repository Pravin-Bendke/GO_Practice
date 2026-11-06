// package main

// import (
// 	"log"
// 	"log/slog"
// 	"net/http"
// 	"os"

// 	"training/config"
// 	"training/handler"
// 	"training/repository"
// )

// func main() {

// 	// Logger
// 	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

// 	// Connect PostgreSQL Database
// 	db, err := config.ConnectDB()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	log.Println("PostgreSQL connected successfully")

// 	// Create Department Repository
// 	departmentRepo := repository.NewDepartmentRepository(db)

// 	// Create Department Handler
// 	departmentHandler := handler.NewDepartmentHandler(departmentRepo)

// 	// Create Router
// 	mux := http.NewServeMux()

// 	// Health Handler API
// 	mux.HandleFunc("GET /p1/training/health", handler.HealthHandler)

// 	// Hello Handler API
// 	mux.HandleFunc("GET /p1/training/hello", handler.HelloHandler)

// 	// Department APIs
// 	mux.HandleFunc("GET /p1/training/department/{id}", departmentHandler.GetByID)
// 	mux.HandleFunc("GET /p1/training/department", departmentHandler.ListByCompanyID)

// 	// Logging Middleware
// 	loggerMux := handler.LoggingMiddleware(logger, mux)

// 	log.Println("Server started on port 8080")

// 	err = http.ListenAndServe(":8080", loggerMux)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

package main

import (
	"log"
	"net/http"
	"os"

	"training/config"
	"training/handler"
	"training/repository"
	domain "training/service"
)

func main() {

	// Database Connection
	db, err := config.ConnectDB()

	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	defer db.Close()

	// Repository
	departmentRepo := repository.NewDepartmentRepository(db)

	// Service
	departmentService := domain.NewDepartmentService(
		departmentRepo,
	)

	// Handler
	departmentHandler := handler.NewDepartmentHandler(
		departmentService,
	)

	// Router
	mux := http.NewServeMux()

	mux.HandleFunc(
		"GET /v1/training/departments",
		departmentHandler.ListDepartments,
	)

	mux.HandleFunc(
		"GET /v1/training/departments/{id}",
		departmentHandler.GetDepartment,
	)

	// Server
	port := os.Getenv("APP_PORT")

	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port:", port)

	err = http.ListenAndServe(
		":"+port,
		mux,
	)

	if err != nil {
		log.Fatal("Server failed:", err)
	}
}
