package function

import "fmt"

// Has determines if a node has a particular predicate.
// https://dgraph.io/docs/query-language/functions/#has
type Has struct {
	pred string
}

func NewHas(pred string) Has {
	return Has{pred: pred}
}

func (op Has) String() string {
	return fmt.Sprintf("has(%s)", op.pred)
}
