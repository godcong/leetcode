package _1034

import (
	"reflect"
	"testing"
)

func Test_colorBorder(t *testing.T) {
	type args struct {
		grid  [][]int
		row   int
		col   int
		color int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := colorBorder(tt.args.grid, tt.args.row, tt.args.col, tt.args.color); !reflect.DeepEqual(got,
				tt.want) {
				t.Errorf("colorBorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
