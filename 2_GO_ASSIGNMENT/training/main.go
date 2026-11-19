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
	"training/middleware"
	"training/repository"
	domain "training/service"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file ;", err)
	}
	// Load Configuration
	cfg := config.LocalConfig()

	if cfg.JWTSecret == "" {
		log.Fatal("JWT_SECRET is not configured")
	}

	// Database Connection
	db, err := config.ConnectDB()

	if err != nil {
		log.Fatal("Database connection failed :", err)
	}

	defer db.Close()

	log.Println("Database connected successfully")

	//dept Repository
	departmentRepo := repository.NewDepartmentRepository(db)

	//dept Service
	departmentService := domain.NewDepartmentService(departmentRepo)

	//dept Handler
	departmentHandler := handler.NewDepartmentHandler(departmentService)

	//User Repository
	userRepo := repository.NewUserRepository(db)

	//Auth Service
	authService := domain.NewAuthService(userRepo)

	//Auth Handler
	authHandler := handler.NewAuthHandler(
		authService,
		cfg,
	)

	// Router
	mux := http.NewServeMux()

	// department Api
	mux.HandleFunc(
		"GET /v1/training/departments",
		departmentHandler.ListDepartments,
	)
	mux.HandleFunc(
		"GET /v1/training/departments/{id}",
		departmentHandler.GetDepartment,
	)
	// Authentication API
	mux.HandleFunc(
		"POST /v1/training/login",
		authHandler.Login,
	)

	//Register
	mux.HandleFunc(
		"POST /v1/training/register",
		authHandler.RegisterRequest,
	)

	//current user
	mux.Handle(
		"GET /v1/training/me",
		middleware.AuthMiddleware(cfg.JWTSecret)(
			http.HandlerFunc(authHandler.Me),
		),
	)

	//logout
	mux.Handle(
		"POST /v1/training/logout",
		middleware.AuthMiddleware(cfg.JWTSecret)(
			http.HandlerFunc(authHandler.Logout),
		),
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
