package main

import (
	"fmt"
)

var numList = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	panic("todo")
}

func main() {
	fmt.Println("III", romanToInt("III"))
	fmt.Println("IV", romanToInt("IV"))
	fmt.Println("IX", romanToInt("IX"))
	fmt.Println("LVIII", romanToInt("LVIII"))
	fmt.Println("MCMXCIV", romanToInt("MCMXCIV"))
}
