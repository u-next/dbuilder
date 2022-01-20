package function

import "fmt"

// anyofterms
// Syntax Example: anyofterms(predicate, "space-separated term list")
type Anyofterms struct {
	pred  string
	value interface{}
}

func NewAnyofterms(pred string, value interface{}) Anyofterms {
	return Anyofterms{pred: pred, value: value}
}

func (op Anyofterms) String() string {
	return fmt.Sprintf("anyofterms(%s, %v)", op.pred, op.value)
}
