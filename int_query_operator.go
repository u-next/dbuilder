package dbuilder

import "github.com/u-next/dbuilder/function"

type IntQueryOperator struct {
	Eq  *int `json:"eq,omitempty"`
	Ne  *int `json:"ne,omitempty"`
	Gt  *int `json:"gt,omitempty"`
	Gte *int `json:"gte,omitempty"`
	Lt  *int `json:"lt,omitempty"`
	Lte *int `json:"lte,omitempty"`
}

func (op *IntQueryOperator) execute(pred string) *Expression {
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
