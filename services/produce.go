package services

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/j7nw4r/produce-store/schemas"
)

type ProduceService struct {
	db *sql.DB
}

func NewProduceService(db *sql.DB) ProduceService {
	return ProduceService{db: db}
}

func (ps ProduceService) GetProduce(eid uuid.UUID) (*schemas.ProduceSchema, error) {
	return nil, errors.New("not implemented")
}

func (ps ProduceService) SearchProduce(name string) ([]schemas.ProduceSchema, error) {
	return nil, errors.New("not implemented")
}

func (ps ProduceService) StoreProduce(p schemas.ProduceSchema) (*schemas.ProduceSchema, error) {
	return nil, errors.New("not implemented")
}
