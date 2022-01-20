package function

import "fmt"

// allofterms
// Syntax Example: allofterms(predicate, "space-separated term list")
type Allofterms struct {
	pred  string
	value interface{}
}

func NewAllofterms(pred string, value interface{}) Allofterms {
	return Allofterms{pred: pred, value: value}
}

func (op Allofterms) String() string {
	return fmt.Sprintf("allofterms(%s, %v)", op.pred, op.value)
}
