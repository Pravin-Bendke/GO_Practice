package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func ConnectDB() (*sql.DB, error) {

	databaseURL := os.Getenv("DATABASE_URL")

	if databaseURL == "" {
		return nil, fmt.Errorf("DATABASE_URL is not set")
	}

	db, err := sql.Open("pgx", databaseURL)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
