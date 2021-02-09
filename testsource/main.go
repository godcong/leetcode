package main

import "fmt"

type SubS struct {
	b string
}

type S struct {
	a    []int
	SubS *SubS
}

func main() {
	s := S{
		a: []int{10, 9, 8, 7},
		SubS: &SubS{
			b: "abc",
		},
	}

	fmt.Println(s)

	s1 := s
	s.a[0] = 7
	s1.SubS.b = "def"
	fmt.Printf("%+v,%+v", s1, s)

}
