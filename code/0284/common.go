package _0284

type Iterator struct {
	current int
	ele     []interface{}
}

func (this *Iterator) hasNext() bool {
	// Returns true if the iteration has more elements.
	return this.current < len(this.ele)
}

func (this *Iterator) next() int {
	return this.current
}
