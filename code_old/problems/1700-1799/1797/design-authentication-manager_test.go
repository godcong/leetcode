package _1797

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func TestConstructor(t *testing.T) {
	type args struct {
		timeToLive int
	}
	tests := []struct {
		name string
		args args
		want AuthenticationManager
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.timeToLive); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthenticationManager_Generate(t *testing.T) {
	type args struct {
		tokenId     string
		currentTime int
	}
	tests := []struct {
		name string
		this *AuthenticationManager
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.Generate(tt.args.tokenId, tt.args.currentTime)
		})
	}
}

func TestAuthenticationManager_Renew(t *testing.T) {
	type args struct {
		tokenId     string
		currentTime int
	}
	tests := []struct {
		name string
		this *AuthenticationManager
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.Renew(tt.args.tokenId, tt.args.currentTime)
		})
	}
}

func TestAuthenticationManager_CountUnexpiredTokens(t *testing.T) {
	type args struct {
		currentTime int
	}
	tests := []struct {
		name string
		this *AuthenticationManager
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.CountUnexpiredTokens(tt.args.currentTime); got != tt.want {
				t.Errorf("AuthenticationManager.CountUnexpiredTokens() = %v, want %v", got, tt.want)
			}
		})
	}
}
