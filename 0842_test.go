package leetcode

import (
	"reflect"
	"testing"
)

func Test_splitIntoFibonacci(t *testing.T) {
	type args struct {
		s string
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
				s: "123456579",
			},
			want: strToIntArray("[123,456,579]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitIntoFibonacci(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitIntoFibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
