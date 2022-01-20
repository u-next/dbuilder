package function

import "fmt"

type Lt struct {
	pred  string
	value interface{}
}

func NewLt(pred string, value interface{}) Lt {
	return Lt{pred: pred, value: value}
}

func (op Lt) String() string {
	return fmt.Sprintf("lt(%s, %v)", op.pred, op.value)
}
