package _2469

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_convertTemperature(t *testing.T) {
	type args struct {
		celsius float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertTemperature(tt.args.celsius); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertTemperature() = %v, want %v", got, tt.want)
			}
		})
	}
}
