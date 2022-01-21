package dbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/u-next/dbuilder/function"
	"github.com/u-next/dbuilder/pointerizer"
)

func TestStringQueryOperator_execute(t *testing.T) {
	type fields struct {
		Eq         *string
		Ne         *string
		In         []string
		Nin        []string
		Regexp     *string
		Allofterms *string
		Anyofterms *string
	}
	type args struct {
		pred string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Expression
	}{
		{
			name: "should return expressions correctly",
			fields: fields{
				Eq:         pointerizer.S("Eq"),
				Ne:         pointerizer.S("Ne"),
				In:         []string{"In"},
				Nin:        []string{"Nin"},
				Regexp:     pointerizer.S("/regular-expression/"),
				Allofterms: pointerizer.S("Allofterms"),
				Anyofterms: pointerizer.S("Anyofterms"),
			},
			args: args{pred: "foo"},
			want: &Expression{
				fns: []function.Function{
					function.NewEq("foo", "\"Eq\""),
					function.NewNe("foo", "\"Ne\""),
					function.NewEq("foo", "\"In\""),
					function.NewNe("foo", "\"Nin\""),
					function.NewRegexp("foo", "/regular-expression/"),
					function.NewAllofterms("foo", "\"Allofterms\""),
					function.NewAnyofterms("foo", "\"Anyofterms\""),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := &StringQueryOperator{
				Eq:         tt.fields.Eq,
				Ne:         tt.fields.Ne,
				In:         tt.fields.In,
				Nin:        tt.fields.Nin,
				Regexp:     tt.fields.Regexp,
				Allofterms: tt.fields.Allofterms,
				Anyofterms: tt.fields.Anyofterms,
			}
			got := op.execute(tt.args.pred)

			assert.Equal(t, tt.want, got)
		})
	}
}
