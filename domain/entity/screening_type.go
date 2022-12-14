package entity

import "time"

type ScreeningType struct {
	ID        UUID
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewScreeningType(_type string) *ScreeningType {
	now := time.Now()
	return &ScreeningType{
		ID:        NewUUID(),
		Type:      _type,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
