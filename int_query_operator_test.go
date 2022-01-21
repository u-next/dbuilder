package dbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/u-next/dbuilder/function"
	"github.com/u-next/dbuilder/pointerizer"
)

func TestIntQueryOperator_execute(t *testing.T) {
	type fields struct {
		Eq  *int
		Ne  *int
		Gt  *int
		Gte *int
		Lt  *int
		Lte *int
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
				Eq:  pointerizer.I(100),
				Ne:  pointerizer.I(200),
				Gt:  pointerizer.I(300),
				Gte: pointerizer.I(400),
				Lt:  pointerizer.I(500),
				Lte: pointerizer.I(600),
			},
			args: args{pred: "foo"},
			want: &Expression{
				fns: []function.Function{
					function.NewEq("foo", 100),
					function.NewNe("foo", 200),
					function.NewGt("foo", 300),
					function.NewGte("foo", 400),
					function.NewLt("foo", 500),
					function.NewLte("foo", 600),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := &IntQueryOperator{
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
