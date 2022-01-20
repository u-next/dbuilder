package function

import "fmt"

type Lte struct {
	pred  string
	value interface{}
}

func NewLte(pred string, value interface{}) Lte {
	return Lte{pred: pred, value: value}
}

func (op Lte) String() string {
	return fmt.Sprintf("le(%s, %v)", op.pred, op.value)
}
