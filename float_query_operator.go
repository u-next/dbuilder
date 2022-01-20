package dbuilder

import "github.com/u-next/dbuilder/function"

type FloatQueryOperator struct {
	Eq  *float64 `json:"eq,omitempty"`
	Ne  *float64 `json:"ne,omitempty"`
	Gt  *float64 `json:"gt,omitempty"`
	Gte *float64 `json:"gte,omitempty"`
	Lt  *float64 `json:"lt,omitempty"`
	Lte *float64 `json:"lte,omitempty"`
}

func (op *FloatQueryOperator) execute(pred string) *Expression {
	var fns []function.Function

	if op == nil {
		return nil
	}

	if op.Eq != nil {
		fns = append(fns, function.NewEq(pred, *op.Eq))
	}

	if op.Ne != nil {
		fns = append(fns, function.NewNe(pred, *op.Ne))
	}

	if op.Gt != nil {
		fns = append(fns, function.NewGt(pred, *op.Gt))
	}

	if op.Gte != nil {
		fns = append(fns, function.NewGte(pred, *op.Gte))
	}

	if op.Lt != nil {
		fns = append(fns, function.NewLt(pred, *op.Lt))
	}

	if op.Lte != nil {
		fns = append(fns, function.NewLte(pred, *op.Lte))
	}

	return &Expression{
		fns: fns,
	}
}
