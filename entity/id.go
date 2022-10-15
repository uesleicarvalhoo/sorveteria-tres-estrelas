package entity

import (
	"errors"

	"github.com/google/uuid"
)

var ErrInvalidID = errors.New("invalid id")

type ID uuid.UUID

func (id ID) String() string {
	return uuid.UUID(id).String()
}

func StringToID(id string) (ID, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return ID{}, ErrInvalidID
	}

	return ID(uid), nil
}

func NewID() ID {
	return ID(uuid.New())
}
