package db

import (
	"database/sql"
	"errors"
	"log/slog"
)

func init() {
	// TODO: Do db migrations
}

func NewDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		slog.Error(err.Error())
		return nil, errors.New("could not create db")
	}
	return db, nil
}
