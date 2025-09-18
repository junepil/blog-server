package database

import (
	"database/sql"
	"fmt"

	"blog-api/internal/config"

	_ "github.com/lib/pq"
)

func Connect(config config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName, config.SSLMode)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
