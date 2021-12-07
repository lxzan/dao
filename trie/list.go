package trie

type Pair struct {
	Key string
	Val int
}

type ListElement struct {
	Value Pair
	Next  *ListElement
}

type Queue struct {
	length int
	head   *ListElement
	tail   *ListElement
}

func NewQueue() *Queue {
	return new(Queue)
}

func (c *Queue) Len() int {
	return c.length
}

func (c *Queue) Push(p Pair) {
	var node = &ListElement{Value: p}
	if c.length == 0 {
		c.head = node
		c.tail = node
	} else {
		c.tail.Next = node
		c.tail = node
	}
	c.length++
}
