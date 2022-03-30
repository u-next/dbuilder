package dbuilder

import (
	"github.com/u-next/dbuilder/pointerizer"
	"reflect"
	"testing"
)

func TestCustomQueryOperator_execute(t *testing.T) {
	tests := []struct {
		name       string
		Expression *string
		pred       string
		want       string
	}{
		{
			name:       "Returns the expected expression string and ignores the pred",
			Expression: pointerizer.S("eq(predicate,\"value\")"),
			pred:       "ignore_this",
			want:       "eq(predicate,\"value\")",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := &CustomQueryOperator{
				Expression: tt.Expression,
			}
			if got := op.execute(tt.pred).Build(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
