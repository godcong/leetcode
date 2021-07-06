package _1418

import (
	"reflect"
	"testing"
)

func Test_displayTable(t *testing.T) {
	type args struct {
		orders [][]string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := displayTable(tt.args.orders); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("displayTable() = %v, want %v", got, tt.want)
			}
		})
	}
}
