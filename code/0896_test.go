package code

import "testing"

func Test_isMonotonic(t *testing.T) {
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
				A: strToIntArray("[1,2,2,3]"),
			},
			want: true,
		},
		{
			name: "",
			args: args{
				A: strToIntArray("[6,5,4,4]"),
			},
			want: true,
		},
		{
			name: "",
			args: args{
				A: strToIntArray("[1,3,2]"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMonotonic(tt.args.A); got != tt.want {
				t.Errorf("isMonotonic() = %v, want %v", got, tt.want)
			}
		})
	}
}
