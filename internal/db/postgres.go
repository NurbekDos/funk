package db

import (
	"database/sql"
	"fmt"

	"github.com/NurbekDos/funk/internal/cfg"
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
	db, err := sql.Open("postgres", cfg.GetConfig().PgUrl)
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
