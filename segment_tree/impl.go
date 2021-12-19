package segment_tree

type Node struct {
	Sum int
}

func (n Node) NewFrom(x int) Node {
	return Node{Sum: x}
}

func (n Node) Merge(a, b Node) Node {
	return Node{Sum: a.Sum + b.Sum}
}
