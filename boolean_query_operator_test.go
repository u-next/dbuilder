package dbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/u-next/dbuilder/function"
	"github.com/u-next/dbuilder/pointerizer"
)

func TestBooleanQueryOperator_execute(t *testing.T) {
	type fields struct {
		Eq *bool
	}
	type args struct {
		pred string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Expression
	}{
		{
			name: "should return true correctly",
			fields: fields{
				Eq: pointerizer.B(true),
			},
			args: args{pred: "foo"},
			want: &Expression{
				fns: []function.Function{
					function.NewEq("foo", "true"),
				},
			},
		},
		{
			name: "should return false correctly",
			fields: fields{
				Eq: pointerizer.B(false),
			},
			args: args{pred: "foo"},
			want: &Expression{
				fns: []function.Function{
					function.NewEq("foo", "false"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := &BooleanQueryOperator{
				Eq: tt.fields.Eq,
			}

			got := op.execute(tt.args.pred)

			assert.Equal(t, tt.want, got)
		})
	}
}
