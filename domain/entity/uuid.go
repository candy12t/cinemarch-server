package entity

import "github.com/google/uuid"

type UUID string

var NewUUID = func() UUID {
	return UUID(uuid.New().String())
}

func (u UUID) String() string {
	return string(u)
}
