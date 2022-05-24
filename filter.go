package dbuilder

import (
	"fmt"
	"strings"

	"github.com/u-next/dbuilder/pointerizer"

	"github.com/u-next/dbuilder/function"
)

// Filter is a builder to construct filter directive
//
// https://github.com/u-next/dbuilder
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

// Has add has() function to determine if a node has a particular predicate
// `not` controls whether a NOT operator should be appended to the front to
// invert the results
func (f *Filter) Has(pred string, not bool) *Filter {
	f.exprs = append(f.exprs, &Expression{
		fns: []function.Function{
			function.NewHas(pred, not),
		},
	})

	return f
}

// Type add type() function to determine if a node belongs to particular type.
// `not` controls whether a NOT operator should be appended to the front to
// invert the results
func (f *Filter) Type(targetType string, not bool) *Filter {
	f.exprs = append(f.exprs, &Expression{
		fns: []function.Function{
			function.NewType(targetType, not),
		},
	})

	return f
}

// BuildRaw returns the content of the @filter directive e.g.
//
//   - gt(original_price, 500)
//   - (gt(original_price, 500) AND gt(popularity, 0.5))
//   - ((gt(original_price, 500) AND lt(original_price, 1000)) OR gt(popularity, 0.5))
func (f *Filter) BuildRaw() string {
	if len(f.exprs) <= 0 {
		return ""
	}

	ret := make([]string, 0, len(f.exprs))

	for _, expr := range f.exprs {
		if cl := expr.Build(); len(cl) > 0 {
			ret = append(ret, cl)
		}
	}

	return strings.Join(ret, f.conj.String())
}

// Build returns @filter directive e.g.
//
//   - @filter(gt(original_price, 500))
//   - @filter(gt(original_price, 500) AND gt(popularity, 0.5))
//   - @filter((gt(original_price, 500) AND lt(original_price, 1000)) OR gt(popularity, 0.5))
func (f *Filter) Build() string {
	if len(f.exprs) <= 0 {
		return ""
	}

	return fmt.Sprintf("@filter(%s)", f.BuildRaw())
}

// ToCustomQueryOperator returns an Operator with the content of the @filter directive e.g.
//
//   - gt(original_price, 500)
//   - (gt(original_price, 500) AND gt(popularity, 0.5))
//   - ((gt(original_price, 500) AND lt(original_price, 1000)) OR gt(popularity, 0.5))
//
// This can be used to build partial filters to be included in other filters, e.g.:
// partialFilter := &dbuilder.NewFilter(dbuilder.ConjunctionAnd).
//		Apply("other", &dbuilder.StringQueryOperator{Eq: pointerizer.S("foo")}).
// 		ToCustomQueryOperator()
// dbuilder.NewFilter(dbuilder.ConjunctionAnd).
//		Apply("popularity", &dbuilder.FloatQueryOperator{Eq: pointerizer.F64(0.5)}).
//		Apply("", partialFilter)
func (f *Filter) ToCustomQueryOperator() *CustomQueryOperator {
	str := f.BuildRaw()
	if len(str) <= 0 {
		return &CustomQueryOperator{Expression: &str}
	}

	// keep the parenthesis if we have more than one expression
	if len(f.exprs) > 1 {
		str = fmt.Sprintf("(%s)", str)
	}

	return &CustomQueryOperator{Expression: pointerizer.S(str)}
}
