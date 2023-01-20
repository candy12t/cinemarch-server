package entity

type ScreeningTypeName string

type ScreeningType struct {
	ID   UUID
	Name ScreeningTypeName
}

func NewScreeningType(name ScreeningTypeName) *ScreeningType {
	return &ScreeningType{
		ID:   NewUUID(),
		Name: name,
	}
}

func NewScreeningTypeName(name string) (ScreeningTypeName, error) {
	if name == "" || len([]rune(name)) > 255 {
		return "", ErrInvalidLengthScreeningTypeName
	}
	return ScreeningTypeName(name), nil
}

func (stn ScreeningTypeName) String() string {
	return string(stn)
}
