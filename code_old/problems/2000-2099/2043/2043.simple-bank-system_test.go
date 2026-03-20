package _2043

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func TestConstructor(t *testing.T) {
	type args struct {
		balance []int64
	}
	tests := []struct {
		name string
		args args
		want Bank
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.balance); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBank_Transfer(t *testing.T) {
	type args struct {
		account1 int
		account2 int
		money    int64
	}
	tests := []struct {
		name string
		this *Bank
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Transfer(tt.args.account1, tt.args.account2, tt.args.money); got != tt.want {
				t.Errorf("Bank.Transfer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBank_Deposit(t *testing.T) {
	type args struct {
		account int
		money   int64
	}
	tests := []struct {
		name string
		this *Bank
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Deposit(tt.args.account, tt.args.money); got != tt.want {
				t.Errorf("Bank.Deposit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBank_Withdraw(t *testing.T) {
	type args struct {
		account int
		money   int64
	}
	tests := []struct {
		name string
		this *Bank
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Withdraw(tt.args.account, tt.args.money); got != tt.want {
				t.Errorf("Bank.Withdraw() = %v, want %v", got, tt.want)
			}
		})
	}
}
