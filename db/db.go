package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
	"log/slog"
)

func NewInMemoryDB() (*sql.DB, error) {
	return newDB(":memory:")
}

func NewExternalDB(ctx context.Context, dbURL string) (*sql.DB, error) {
	db, err := newDB(dbURL)
	if err != nil {
		return nil, err
	}

	if err := db.PingContext(ctx); err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	return db, nil
}

func newDB(dbName string) (*sql.DB, error) {
	connStr := fmt.Sprintf("file:%s?cache=shared", dbName)
	db, err := sql.Open("sqlite3", connStr)
	if err != nil {
		slog.Error(err.Error())
		return nil, errors.New("could not create db")
	}

	migrationConnStr := fmt.Sprintf("sqlite3://%s", dbName)
	m, err := migrate.New(
		"file://migrations",
		migrationConnStr,
	)
	if err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			slog.Error(err.Error())
			return nil, errors.New("could not migrate db")
		}
	}

	if err := m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			slog.Error(err.Error())
			return nil, errors.New("could not migrate db")
		}
	}

	return db, nil
}
