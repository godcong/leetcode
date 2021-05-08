package _1723

import "testing"

func Test_minimumTimeRequired(t *testing.T) {
	type args struct {
		jobs []int
		k    int
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
				if got := minimumTimeRequired(tt.args.jobs, tt.args.k); got != tt.want {
					t.Errorf("minimumTimeRequired() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
