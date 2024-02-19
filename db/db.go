package db

import (
	"database/sql"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log/slog"
	_ "modernc.org/sqlite"
)

func NewInMemoryDB() (*sql.DB, error) {
	return newDB(":memory:")
}

func NewExternalDB(dbURL string) (*sql.DB, error) {
	return newDB(dbURL)
}

func newDB(source string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", source)
	if err != nil {
		slog.Error(err.Error())
		return nil, errors.New("could not create db")
	}

	m, err := migrate.New(
		"file://migrations",
		source,
	)
	if err != nil {
		if err != migrate.ErrNoChange {
			slog.Error(err.Error())
			return nil, errors.New("could not migrate db")
		}
	}

	if err := m.Up(); err != nil {
		slog.Error(err.Error())
		return nil, errors.New("could not migrate db")
	}

	return db, nil
}
