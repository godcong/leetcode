package code

import (
	"reflect"
	"testing"
)

func Test_accountsMerge(t *testing.T) {
	type args struct {
		accounts [][]string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				accounts: strToStringArrayArray("[[\"John\", \"johnsmith@mail.com\", \"john00@mail.com\"], [\"John\", \"johnnybravo@mail.com\"], [\"John\", \"johnsmith@mail.com\", \"john_newyork@mail.com\"], [\"Mary\", \"mary@mail.com\"]]"),
			},
			want: strToStringArrayArray("[[\"John\", 'john00@mail.com', 'john_newyork@mail.com', 'johnsmith@mail.com'],  [\"John\", \"johnnybravo@mail.com\"], [\"Mary\", \"mary@mail.com\"]]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := accountsMerge(tt.args.accounts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("accountsMerge() = %v, want %v", got, tt.want)
			}
		})
	}
}
