package models

import (
	"errors"
	"github.com/google/uuid"
	"github.com/j7nw4r/produce-store/schemas"
	"log/slog"
)

type Produce struct {
	Id    string  `json:"id"`
	Code  string  `json:"code"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

func FromProduceSchemaToProduce(s schemas.ProduceSchema) Produce {
	return Produce{
		Id:    s.Eid.String(),
		Code:  s.Code,
		Name:  s.Name,
		Price: s.Price,
	}
}

func FromProduceSchemasToProduces(schemas []schemas.ProduceSchema) []Produce {
	var responses []Produce
	for _, s := range schemas {
		resp := Produce{
			Code:  s.Code,
			Name:  s.Name,
			Price: s.Price,
		}
		responses = append(responses, resp)
	}
	return responses
}

func FromProduceToProduceSchema(p Produce) (*schemas.ProduceSchema, error) {
	eid, err := uuid.Parse(p.Id)
	if err != nil {
		slog.Error(err.Error())
		return nil, errors.New("could not parse id ")
	}

	return &schemas.ProduceSchema{
		Eid:   eid,
		Code:  p.Code,
		Name:  p.Name,
		Price: p.Price,
	}, nil
}
