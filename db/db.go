package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DBConfig struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func NewPostgresStorage(cfg DBConfig) (*sql.DB, error) {
	var connString = fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	// var connString = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection ⛔: %w", err)
	}

	//Test db connection with a ping
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database ⛔: %w", err)
	}

	fmt.Println("Connected to database successfully ✅ ...")

	return db, nil
}
