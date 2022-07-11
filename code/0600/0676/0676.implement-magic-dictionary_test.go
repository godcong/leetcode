package _0676

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want MagicDictionary
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMagicDictionary_BuildDict(t *testing.T) {
	type args struct {
		dictionary []string
	}
	tests := []struct {
		name string
		this *MagicDictionary
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.BuildDict(tt.args.dictionary)
		})
	}
}

func TestMagicDictionary_Search(t *testing.T) {
	type args struct {
		searchWord string
	}
	tests := []struct {
		name string
		this *MagicDictionary
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Search(tt.args.searchWord); got != tt.want {
				t.Errorf("MagicDictionary.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
