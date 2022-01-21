package function

import "fmt"

type Gte struct {
	pred  string
	value interface{}
}

func NewGte(pred string, value interface{}) Gte {
	return Gte{pred: pred, value: value}
}

func (op Gte) String() string {
	return fmt.Sprintf("ge(%s, %v)", op.pred, op.value)
}
