package dbuilder

import "github.com/u-next/dbuilder/function"

type BooleanQueryOperator struct {
	Eq *bool `json:"eq,omitempty"`
}

func (op *BooleanQueryOperator) execute(pred string) *Expression {
	var fns []function.Function

	if op == nil {
		return nil
	}

	if op.Eq != nil {
		if *op.Eq {
			fns = append(fns, function.NewEq(pred, "true"))
		} else {
			fns = append(fns, function.NewEq(pred, "false"))
		}
	}

	return &Expression{
		fns: fns,
	}
}
