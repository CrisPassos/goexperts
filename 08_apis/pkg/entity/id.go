package entity

import "github.com/google/uuid"

// esse package pode ser utilizado por outras pessoas
type ID = uuid.UUID

func NewID() ID {
	return uuid.New()
}

func ParseID(s string) (ID, error) {
	id, err := uuid.Parse(s)
	return ID(id), err
}
