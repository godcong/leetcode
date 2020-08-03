package main

import (
	"fmt"
	"strings"
)

//var numList = map[rune]uint8{
//	'0': 0,
//	'1': 1,
//	'2': 2,
//	'3': 3,
//	'4': 4,
//	'5': 5,
//	'6': 6,
//	'7': 7,
//	'8': 8,
//	'9': 9,
//}

var numList = "01234567890123456789"

func getMin(num1 string, num2 string) (string, string) {
	if len(num1) > len(num2) {
		return num2, num1
	}
	return num1, num2
}

func sumString(num []byte, index int, add int) []byte {
	if index < 0 {
		return append([]byte{'1'}, num...)
	}
	indexByte := strings.IndexByte(numList, num[index])
	num[index] = numList[indexByte+add]
	if (indexByte + add) > 9 {
		num = sumString(num, index-1, 1)
	}
	return num
}

func addStrings(num1 string, num2 string) string {
	num1, num2 = getMin(num1, num2)
	loop := len(num1) - 1
	for i := len(num2) - 1; loop >= 0; loop-- {
		indexByte := strings.IndexByte(numList, num1[loop])
		num2 = string(sumString([]byte(num2), i, indexByte))
		i--
	}
	return num2
}

func main() {
	fmt.Println("test string 1+1", addStrings("1", "1"))
	fmt.Println("test string 4+4", addStrings("4", "4"))
	fmt.Println("test string 9+1", addStrings("9", "1"))
	fmt.Println("test string 99+1", addStrings("99", "1"))
}
