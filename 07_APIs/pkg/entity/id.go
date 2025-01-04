package entity

import "github.com/google/uuid"

type ID = uuid.UUID

// essa função pode ser usada por outras aplicacoes
func NewID() ID {
	return ID(uuid.New())
}

func ParseID(id string) (ID, error) {
	parsed, err := uuid.Parse(id)
	return ID(parsed), err
}
