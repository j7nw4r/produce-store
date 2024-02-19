package models

import (
	"github.com/j7nw4r/produce-store/schemas"
	"math"
)

// Produce Example
type Produce struct {
	Id    int     `json:"id" example:"1" format:"int32"`
	Code  string  `json:"code" example:"A12T-4GH7-QPL9-3N4M"`
	Name  string  `json:"name" example:"bannana"`
	Price float32 `json:"price" example:"3.32" format:"float32"`
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
			Id:    s.Id,
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
			Price: float32(roundFloat(float64(s.Price), 2)),
		}
		responses = append(responses, resp)
	}
	return responses
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
