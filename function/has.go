package function

import "fmt"

// Has determines if a node has a particular predicate.
// https://dgraph.io/docs/query-language/functions/#has
type Has struct {
	pred string

	// invert controls whether a NOT operator should be appended to the front to invert the results
	invert bool
}

func NewHas(pred string, invert bool) Has {
	return Has{pred: pred, invert: invert}
}

func (op Has) String() string {
	cls := fmt.Sprintf("has(%s)", op.pred)

	if op.invert {
		return fmt.Sprintf("NOT %s", cls)
	}

	return cls
}
