package old

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				strs: []string{
					"flower", "flow", "flight",
				},
			},
			want: "fl",
		},
		{
			name: "",
			args: args{
				strs: []string{
					"dog", "racecar", "car",
				},
			},
			want: "",
		},
		{
			name: "",
			args: args{
				strs: []string{
					"flower", "flow", "flowight",
				},
			},
			want: "flow",
		},
		{
			name: "",
			args: args{
				strs: []string{
					"flower",
				},
			},
			want: "flower",
		},
		{
			name: "",
			args: args{
				strs: []string{},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
