package _0786

import (
	"reflect"
	"testing"
)

func Test_kthSmallestPrimeFraction(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallestPrimeFraction(tt.args.arr, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kthSmallestPrimeFraction() = %v, want %v", got, tt.want)
			}
		})
	}
}
