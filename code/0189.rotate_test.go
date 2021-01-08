package code

import "testing"

func Test_rotate0189(t *testing.T) {
	type args struct {
		nums []int
		k    int
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
				nums: strToIntArray("[1,2,3,4,5,6,7]"),
				k:    3,
			},
			want: strToIntArray("[5,6,7,1,2,3,4]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
