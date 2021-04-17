package _0220

import "testing"

func Test_containsNearbyAlmostDuplicate(t *testing.T) {
	type args struct {
		nums []int
		k    int
		t    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := containsNearbyAlmostDuplicate(tt.args.nums, tt.args.k, tt.args.t); got != tt.want {
					t.Errorf("containsNearbyAlmostDuplicate() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
