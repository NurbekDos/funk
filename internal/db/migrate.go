package db

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate"
)

func RunMigrations() error {
	migrateInstance, err := migrate.New(migrationsFolder, getPostgresUrl())
	if err != nil {
		return fmt.Errorf("golang-migrate: create migrate instance error: %s", err.Error())
	}

	err = migrateInstance.Up()
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("RunMigrations: Changes not found")
			return nil
		}

		return fmt.Errorf("golang-migrate: up migration error: %s", err.Error())
	}

	fmt.Println("RunMigrations: migrated")
	return nil
}
