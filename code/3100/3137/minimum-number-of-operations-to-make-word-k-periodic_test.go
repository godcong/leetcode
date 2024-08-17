package _3137

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minimumOperationsToMakeKPeriodic(t *testing.T) {
	type args struct {
		word string
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
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumOperationsToMakeKPeriodic(tt.args.word, tt.args.k); got != tt.want {
				t.Errorf("minimumOperationsToMakeKPeriodic() = %v, want %v", got, tt.want)
			}
		})
	}
}
