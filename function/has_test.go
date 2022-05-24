package function

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHas_String(t *testing.T) {
	type fields struct {
		pred   string
		invert bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "should return has() correctly",
			fields: fields{
				pred: "foo",
			},
			want: "has(foo)",
		},
		{
			name: "should return NOT has() correctly",
			fields: fields{
				pred:   "foo",
				invert: true,
			},
			want: "NOT has(foo)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := Has{
				pred:   tt.fields.pred,
				invert: tt.fields.invert,
			}

			got := op.String()

			assert.Equal(t, tt.want, got)
		})
	}
}
