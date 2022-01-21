package dbuilder

import (
	"fmt"
	"strings"

	"github.com/u-next/dbuilder/function"
)

// Expression is consist of either single or multiple functions.
type Expression struct {
	fns []function.Function
}

// Build return expression with functions in string. Must be in conjunction with `AND`. e.g.
//   - (fn)
//   - (fn AND fn)
func (expr *Expression) Build() string {
	switch {
	case len(expr.fns) <= 0:
		return ""
	case len(expr.fns) == 1:
		return expr.fns[0].String()
	default:
		ret := make([]string, 0, len(expr.fns))
		for _, fn := range expr.fns {
			ret = append(ret, fn.String())
		}

		return fmt.Sprintf("(%s)", strings.Join(ret, Conjunction(ConjunctionAnd).String()))
	}
}
