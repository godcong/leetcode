package old

import "testing"

func Test_reorganizeString(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				S: "aab",
			},
			want: "aba",
		},
		{
			name: "",
			args: args{
				S: "aaab",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reorganizeString(tt.args.S); got != tt.want {
				t.Errorf("reorganizeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
