package entity

type ScreeningType struct {
	ID   UUID
	Type string
}

func NewScreeningType(_type string) *ScreeningType {
	return &ScreeningType{
		ID:   NewUUID(),
		Type: _type,
	}
}
