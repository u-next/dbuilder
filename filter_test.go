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

func TestFilter_Has(t *testing.T) {
	type fields struct {
		conj  Conjunction
		exprs []*Expression
	}
	type args struct {
		pred string
		not  bool
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

			got := f.Has(tt.args.pred, tt.args.not).Build()

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFilter_Type(t *testing.T) {
	type fields struct {
		conj  Conjunction
		exprs []*Expression
	}
	type args struct {
		targetType string
		not        bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "should return type() correctly",
			fields: fields{
				conj: ConjunctionAnd,
			},
			args: args{
				targetType: "Foo",
			},
			want: "@filter(type(Foo))",
		},
		{
			name: "should return NOT type() correctly",
			fields: fields{
				conj: ConjunctionAnd,
			},
			args: args{
				targetType: "Foo",
				not:        true,
			},
			want: "@filter(NOT type(Foo))",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Filter{
				conj:  tt.fields.conj,
				exprs: tt.fields.exprs,
			}
			got := f.Type(tt.args.targetType, tt.args.not).Build()

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFilter_BuildRaw(t *testing.T) {
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
			name: "should return clause correctly",
			fields: fields{
				exprs: []*Expression{
					{
						fns: []function.Function{
							function.NewGt("original_price", 500),
						},
					},
				},
			},
			want: `gt(original_price, 500)`,
		},
		{
			name: "should return clause with multiple expressions correctly",
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
					(&CustomQueryOperator{Expression: pointerizer.S(`(lt(media.original_release_date, "1977-01-01") OR gt(media.original_price, 10))`)}).execute("ignore"),
					NewFilter(ConjunctionAnd).Apply("other", &StringQueryOperator{Eq: pointerizer.S("foo")}).ToCustomQueryOperator().execute("ignore"),
				},
			},
			want: `gt(original_price, 500) AND gt(popularity, 0.5) AND (lt(media.original_release_date, "1977-01-01") OR gt(media.original_price, 10)) AND eq(other, "foo")`,
		},
		{
			name: "should return clause with multiple functions in one expression correctly",
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
			want: `(gt(original_price, 500) AND lt(original_price, 1000)) OR gt(popularity, 0.5)`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Filter{
				conj:  tt.fields.conj,
				exprs: tt.fields.exprs,
			}

			got := f.BuildRaw()

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

func TestFilter_ToCustomQueryOperator(t *testing.T) {
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
			name: "should return @filter content correctly",
			fields: fields{
				exprs: []*Expression{
					{
						fns: []function.Function{
							function.NewGt("original_price", 500),
						},
					},
				},
			},
			want: `gt(original_price, 500)`,
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
			want: `(gt(original_price, 500) AND gt(popularity, 0.5))`,
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
			want: `((gt(original_price, 500) AND lt(original_price, 1000)) OR gt(popularity, 0.5))`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Filter{
				conj:  tt.fields.conj,
				exprs: tt.fields.exprs,
			}

			got := *f.ToCustomQueryOperator().Expression

			assert.Equal(t, tt.want, got)
		})
	}
}
