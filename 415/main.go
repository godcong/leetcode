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

func getSumByte(num1 string, num2 string) ([]byte, []byte) {
	if len(num1) > len(num2) {
		return []byte(num2), []byte(num1)
	}
	return []byte(num1), []byte(num2)
}

func sumString(num []byte, index int, add int) (ret []byte) {
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
	n1, n2 := getSumByte(num1, num2)
	loop := len(n1) - 1
	for i := 1; loop >= 0; loop-- {
		indexByte := strings.IndexByte(numList, n1[loop])
		n2 = sumString(n2, len(n2)-i, indexByte)
		i++
	}
	return string(n2)
}

func main() {
	fmt.Println("test string 1+1", addStrings("1", "1"))
	fmt.Println("test string 4+4", addStrings("4", "4"))
	fmt.Println("test string 9+1", addStrings("9", "1"))
	fmt.Println("test string 99+1", addStrings("99", "1"))
	fmt.Println("test string 999+999", addStrings("999", "999"))
	fmt.Println("test string 123+456", addStrings("123", "456"))
	fmt.Println("test string 9133+0", addStrings("9133", "0"))
}
