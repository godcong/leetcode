package _3280

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_convertDateToBinary(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertDateToBinary(tt.args.date); got != tt.want {
				t.Errorf("convertDateToBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
