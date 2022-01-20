package dbuilder

// Conjunction specify conjunction implementation in dgraph
type Conjunction int

const (
	ConjunctionAnd Conjunction = iota + 1 // (A and B)
	ConjunctionOr                         // (A or B)
)

// String return conjunction in string. Return `AND` as fallback.
func (conj Conjunction) String() string {
	switch conj {
	case ConjunctionOr:
		return " OR "
	default:
		return " AND "
	}
}
