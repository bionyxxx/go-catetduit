package main

import (
	"catetduit/internal/config"
	"catetduit/internal/database"
	"catetduit/internal/module/auth"
	"catetduit/internal/module/user"
	customValidator "catetduit/internal/validator"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

var validate = validator.New()

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

	customValidator.SetDB(db)
	validate = customValidator.NewCustomValidator()

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

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	userRepo := user.NewRepository(db)
	authService := auth.NewService(userRepo)

	// Register routes
	r.Route("/api/v1", func(r chi.Router) {
		auth.RegisterRoutes(r, validate, authService)
		user.RegisterRoutes(r, db)
	})

	// Start server
	port := ":" + fmt.Sprintf("%d", mainConfig.APIPort)
	fmt.Printf("Server starting on port %s\n", port)
	if err := http.ListenAndServe(port, r); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
