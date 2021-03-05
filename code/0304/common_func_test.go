package _0304

import (
	"reflect"
	"testing"
)

func Test_strToIntArray(t *testing.T) {
	type args struct {
		nums string
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
				nums: "[1,2,3]",
			},
			want: []int{
				1, 2, 3,
			},
		},
		{
			name: "",
			args: args{
				nums: "[]",
			},
			want: nil,
		},
		{
			name: "",
			args: args{
				nums: "",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strToIntArray(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("strToIntArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_strToIntArrArray(t *testing.T) {
	type args struct {
		nums string
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				nums: "[[0,1,0,0],[1,1,1,0],[0,1,0,0],[1,1,0,0]]",
			},
			want: [][]int{
				{0, 1, 0, 0},
				{1, 1, 1, 0},
				{0, 1, 0, 0},
				{1, 1, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strToIntArrArray(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("strToIntArrArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
