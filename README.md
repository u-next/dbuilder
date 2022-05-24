# dbuilder

Dgraph query builder aims to create dgraph query filter directive with semantic and type-safe [QueryOperator](#queryoperators).

## Install

```
go get github.com/u-next/dbuilder
```

## Quick Start

**Basic**

```go
func main() {
	clause := dbuilder.NewFilter(dbuilder.ConjunctionAnd).
		Apply("category", &dbuilder.StringQueryOperator{Eq: pointerizer.S("foo")}).
		Build()

	// @filter(eq(category, "foo"))
	fmt.Print(clause)
}
```

**CustomQueryOperator**

```go
func main() {
    partialFilter := &dbuilder.NewFilter(dbuilder.ConjunctionAnd).
		Apply("other", &dbuilder.StringQueryOperator{Eq: pointerizer.S("foo")}).
 		ToCustomQueryOperator()

    clause := dbuilder.NewFilter(dbuilder.ConjunctionAnd).
		Apply("popularity", &dbuilder.FloatQueryOperator{Eq: pointerizer.F64(0.5)}).
		Apply("category", &dbuilder.StringQueryOperator{Eq: pointerizer.S("foo")}).
		Apply("ignored", &dbuilder.CustomQueryOperator{Expression: pointerizer.S(`lt(media.original_release_date, "1977-01-01") AND gt(media.original_price, 10)`)}).
		Apply("ignored", partialFilter).
		Build()

    // @filter(eq(popularity, 0.5) AND eq(category, "foo") AND lt(media.original_release_date, "1977-01-01") AND gt(media.original_price, 10) AND eq(other, "foo") )
    fmt.Print(clause)
}
```

**Raw Filter String**

```go
func main() {
	clause := dbuilder.NewFilter(dbuilder.ConjunctionAnd).
		Apply("category", &dbuilder.StringQueryOperator{Eq: pointerizer.S("foo")}).
		Build()

	// eq(category, "foo")
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

```go
type CustomQueryOperator struct {
	// any valid filter expression in string format
	// it should contain also the full predicate name.
	// The predicate passed in the "Apply" method call is ignored.
	Expression *string
}
```

### Has

Determines if a node has a particular predicate.

```go
func main() {
    clause := dbuilder.NewFilter(dbuilder.ConjunctionAnd).
		Has("popularity", false).
		Build()

    // @filter(has(popularity))
    fmt.Print(clause)
}
```

### Type

Determine if a node belongs to particular type.

```go
func main() {
    clause := dbuilder.NewFilter(dbuilder.ConjunctionAnd).
		Type("Foo", false).
		Build()

    // @filter(type(Foo))
    fmt.Print(clause)
}
```
