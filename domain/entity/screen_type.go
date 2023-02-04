package entity

type ScreenTypeName string

type ScreenType struct {
	ID   UUID
	Name ScreenTypeName
}

func NewScreenType(name ScreenTypeName) *ScreenType {
	return &ScreenType{
		ID:   NewUUID(),
		Name: name,
	}
}

func NewScreenTypeName(name string) (ScreenTypeName, error) {
	if name == "" || len([]rune(name)) > 255 {
		return "", ErrInvalidLengthScreenTypeName
	}
	return ScreenTypeName(name), nil
}

func (stn ScreenTypeName) String() string {
	return string(stn)
}
