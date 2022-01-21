package function

import "fmt"

// Regular Expressions
// Syntax Examples:
//   - regexp(predicate, /regular-expression/)
//   - regexp(predicate, /regular-expression/i) (case insensitive)
type Regexp struct {
	pred  string
	value interface{}
}

func NewRegexp(pred string, value interface{}) Regexp {
	return Regexp{pred: pred, value: value}
}

func (op Regexp) String() string {
	return fmt.Sprintf("regexp(%s, %v)", op.pred, op.value)
}
