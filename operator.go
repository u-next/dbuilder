package dbuilder

type QueryOperatorable interface {
	// execute returns expression which contains functions for equality of a predicate
	// to a value or find in a list of values.
	execute(pred string) *Expression
}
