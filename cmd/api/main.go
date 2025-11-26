package main

import (
	"catetduit/internal/config"
	"catetduit/internal/database"
	"catetduit/internal/helper"
	middleware2 "catetduit/internal/middleware"
	"catetduit/internal/module/auth"
	"catetduit/internal/module/oauth"
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

var (
	mainConfig   *config.Config
	dbConfig     *config.DatabaseConfig
	db           *sqlx.DB
	oauth2Config *config.OAuth2Config
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	mainConfig = config.NewConfig()
	dbConfig = config.NewDatabaseConfig()

	db = database.DBConnect(dbConfig)

	database.DBMigration(db)

	customValidator.SetDB(db)

	oauth2Config = config.NewOAuth2Config()
}

func main() {
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	validate := customValidator.NewCustomValidator()

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" || name == "" {
			return fld.Name
		}
		return name
	})

	r := chi.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://127.0.0.1:3000", "http://localhost:3000", "https://catetduit.duttafachrezy.my.id"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	userRepo := user.NewRepository(db)
	transactionRepo := transaction.NewRepository(db)

	jwtHelper := helper.NewJWTHelper(mainConfig.JWTSecret, mainConfig.JWTExpiredInHour, mainConfig.JWTRefreshExpiredInHour)

	oauth2Service := oauth.NewService(oauth2Config, userRepo, jwtHelper)
	authService := auth.NewService(userRepo, jwtHelper, *oauth2Config)
	userService := user.NewService(userRepo)
	transactionService := transaction.NewService(transactionRepo)

	r.Use(c.Handler)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Routes

	// API v1 Routes
	r.Route("/api/v1", func(r chi.Router) {
		auth.RegisterRoutes(r, validate, authService, oauth2Config)
		r.Group(func(r chi.Router) {
			authMiddleware := middleware2.NewAuthMiddleware(jwtHelper)
			r.Use(authMiddleware.RequireAuth)
			user.RegisterRoutes(r, validate, userService)
			transaction.RegisterRoutes(r, validate, transactionService)
		})
	})

	// Oauth2 Routes
	r.Route("/oauth2", func(r chi.Router) {
		oauth.RegisterRoutes(r, validate, authService, oauth2Service, oauth2Config)
	})

	port := ":" + fmt.Sprintf("%d", mainConfig.APIPort)
	fmt.Printf("Server starting on port %s\n", port)
	if err := http.ListenAndServe(port, r); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
