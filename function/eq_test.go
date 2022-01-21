package function

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEq(t *testing.T) {
	type args struct {
		pred   string
		values []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should return eq string correctly",
			args: args{
				pred:   "test",
				values: []interface{}{"\"foo\"", "\"bar\""},
			},
			want: `eq(test, "foo", "bar")`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewEq(tt.args.pred, tt.args.values...).String()

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestNewNe(t *testing.T) {
	type args struct {
		pred   string
		values []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should return eq string correctly",
			args: args{
				pred:   "test",
				values: []interface{}{"\"foo\"", "\"bar\""},
			},
			want: `NOT eq(test, "foo", "bar")`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewNe(tt.args.pred, tt.args.values...).String()

			assert.Equal(t, tt.want, got)
		})
	}
}
