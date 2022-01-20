package function

import "fmt"

// IE: gt(), ge(), lt(), le()
// https://dgraph.io/docs/query-language/functions/#less-than-less-than-or-equal-to-greater-than-and-greater-than-or-equal-to
type Gt struct {
	pred  string
	value interface{}
}

func NewGt(pred string, value interface{}) Gt {
	return Gt{pred: pred, value: value}
}

func (op Gt) String() string {
	return fmt.Sprintf("gt(%s, %v)", op.pred, op.value)
}
