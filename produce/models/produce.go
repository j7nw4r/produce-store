package models

import "github.com/google/uuid"

type ProduceSchema struct {
	Id    int
	Eid   uuid.UUID
	name  string
	price float32
}
