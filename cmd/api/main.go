package main

import (
	"catetduit/internal/config"
	"catetduit/internal/database"
	"catetduit/internal/module/auth"
	"catetduit/internal/module/user"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Register routes
	r.Route("/api/v1", func(r chi.Router) {
		user.RegisterRoutes(r, db)
		auth.RegisterRoutes(r)
	})

	// Start server
	port := ":" + fmt.Sprintf("%d", mainConfig.APIPort)
	fmt.Printf("Server starting on port %s\n", port)
	if err := http.ListenAndServe(port, r); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
