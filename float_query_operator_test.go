package dbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/u-next/dbuilder/function"
	"github.com/u-next/dbuilder/pointerizer"
)

func TestFloatQueryOperator_execute(t *testing.T) {
	type fields struct {
		Eq  *float64
		Ne  *float64
		Gt  *float64
		Gte *float64
		Lt  *float64
		Lte *float64
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
			name: "should return expressions correctly",
			fields: fields{
				Eq:  pointerizer.F64(0.1),
				Ne:  pointerizer.F64(0.2),
				Gt:  pointerizer.F64(0.3),
				Gte: pointerizer.F64(0.4),
				Lt:  pointerizer.F64(0.5),
				Lte: pointerizer.F64(0.6),
			},
			args: args{pred: "foo"},
			want: &Expression{
				fns: []function.Function{
					function.NewEq("foo", 0.1),
					function.NewNe("foo", 0.2),
					function.NewGt("foo", 0.3),
					function.NewGte("foo", 0.4),
					function.NewLt("foo", 0.5),
					function.NewLte("foo", 0.6),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := &FloatQueryOperator{
				Eq:  tt.fields.Eq,
				Ne:  tt.fields.Ne,
				Gt:  tt.fields.Gt,
				Gte: tt.fields.Gte,
				Lt:  tt.fields.Lt,
				Lte: tt.fields.Lte,
			}
			got := op.execute(tt.args.pred)

			assert.Equal(t, tt.want, got)
		})
	}
}
