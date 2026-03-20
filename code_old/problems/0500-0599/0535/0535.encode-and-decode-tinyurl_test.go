package _0535

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want Codec
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

func TestCodec_encode(t *testing.T) {
	type args struct {
		longUrl string
	}
	tests := []struct {
		name string
		this *Codec
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.encode(tt.args.longUrl); got != tt.want {
				t.Errorf("Codec.encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCodec_decode(t *testing.T) {
	type args struct {
		shortUrl string
	}
	tests := []struct {
		name string
		this *Codec
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.decode(tt.args.shortUrl); got != tt.want {
				t.Errorf("Codec.decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
