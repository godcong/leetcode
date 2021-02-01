package code

import (
	"reflect"
	"testing"
)

func Test_fairCandySwap(t *testing.T) {
	type args struct {
		A []int
		B []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				A: strToIntArray("[1,1]"),
				B: strToIntArray("[2,2]"),
			},
			want: strToIntArray("[1,2]"),
		},
		{
			name: "",
			args: args{
				A: strToIntArray("[1,2]"),
				B: strToIntArray("[2,3]"),
			},
			want: strToIntArray("[1,2]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fairCandySwap(tt.args.A, tt.args.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fairCandySwap() = %v, want %v", got, tt.want)
			}
		})
	}
}
