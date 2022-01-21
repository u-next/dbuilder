package dbuilder

import (
	"strconv"
	"time"

	"github.com/u-next/dbuilder/function"
)

type TimeQueryOperator struct {
	Eq  *time.Time `json:"eq,omitempty"`
	Ne  *time.Time `json:"ne,omitempty"`
	Gt  *time.Time `json:"gt,omitempty"`
	Gte *time.Time `json:"gte,omitempty"`
	Lt  *time.Time `json:"lt,omitempty"`
	Lte *time.Time `json:"lte,omitempty"`
}

func (op *TimeQueryOperator) execute(pred string) *Expression {
	var fns []function.Function

	if op == nil {
		return nil
	}

	if op.Eq != nil {
		fns = append(fns, function.NewEq(pred, strconv.Quote(op.Eq.Format(time.RFC3339))))
	}

	if op.Ne != nil {
		fns = append(fns, function.NewNe(pred, strconv.Quote(op.Ne.Format(time.RFC3339))))
	}

	if op.Gt != nil {
		fns = append(fns, function.NewGt(pred, strconv.Quote(op.Gt.Format(time.RFC3339))))
	}

	if op.Gte != nil {
		fns = append(fns, function.NewGte(pred, strconv.Quote(op.Gte.Format(time.RFC3339))))
	}

	if op.Lt != nil {
		fns = append(fns, function.NewLt(pred, strconv.Quote(op.Lt.Format(time.RFC3339))))
	}

	if op.Lte != nil {
		fns = append(fns, function.NewLte(pred, strconv.Quote(op.Lte.Format(time.RFC3339))))
	}

	return &Expression{
		fns: fns,
	}
}
