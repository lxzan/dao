package rbtree

import "github.com/lxzan/dao"

type Color uint8

const (
	BLACK Color = 0
	RED   Color = 1
)

type rbtree_node[T any] struct {
	left   *rbtree_node[T]
	right  *rbtree_node[T]
	parent *rbtree_node[T]
	color  Color
	data   *T // unique, not empty
}

func (c *rbtree_node[T]) resert_key() {
	var data T
	c.data = &data
}

func (c *rbtree_node[T]) set_black() {
	c.color = BLACK
}

func (c *rbtree_node[T]) set_red() {
	c.color = RED
}

func (c *rbtree_node[T]) is_black() bool {
	return c.color == BLACK
}

func (c *rbtree_node[T]) is_red() bool {
	return c.color == RED
}

func rbt_copy_color[T any](n1, n2 *rbtree_node[T]) {
	n1.color = n2.color
}

func rbtree_min[T any](node *rbtree_node[T], sentinel *rbtree_node[T]) *rbtree_node[T] {
	for node.left != sentinel {
		node = node.left
	}
	return node
}

type RBTree[T any] struct {
	length   int
	cmp      func(a, b *T) dao.Ordering
	root     *rbtree_node[T]
	sentinel *rbtree_node[T]
}

func New[T any](cmp func(a, b *T) dao.Ordering) *RBTree[T] {
	var node rbtree_node[T]
	return &RBTree[T]{root: &node, sentinel: &node, length: 0, cmp: cmp}
}

func (c *RBTree[T]) is_key_empty(d *T) bool {
	return d == nil
}

func (c *RBTree[T]) begin() *rbtree_node[T] {
	return c.root
}

func (c *RBTree[T]) next(iter *rbtree_node[T], ele *T) *rbtree_node[T] {
	if c.cmp(ele, iter.data) == dao.Greater {
		return iter.right
	}
	return iter.left
}

func (c *RBTree[T]) end(iter *rbtree_node[T]) bool {
	return iter.data == nil
}

func (c *RBTree[T]) left_rotate(root **rbtree_node[T], sentinel *rbtree_node[T], node *rbtree_node[T]) {
	var temp *rbtree_node[T]
	temp = node.right
	node.right = temp.left
	if temp.left != sentinel {
		temp.left.parent = node
	}
	temp.parent = node.parent
	if node == *root {
		*root = temp
	} else if node == node.parent.left {
		node.parent.left = temp
	} else {
		node.parent.right = temp
	}
	temp.left = node
	node.parent = temp
}

func (c *RBTree[T]) right_rotate(root **rbtree_node[T], sentinel *rbtree_node[T], node *rbtree_node[T]) {
	var temp *rbtree_node[T]
	temp = node.left
	node.left = temp.right
	if temp.right != sentinel {
		temp.right.parent = node
	}
	temp.parent = node.parent
	if node == *root {
		*root = temp
	} else if node == node.parent.right {
		node.parent.right = temp
	} else {
		node.parent.left = temp
	}
	temp.right = node
	node.parent = temp
}

func (c *RBTree[T]) do_insert(temp *rbtree_node[T], node *rbtree_node[T], sentinel *rbtree_node[T]) {
	var p **rbtree_node[T]
	for {
		if c.cmp(node.data, temp.data) == dao.Less {
			p = &temp.left
		} else {
			p = &temp.right
		}
		if *p == sentinel {
			break
		}
		temp = *p
	}

	*p = node
	node.parent = temp
	node.left = sentinel
	node.right = sentinel
	node.set_red()
}

func (c *RBTree[T]) do_delete(node *rbtree_node[T]) {
	var red bool
	var root **rbtree_node[T]
	var sentinel, subst, temp, w *rbtree_node[T]

	/* a binary tree delete */

	root = &c.root
	sentinel = c.sentinel
	if node.left == sentinel {
		temp = node.right
		subst = node
	} else if node.right == sentinel {
		temp = node.left
		subst = node
	} else {
		subst = rbtree_min(node.right, sentinel)
		temp = subst.right
	}
	if subst == *root {
		*root = temp
		(temp).set_black()
		/* DEBUG stuff */
		node.left = nil
		node.right = nil
		node.parent = nil
		node.resert_key()
		return
	}

	red = subst.is_red()
	if subst == subst.parent.left {
		subst.parent.left = temp
	} else {
		subst.parent.right = temp
	}
	if subst == node {
		temp.parent = subst.parent
	} else {

		if subst.parent == node {
			temp.parent = subst
		} else {
			temp.parent = subst.parent
		}
		subst.left = node.left
		subst.right = node.right
		subst.parent = node.parent
		rbt_copy_color(subst, node)
		if node == *root {
			*root = subst
		} else {
			if node == node.parent.left {
				node.parent.left = subst
			} else {
				node.parent.right = subst
			}
		}
		if subst.left != sentinel {
			subst.left.parent = subst
		}
		if subst.right != sentinel {
			subst.right.parent = subst
		}
	}

	/* DEBUG stuff */
	node.left = nil
	node.right = nil
	node.parent = nil
	node.resert_key()

	if red {
		return
	}

	/* a delete fixup */

	for temp != *root && temp.is_black() {
		if temp == temp.parent.left {
			w = temp.parent.right

			if w.is_red() {
				w.set_black()
				temp.parent.set_red()
				c.left_rotate(root, sentinel, temp.parent)
				w = temp.parent.right
			}
			if w.left.is_black() && w.right.is_black() {
				w.set_red()
				temp = temp.parent

			} else {
				if w.right.is_black() {
					w.left.set_black()
					w.set_red()
					c.right_rotate(root, sentinel, w)
					w = temp.parent.right
				}

				rbt_copy_color(w, temp.parent)
				temp.parent.set_black()
				w.right.set_black()
				c.left_rotate(root, sentinel, temp.parent)
				temp = *root
			}
		} else {
			w = temp.parent.left
			if w.is_red() {
				w.set_black()
				temp.parent.set_red()
				c.right_rotate(root, sentinel, temp.parent)
				w = temp.parent.left
			}
			if w.left.is_black() && w.right.is_black() {
				w.set_red()
				temp = temp.parent

			} else {
				if w.left.is_black() {
					w.right.set_black()
					w.set_red()
					c.left_rotate(root, sentinel, w)
					w = temp.parent.left
				}
				rbt_copy_color(w, temp.parent)
				temp.parent.set_black()
				w.left.set_black()
				c.right_rotate(root, sentinel, temp.parent)
				temp = *root
			}
		}
	}
	temp.set_black()
}
