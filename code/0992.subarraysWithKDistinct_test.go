package code

import "testing"

func Test_subarraysWithKDistinct(t *testing.T) {
	type args struct {
		A []int
		K int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				A: strToIntArray("[1,2,1,2,3]"),
				K: 2,
			},
			want: 7,
		},
		{
			name: "",
			args: args{
				A: strToIntArray("[1,2,1,3,4]"),
				K: 3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraysWithKDistinct(tt.args.A, tt.args.K); got != tt.want {
				t.Errorf("subarraysWithKDistinct() = %v, want %v", got, tt.want)
			}
		})
	}
}
