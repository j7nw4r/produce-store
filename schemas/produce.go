package schemas

import "github.com/google/uuid"

type ProduceSchema struct {
	Id    int
	Eid   uuid.UUID
	Code  string
	Name  string
	Price float32
}
