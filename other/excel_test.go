package other

import "testing"

func Test_excel(t *testing.T) {
	type args struct {
		number int
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
				number: 1,
			},
			want: "A",
		},
		{
			name: "",
			args: args{
				number: 27,
			},
			want: "AA",
		},
		{
			name: "",
			args: args{
				number: 29,
			},
			want: "AC",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := excel(tt.args.number); got != tt.want {
				t.Errorf("excel() = %v, want %v", got, tt.want)
			}
		})
	}
}
