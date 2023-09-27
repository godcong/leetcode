package _1333

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_filterRestaurants(t *testing.T) {
	type args struct {
		restaurants   [][]int
		veganFriendly int
		maxPrice      int
		maxDistance   int
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
			if got := filterRestaurants(tt.args.restaurants, tt.args.veganFriendly, tt.args.maxPrice, tt.args.maxDistance); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterRestaurants() = %v, want %v", got, tt.want)
			}
		})
	}
}
