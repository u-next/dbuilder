package dbuilder

import "github.com/u-next/dbuilder/function"

type CustomQueryOperator struct {
	Expression *string `json:"expression,omitempty"`
}

func (op *CustomQueryOperator) execute(string) *Expression {
	var fns []function.Function

	if op == nil || op.Expression == nil {
		return nil
	}

	if op.Expression != nil {
		fns = append(fns, function.NewEcho(*op.Expression))
	}

	return &Expression{
		fns: fns,
	}
}
