package main

import (
	"fmt"
)

type SubS struct {
	b string
}

type S struct {
	a    []int
	SubS *SubS
}

func main() {
	//markdownHTML()
	fmt.Println("system", GetPhysicalID())
}
