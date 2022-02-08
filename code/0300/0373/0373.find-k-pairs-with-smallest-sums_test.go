package _0373

import (
	"reflect"
	"testing"
)

func Test_kSmallestPairs(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
		k     int
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
			if got := kSmallestPairs(tt.args.nums1, tt.args.nums2, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kSmallestPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
