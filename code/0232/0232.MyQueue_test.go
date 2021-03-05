package _0232

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		nums [][]int
		step []string
	}
	tests := []struct {
		name string
		args args
		want MyQueue
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				nums: strto,
				step: nil,
			},
			want: MyQueue{
				inStack:  nil,
				outStack: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}
