package dbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/u-next/dbuilder/function"
	"github.com/u-next/dbuilder/pointerizer"
)

func TestFilter_Apply(t *testing.T) {
	type fields struct {
		conj  Conjunction
		exprs []*Expression
	}
	type args struct {
		pred string
		op   QueryOperatorable
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "should apply operator correctly",
			fields: fields{
				conj: ConjunctionAnd,
			},
			args: args{
				pred: "media.popularity",
				op:   &FloatQueryOperator{Gt: pointerizer.F64(0.5)},
			},
			want: "@filter(gt(media.popularity, 0.5))",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Filter{
				conj:  tt.fields.conj,
				exprs: tt.fields.exprs,
			}

			got := f.Apply(tt.args.pred, tt.args.op).Build()

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFilter_ApplyList(t *testing.T) {
	type fields struct {
		conj  Conjunction
		exprs []*Expression
	}
	type args struct {
		pred string
		ops  []QueryOperatorable
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "should apply operators correctly",
			fields: fields{
				conj: ConjunctionAnd,
			},
			args: args{
				pred: "media.popularity",
				ops: []QueryOperatorable{
					&FloatQueryOperator{Gt: pointerizer.F64(0.5)},
				},
			},
			want: "@filter(gt(media.popularity, 0.5))",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Filter{
				conj:  tt.fields.conj,
				exprs: tt.fields.exprs,
			}

			got := f.ApplyList(tt.args.pred, tt.args.ops).Build()

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFilter_Has(t *testing.T) {
	type fields struct {
		conj  Conjunction
		exprs []*Expression
	}
	type args struct {
		pred string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "should apply operators correctly",
			fields: fields{
				conj: ConjunctionAnd,
			},
			args: args{
				pred: "media.popularity",
			},
			want: "@filter(has(media.popularity))",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Filter{
				conj:  tt.fields.conj,
				exprs: tt.fields.exprs,
			}

			got := f.Has(tt.args.pred).Build()

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFilter_Build(t *testing.T) {
	type fields struct {
		conj  Conjunction
		exprs []*Expression
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "should return empty string if no expressions",
			fields: fields{},
			want:   "",
		},
		{
			name: "should return @filter clause correctly",
			fields: fields{
				exprs: []*Expression{
					{
						fns: []function.Function{
							function.NewGt("original_price", 500),
						},
					},
				},
			},
			want: `@filter(gt(original_price, 500))`,
		},
		{
			name: "should return @filter clause with multiple expressions correctly",
			fields: fields{
				exprs: []*Expression{
					{
						fns: []function.Function{
							function.NewGt("original_price", 500),
						},
					},
					{
						fns: []function.Function{
							function.NewGt("popularity", 0.5),
						},
					},
				},
			},
			want: `@filter(gt(original_price, 500) AND gt(popularity, 0.5))`,
		},
		{
			name: "should return @filter clause with multiple functions in one expression correctly",
			fields: fields{
				conj: ConjunctionOr,
				exprs: []*Expression{
					{
						fns: []function.Function{
							function.NewGt("original_price", 500),
							function.NewLt("original_price", 1000),
						},
					},
					{
						fns: []function.Function{
							function.NewGt("popularity", 0.5),
						},
					},
				},
			},
			want: `@filter((gt(original_price, 500) AND lt(original_price, 1000)) OR gt(popularity, 0.5))`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Filter{
				conj:  tt.fields.conj,
				exprs: tt.fields.exprs,
			}

			got := f.Build()

			assert.Equal(t, tt.want, got)
		})
	}
}
