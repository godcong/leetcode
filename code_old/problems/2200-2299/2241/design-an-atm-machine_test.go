package _2241

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want ATM
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

func TestATM_Deposit(t *testing.T) {
	type args struct {
		banknotesCount []int
	}
	tests := []struct {
		name string
		this *ATM
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.Deposit(tt.args.banknotesCount)
		})
	}
}

func TestATM_Withdraw(t *testing.T) {
	type args struct {
		amount int
	}
	tests := []struct {
		name string
		this *ATM
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Withdraw(tt.args.amount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ATM.Withdraw() = %v, want %v", got, tt.want)
			}
		})
	}
}
