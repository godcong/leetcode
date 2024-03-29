package old

import "testing"

func Test_findTheDifference(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				s: "abcd",
				t: "abcde",
			},
			want: 'e',
		},
		{
			name: "",
			args: args{
				s: "",
				t: "y",
			},
			want: 'y',
		},
		{
			name: "",
			args: args{
				s: "a",
				t: "aa",
			},
			want: 'a',
		},
		{
			name: "",
			args: args{
				s: "ae",
				t: "aea",
			},
			want: 'a',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheDifference(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("findTheDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
