package code

import (
	"testing"
)

func Test_validMountainArray(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				A: strToIntArray("[2,1]"),
			},
			want: false,
		},
		{
			name: "",
			args: args{
				A: strToIntArray("[3,5,5]"),
			},
			want: false,
		},
		{
			name: "",
			args: args{
				A: strToIntArray("[0,3,2,1]"),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validMountainArray(tt.args.A); got != tt.want {
				t.Errorf("validMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
