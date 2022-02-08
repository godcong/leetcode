package old

import (
	"reflect"
	"testing"
)

func Test_reverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				s: []byte("hello"),
			},
			want: []byte{
				'o', 'l', 'l', 'e', 'h',
			},
		},
		{
			name: "",
			args: args{
				s: []byte{
					'H', 'a', 'n', 'n', 'a', 'h',
				},
			},
			want: []byte{
				'h', 'a', 'n', 'n', 'a', 'H',
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.args.s)
			if !reflect.DeepEqual(tt.args.s, tt.want) {
				t.Errorf("findItinerary() = %v, want %v", tt.args.s, tt.want)
			}
		})
	}
}
