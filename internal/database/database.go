package database

import (
	"catetduit/internal/config"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
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
