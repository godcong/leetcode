package _0027

import "testing"

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := removeElement(tt.args.nums, tt.args.val); got != tt.want {
					t.Errorf("removeElement() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
