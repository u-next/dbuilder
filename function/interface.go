package function

// Function allows filtering based on properties of nodes or variables
// https://dgraph.io/docs/query-language/functions/
type Function interface {
	// String converts function AST to string expression in dgraph
	String() string
}
