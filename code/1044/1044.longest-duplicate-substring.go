package _1044

import (
	"index/suffixarray"
	"reflect"
	"unsafe"
)

func longestDupSubstring(s string) string {
	n := len(s)
	sa := *(*[]int32)(unsafe.Pointer(reflect.ValueOf(suffixarray.New([]byte(s))).Elem().FieldByName("sa").Field(0).UnsafeAddr()))
	rank := make([]int, n)
	for i := range rank {
		rank[sa[i]] = i
	}
	height := make([]int, n)
	h := 0
	for i, rk := range rank {
		if h > 0 {
			h--
		}
		if rk > 0 {
			for j := int(sa[rk-1]); i+h < n && j+h < n && s[i+h] == s[j+h]; h++ {
			}
		}
		height[rk] = h
	}

	idx, maxH := 0, 0
	for i, h := range height {
		if h > maxH {
			idx, maxH = i, h
		}
	}
	return s[sa[idx] : int(sa[idx])+maxH]
}
