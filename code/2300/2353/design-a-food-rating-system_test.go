package _2353

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func TestConstructor(t *testing.T) {
	type args struct {
		foods    []string
		cuisines []string
		ratings  []int
	}
	tests := []struct {
		name string
		args args
		want FoodRatings
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.foods, tt.args.cuisines, tt.args.ratings); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFoodRatings_ChangeRating(t *testing.T) {
	type args struct {
		food      string
		newRating int
	}
	tests := []struct {
		name string
		this *FoodRatings
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.ChangeRating(tt.args.food, tt.args.newRating)
		})
	}
}

func TestFoodRatings_HighestRated(t *testing.T) {
	type args struct {
		cuisine string
	}
	tests := []struct {
		name string
		this *FoodRatings
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.HighestRated(tt.args.cuisine); got != tt.want {
				t.Errorf("FoodRatings.HighestRated() = %v, want %v", got, tt.want)
			}
		})
	}
}
