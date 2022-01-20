package function

// Function allows filtering based on properties of nodes or variables
// Now what the supporting functions are:
//   - Eq
//   - Gt
//   - Gte
//   - Lt
//   - Lte
//   - Regexp
//   - Allofterms
//   - Anyofterms
// https://dgraph.io/docs/query-language/functions/
type Function interface {
	// String converts function AST to string expression in dgraph
	String() string
}
