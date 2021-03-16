package common

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
			if got := StrToIntArray(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StrToIntArray() = %v, want %v", got, tt.want)
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
			if got := StrToIntArrArray(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StrToIntArrArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
