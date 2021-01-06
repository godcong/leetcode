package code

import (
	"reflect"
	"testing"
)

func Test_calcEquation(t *testing.T) {
	type args struct {
		equations [][]string
		values    []float64
		queries   [][]string
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				equations: strToStringArrayArray("[[\"a\",\"b\"],[\"b\",\"c\"]]"),
				values: []float64{
					2.0, 3.0,
				},
				queries: strToStringArrayArray("[[\"a\",\"c\"],[\"b\",\"a\"],[\"a\",\"e\"],[\"a\",\"a\"],[\"x\",\"x\"]]"),
			},
			want: []float64{
				6.00000, 0.50000, -1.00000, 1.00000, -1.00000,
			},
		},
		{
			name: "",
			args: args{
				equations: strToStringArrayArray("[[\"a\",\"b\"],[\"b\",\"c\"],[\"bc\",\"cd\"]]"),
				values: []float64{
					1.5, 2.5, 5.0,
				},
				queries: strToStringArrayArray("[[\"a\",\"c\"],[\"c\",\"b\"],[\"bc\",\"cd\"],[\"cd\",\"bc\"]]"),
			},
			want: []float64{
				3.75000, 0.40000, 5.00000, 0.20000,
			},
		},
		{
			name: "",
			args: args{
				equations: strToStringArrayArray("[[\"a\",\"b\"]]"),
				values:    nil,
				queries:   nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcEquation(tt.args.equations, tt.args.values, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calcEquation() = %v, want %v", got, tt.want)
			}
		})
	}
}
