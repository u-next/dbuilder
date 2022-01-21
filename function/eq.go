package function

import (
	"fmt"
	"strings"
)

// Syntax Examples:
//   - eq(predicate, value)
//   - eq(val(varName), value)
//   - eq(predicate, val(varName))
//   - eq(count(predicate), value)
//   - eq(predicate, [val1, val2, ..., valN])
//   - eq(predicate, [$var1, "value", ..., $varN])
// https://dgraph.io/docs/query-language/functions/#equal-to
type Eq struct {
	pred   string
	values []interface{}
	not    bool
}

func NewEq(pred string, values ...interface{}) Eq {
	return Eq{
		pred:   pred,
		values: values,
	}
}

func NewNe(pred string, values ...interface{}) Eq {
	return Eq{
		pred:   pred,
		values: values,
		not:    true,
	}
}

func (op Eq) String() string {
	val := make([]string, 0, len(op.values))
	for _, v := range op.values {
		val = append(val, fmt.Sprintf("%v", v))
	}

	if op.not {
		return fmt.Sprintf("NOT eq(%s, %s)", op.pred, strings.Join(val, ", "))
	}

	return fmt.Sprintf("eq(%s, %s)", op.pred, strings.Join(val, ", "))
}
