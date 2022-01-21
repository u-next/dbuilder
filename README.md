# dbuilder

Dgraph query builder

## Install

```
go get github.com/u-next/dbuilder
```

## Quick Start

```go
func main() {
    clause := dbuilder.NewFilter(dbuilder.ConjunctionAnd).
		Apply("popularity", &dbuilder.FloatQueryOperator{
			Eq: pointerizer.F64(0.5),
		}).
		Apply("category", &dbuilder.StringQueryOperator{
			Eq: pointerizer.S("foo"),
		}).
		Build()

    // @filter(eq(popularity, 0.5) AND eq(category, "foo"))
    fmt.Print(clause)
}
```
