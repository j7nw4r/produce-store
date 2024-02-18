package models

import (
	"github.com/j7nw4r/produce-store/schemas"
)

type Produce struct {
	Code  string  `json:"code"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

func FromSchemaToResponse(s schemas.ProduceSchema) Produce {
	return Produce{
		Code:  s.Code,
		Name:  s.Name,
		Price: s.Price,
	}
}

func FromSchemasToResponses(schemas []schemas.ProduceSchema) []Produce {
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
