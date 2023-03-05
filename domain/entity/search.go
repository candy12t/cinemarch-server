package entity

import "fmt"

type Condition struct {
	Query string
	Arg   any
}

type Conditions []Condition

func (cs Conditions) Build() (string, []any) {
	query := fmt.Sprintf("WHERE %s", cs[0].Query)

	args := make([]any, 0, len(cs))
	args = append(args, cs[0].Arg)

	for _, c := range cs[1:] {
		query = fmt.Sprintf("%s AND %s", query, c.Query)
		args = append(args, c.Arg)
	}
	return query, args
}
