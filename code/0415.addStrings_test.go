package code

import "testing"

func Test_addStrings(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				num1: "0",
				num2: "0",
			},
			want: "0",
		},
		{
			name: "",
			args: args{
				num1: "9",
				num2: "1",
			},
			want: "10",
		},
		{
			name: "",
			args: args{
				num1: "99",
				num2: "1",
			},
			want: "100",
		},
		{
			name: "",
			args: args{
				num1: "999",
				num2: "999",
			},
			want: "1998",
		},
		{
			name: "",
			args: args{
				num1: "123",
				num2: "456",
			},
			want: "579",
		},
		{
			name: "",
			args: args{
				num1: "9133",
				num2: "0",
			},
			want: "9133",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addStrings(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("addStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
