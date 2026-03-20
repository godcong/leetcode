package _0472

import (
	"reflect"
	"testing"
)

func Test_findAllConcatenatedWordsInADict(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAllConcatenatedWordsInADict(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAllConcatenatedWordsInADict() = %v, want %v", got, tt.want)
			}
		})
	}
}
