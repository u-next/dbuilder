package dbuilder

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/u-next/dbuilder/function"
	"github.com/u-next/dbuilder/pointerizer"
)

func TestTimeQueryOperator_execute(t *testing.T) {
	type fields struct {
		Eq  *time.Time
		Ne  *time.Time
		Gt  *time.Time
		Gte *time.Time
		Lt  *time.Time
		Lte *time.Time
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
				Eq:  pointerizer.T(time.Date(2021, 1, 1, 17, 30, 0, 0, time.UTC)),
				Ne:  pointerizer.T(time.Date(2021, 2, 1, 17, 30, 0, 0, time.UTC)),
				Gt:  pointerizer.T(time.Date(2021, 3, 1, 17, 30, 0, 0, time.UTC)),
				Gte: pointerizer.T(time.Date(2021, 4, 1, 17, 30, 0, 0, time.UTC)),
				Lt:  pointerizer.T(time.Date(2021, 5, 1, 17, 30, 0, 0, time.UTC)),
				Lte: pointerizer.T(time.Date(2021, 6, 1, 17, 30, 0, 0, time.UTC)),
			},
			args: args{pred: "foo"},
			want: &Expression{
				fns: []function.Function{
					function.NewEq("foo", "\"2021-01-01T17:30:00Z\""),
					function.NewNe("foo", "\"2021-02-01T17:30:00Z\""),
					function.NewGt("foo", "\"2021-03-01T17:30:00Z\""),
					function.NewGte("foo", "\"2021-04-01T17:30:00Z\""),
					function.NewLt("foo", "\"2021-05-01T17:30:00Z\""),
					function.NewLte("foo", "\"2021-06-01T17:30:00Z\""),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := &TimeQueryOperator{
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
