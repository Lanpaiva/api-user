package models

import (
	"github.com/google/uuid"
)

type ID = uuid.UUID

func NewID() ID {
	return ID(uuid.New())
}

func ParseID(s string) (ID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		panic(err)
	}
	return ID(id), err
}
