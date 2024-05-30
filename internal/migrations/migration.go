package migrations

import (
	"embed"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
)

//go:embed *.sql
var embedMigrations embed.FS

func MigrationUp(db *sqlx.DB) error {
	migrationDir := "."

	goose.SetBaseFS(embedMigrations)

	err := goose.SetDialect("postgres")
	if err != nil {
		return fmt.Errorf("failed to set dialect: %w", err)
	}

	err = goose.Up(db.DB, migrationDir)
	if err != nil {
		return fmt.Errorf("failed to migrate db: %w", err)
	}

	log.Println("Database migrated successfully!")
	return nil
}
