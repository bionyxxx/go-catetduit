package main

import (
	"catetduit/internal/config"
	"catetduit/internal/database"
	"catetduit/internal/helper"
	middleware2 "catetduit/internal/middleware"
	"catetduit/internal/module/auth"
	"catetduit/internal/module/transaction"
	"catetduit/internal/module/user"
	customValidator "catetduit/internal/validator"
	"fmt"
	_ "log"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// Load configuration (if any)
	mainConfig := config.NewConfig()
	dbConfig := config.NewDatabaseConfig()

	// Init DB connection here (omitted for brevity)
	db := database.DBConnect(dbConfig)
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	// Database migrations (if any)
	database.DBMigration(db)

	customValidator.SetDB(db)
	validate := customValidator.NewCustomValidator()

	// This function registers a custom "tag name" function.
	// It tells the validator to use the `json` tag value as the field name.
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		// Look for the json tag and split it by comma
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		// Use the struct field name as a fallback
		if name == "-" || name == "" {
			return fld.Name
		}
		return name
	})

	r := chi.NewRouter()

	// CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	userRepo := user.NewRepository(db)
	transactionRepo := transaction.NewRepository(db)

	// Init helpers
	jwtHelper := helper.NewJWTHelper(mainConfig.JWTSecret)

	authService := auth.NewService(userRepo, jwtHelper)
	userService := user.NewService(userRepo)
	transactionService := transaction.NewService(transactionRepo)

	// Middleware
	r.Use(c.Handler)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Register routes
	r.Route("/api/v1", func(r chi.Router) {
		auth.RegisterRoutes(r, validate, authService)
		r.Group(func(r chi.Router) {
			authMiddleware := middleware2.NewAuthMiddleware(jwtHelper)
			r.Use(authMiddleware.RequireAuth)
			user.RegisterRoutes(r, validate, userService)
			transaction.RegisterRoutes(r, validate, transactionService)
		})
	})

	// Start server
	port := ":" + fmt.Sprintf("%d", mainConfig.APIPort)
	fmt.Printf("Server starting on port %s\n", port)
	if err := http.ListenAndServe(port, r); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
