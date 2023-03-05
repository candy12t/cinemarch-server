package entity

import (
	"reflect"
	"testing"
)

func TestCondition_Build(t *testing.T) {
	tests := []struct {
		name       string
		conditions Conditions
		wantQuery  string
		wantArgs   []any
	}{
		{
			name: "build condition query",
			conditions: Conditions{
				{
					Query: "title LIKE ?",
					Arg:   "%title%",
				},
				{
					Query: "release_status = ?",
					Arg:   "NOW OPEN",
				},
			},
			wantQuery: "WHERE title LIKE ? AND release_status = ?",
			wantArgs:  []any{"%title%", "NOW OPEN"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotQuery, gotArgs := tt.conditions.Build()
			if gotQuery != tt.wantQuery {
				t.Errorf("Conditions.Build() got query is %v, wantQuery is %v", gotQuery, tt.wantQuery)
			}

			if len(gotArgs) != len(tt.wantArgs) {
				t.Errorf("Conditions.Build() got args size is %v, wantArgs size is %v", len(gotQuery), len(tt.wantQuery))
			}

			if !reflect.DeepEqual(gotArgs, tt.wantArgs) {
				t.Errorf("Conditions.Build() got args is %v, wantArgs is %v", gotArgs, tt.wantArgs)
			}
		})
	}
}
