package main

import (
	"catetduit/internal/config"
	"catetduit/internal/database"
	"catetduit/internal/helper"
	"catetduit/internal/module/auth"
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

	// Database migrations (if any)
	database.DBMigration(db)

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

	// CORS
	c := cors.New(cors.Options{
		// GANTI INI dengan URL Next.js Anda.
		// Saat pengembangan (misalnya): "http://localhost:3000"
		// Saat produksi (misalnya): "https://app.domainanda.com"
		// Jika ingin mengizinkan SEMUA (Hanya untuk testing/dev!): []string{"*"}
		AllowedOrigins: []string{"http://localhost:3000"},

		// Metode yang diizinkan oleh API Anda
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},

		// Header yang diizinkan untuk dikirim oleh Next.js (misalnya, untuk Auth Token)
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},

		// Izinkan kredensial (penting jika Anda menggunakan Cookie/Session)
		AllowCredentials: true,

		// Maksimum waktu (detik) preflight OPTIONS request dapat di-cache oleh browser
		MaxAge: 300, // 5 menit
	})

	// Middleware
	r.Use(c.Handler)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	userRepo := user.NewRepository(db)

	// Init helpers
	jwtHelper := helper.NewJWTHelper(mainConfig.JWTSecret)

	authService := auth.NewService(userRepo, jwtHelper)

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
