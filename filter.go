package dbuilder

import (
	"fmt"
	"strings"

	"github.com/u-next/dbuilder/function"
)

// Filter is a builder to construct filter directive
type Filter struct {
	conj  Conjunction
	exprs []*Expression
}

// NewFilter return a builder to construct filter directive
func NewFilter(conj Conjunction) *Filter {
	return &Filter{conj: conj}
}

// Apply applies query functions to given predicate.
//
// @example: "media.popularity" should greater than 0.5
//   -> Apply("media.popularity", &FloatQueryOperator{Gt: pointerizer.F64(0.5)})
//
func (f *Filter) Apply(pred string, op QueryOperatorable) *Filter {
	expr := op.execute(pred)

	if expr != nil {
		f.exprs = append(f.exprs, expr)
	}

	return f
}

func (f *Filter) ApplyList(pred string, ops []QueryOperatorable) *Filter {
	for _, op := range ops {
		f.Apply(pred, op)
	}

	return f
}

// Has add has() function to determine if a node has a particular predicate
func (f *Filter) Has(pred string) *Filter {
	f.exprs = append(f.exprs, &Expression{
		fns: []function.Function{
			function.NewHas(pred),
		},
	})

	return f
}

// Build returns @filter directive e.g.
//
//   - @filter(gt(original_price, 500))
//   - @filter(gt(original_price, 500) AND gt(popularity, 0.5))
//   - @filter((gt(original_price, 500) AND lt(original_price, 1000)) OR gt(popularity, 0.5))
//
func (f *Filter) Build() string {
	if len(f.exprs) <= 0 {
		return ""
	}

	ret := make([]string, 0, len(f.exprs))

	for _, expr := range f.exprs {
		if cl := expr.Build(); len(cl) > 0 {
			ret = append(ret, cl)
		}
	}

	return fmt.Sprintf("@filter(%s)", strings.Join(ret, f.conj.String()))
}
