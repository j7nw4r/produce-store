package models

import (
	"github.com/j7nw4r/produce-store/schemas"
)

type Produce struct {
	Id    int     `json:"id"`
	Code  string  `json:"code"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

func FromProduceSchemaToProduce(s schemas.ProduceSchema) Produce {
	return Produce{
		Id:    s.Id,
		Code:  s.Code,
		Name:  s.Name,
		Price: s.Price,
	}
}

func FromProduceSchemasToProduces(ss []schemas.ProduceSchema) []Produce {
	if ss == nil {
		return []Produce{}
	}

	responses := []Produce{}
	for _, s := range ss {
		resp := Produce{
			Code:  s.Code,
			Name:  s.Name,
			Price: s.Price,
		}
		responses = append(responses, resp)
	}
	return responses
}

func FromProduceToProduceSchema(p Produce) *schemas.ProduceSchema {
	return &schemas.ProduceSchema{
		Id:    p.Id,
		Code:  p.Code,
		Name:  p.Name,
		Price: p.Price,
	}
}

func FromProducesToProduceSchemas(ss []Produce) []schemas.ProduceSchema {
	if ss == nil {
		return []schemas.ProduceSchema{}
	}

	responses := []schemas.ProduceSchema{}
	for _, s := range ss {
		resp := schemas.ProduceSchema{
			Code:  s.Code,
			Name:  s.Name,
			Price: s.Price,
		}
		responses = append(responses, resp)
	}
	return responses
}
