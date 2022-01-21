package dbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConjunction_String(t *testing.T) {
	tests := []struct {
		name string
		conj Conjunction
		want string
	}{
		{
			name: "should return OR correctly",
			conj: ConjunctionOr,
			want: " OR ",
		},
		{
			name: "should return AND correctly",
			conj: ConjunctionAnd,
			want: " AND ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := tt.conj.String()

			assert.Equal(t, tt.want, got)
		})
	}
}
