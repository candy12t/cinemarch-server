package entity

type Prefecture string

func NewPrefecture(prefecture string) (Prefecture, error) {
	return Prefecture(prefecture), nil
}

func (p Prefecture) String() string {
	return string(p)
}
