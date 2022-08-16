package _1656

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func TestConstructor(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want OrderedStream
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderedStream_Insert(t *testing.T) {
	type args struct {
		idKey int
		value string
	}
	tests := []struct {
		name string
		this *OrderedStream
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Insert(tt.args.idKey, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderedStream.Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
