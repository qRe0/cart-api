package migrations

import (
	"embed"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
)

var (
	//go:embed *.sql
	embedMigrations embed.FS
	migrationDir    = "."
)

type Migration struct {
	Database *sqlx.DB
}

type Migrator interface {
	Up() error
	Down() error
}

func NewMigrator(db *sqlx.DB) (Migrator, error) {
	goose.SetBaseFS(embedMigrations)
	err := goose.SetDialect("postgres")
	if err != nil {
		return nil, fmt.Errorf("failed to set dialect: %w", err)
	}

	return &Migration{Database: db}, nil
}

func (m *Migration) Up() error {
	err := goose.Up(m.Database.DB, migrationDir)
	if err != nil {
		return fmt.Errorf("failed to migrate db: %w", err)
	}

	log.Println("Database migrated successfully!")
	return nil
}

func (m *Migration) Down() error {
	err := goose.Down(m.Database.DB, migrationDir)
	if err != nil {
		return fmt.Errorf("failed to migrate db: %w", err)
	}

	log.Println("Database down successfully!")
	return nil
}
