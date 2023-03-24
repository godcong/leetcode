package _1032

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func TestConstructor(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want StreamChecker
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStreamChecker_Query(t *testing.T) {
	type args struct {
		letter byte
	}
	tests := []struct {
		name string
		this *StreamChecker
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Query(tt.args.letter); got != tt.want {
				t.Errorf("StreamChecker.Query() = %v, want %v", got, tt.want)
			}
		})
	}
}
