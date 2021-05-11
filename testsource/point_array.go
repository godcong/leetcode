package main

type Val struct {
	Val string
}

func pointArray() []*Val {
	valList := []*Val{
		{
			Val: "1",
		},
		{
			Val: "2",
		},
		{
			Val: "3",
		},
		{
			Val: "4",
		},
	}

	var vals []*Val
	for _, val := range valList {
		vals = append(vals, val)
	}

	return vals

}
