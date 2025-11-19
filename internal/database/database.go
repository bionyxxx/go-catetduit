package database

import (
	"catetduit/internal/config"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
)

func DBConnect(databaseConfig *config.DatabaseConfig) *sqlx.DB {
	db, err := sqlx.Connect("pgx",
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			databaseConfig.Host,
			databaseConfig.Port,
			databaseConfig.User,
			databaseConfig.Password,
			databaseConfig.DBName,
		),
	)

	if err != nil {
		panic(err)
	}

	// Ping to verify connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func DBMigration(database *sqlx.DB) {
	goose.SetTableName(os.Getenv("DB_MIGRATIONS_TABLE"))

	err := goose.Up(database.DB, "./migrations")
	if err != nil {
		panic(err)
	}
}
