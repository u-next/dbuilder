# dbuilder

Dgraph query builder aims to create dgraph query filter directive with semantic and type-safe [QueryOperator](#queryoperators).

## Install

```
go get github.com/u-next/dbuilder
```

## Quick Start

```go
func main() {
    clause := dbuilder.NewFilter(dbuilder.ConjunctionAnd).
		Apply("popularity", &dbuilder.FloatQueryOperator{Eq: pointerizer.F64(0.5)}).
		Apply("category", &dbuilder.StringQueryOperator{Eq: pointerizer.S("foo")}).
		Build()

    // @filter(eq(popularity, 0.5) AND eq(category, "foo"))
    fmt.Print(clause)
}
```

## Interface

### QueryOperators

```go
type IntQueryOperator struct {
	Eq  *int // equal to
	Ne  *int // not equal to
	Gt  *int // greater than
	Gte *int // greater than or equal to
	Lt  *int // less than
	Lte *int // less than or equal to
}
```

```go
type FloatQueryOperator struct {
	Eq  *float64 // equal to
	Ne  *float64 // not equal to
	Gt  *float64 // greater than
	Gte *float64 // greater than or equal to
	Lt  *float64 // less than
	Lte *float64 // less than or equal to
}
```

```go
type StringQueryOperator struct {
	Eq         *string  // equal to
	Ne         *string  // not equal to
	In         []string // contains
	Nin        []string // not contains
	Regexp     *string  // regular expression

	// Matches strings that have all specified terms in any order; case insensitive.
	// https://dgraph.io/docs/query-language/functions/#allofterms
	Allofterms *string

	// Matches strings that have any of the specified terms in any order; case insensitive.
	// https://dgraph.io/docs/query-language/functions/#anyofterms
	Anyofterms *string
}
```

```go
type TimeQueryOperator struct {
	Eq  *time.Time // equal to
	Ne  *time.Time // not equal to
	Gt  *time.Time // greater than
	Gte *time.Time // greater than or equal to
	Lt  *time.Time // less than
	Lte *time.Time // less than or equal to
}
```

```go
type BooleanQueryOperator struct {
	Eq *bool // equal to
}
```

### Has

Determines if a node has a particular predicate.

```go
func main() {
    clause := dbuilder.NewFilter(dbuilder.ConjunctionAnd).
		Has("popularity").
		Build()

    // @filter(has(popularity))
    fmt.Print(clause)
}
```
