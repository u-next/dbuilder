package dbuilder

import (
	"strconv"

	"github.com/u-next/dbuilder/function"
)

type StringQueryOperator struct {
	Eq         *string  `json:"eq,omitempty"`
	Ne         *string  `json:"ne,omitempty"`
	In         []string `json:"in,omitempty"`
	Nin        []string `json:"nin,omitempty"`
	Regexp     *string  `json:"regexp,omitempty"`
	Allofterms *string  `json:"allofterms,omitempty"`
	Anyofterms *string  `json:"anyofterms,omitempty"`
}

func (op *StringQueryOperator) execute(pred string) *Expression {
	var fns []function.Function

	if op == nil {
		return nil
	}

	if op.Eq != nil {
		fns = append(fns, function.NewEq(pred, strconv.Quote(*op.Eq)))
	}

	if op.Ne != nil {
		fns = append(fns, function.NewNe(pred, strconv.Quote(*op.Ne)))
	}

	if len(op.In) > 0 {
		val := make([]interface{}, 0, len(op.In))
		for _, s := range QuoteStrings(op.In) {
			val = append(val, s)
		}

		fns = append(fns, function.NewEq(pred, val...))
	}

	if len(op.Nin) > 0 {
		val := make([]interface{}, 0, len(op.Nin))
		for _, s := range QuoteStrings(op.Nin) {
			val = append(val, s)
		}

		fns = append(fns, function.NewNe(pred, val...))
	}

	if op.Regexp != nil {
		fns = append(fns, function.NewRegexp(pred, *op.Regexp))
	}

	if op.Allofterms != nil {
		fns = append(fns, function.NewAllofterms(pred, strconv.Quote(*op.Allofterms)))
	}

	if op.Anyofterms != nil {
		fns = append(fns, function.NewAnyofterms(pred, strconv.Quote(*op.Anyofterms)))
	}

	return &Expression{
		fns: fns,
	}
}
