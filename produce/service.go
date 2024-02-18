package produce

import "database/sql"

type Service struct {
	db *sql.DB
}

func NewProduceService(db *sql.DB) Service {
	return Service{db: db}
}
