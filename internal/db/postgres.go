package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
)

var (
	migrationsFolder string
	DB               *sql.DB
)

func init() {
	migrationsFolder = fmt.Sprintf("file://./internal/db/migrations/")
}

func ConnectToPostgres() error {
	db, err := sql.Open("postgres", getPostgresUrl())
	if err != nil {
		return fmt.Errorf("postgres: open connection error: %s", err.Error())
	}
	DB = db

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("postgres: ping error: %s", err.Error())
	}

	return nil
}

func Close() {
	err := DB.Close()
	if err != nil {
		fmt.Printf("postgres: closing DB error: %s\n", err.Error())
	}
}

func getPostgresUrl() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_DB_NAME"),
		os.Getenv("PG_SSL_MODE"),
	)
}
