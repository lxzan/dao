package segment_tree

import "testing"

func TestNew(t *testing.T) {
	var arr = []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10}
	var tree = New[int, Node](arr)
	println(&tree)
}
