package rbtree

import "github.com/lxzan/dao"

type Color uint8

const (
	BLACK Color = 0
	RED   Color = 1
)

type rbtree_node[K dao.Comparable[K], V any] struct {
	left   *rbtree_node[K, V]
	right  *rbtree_node[K, V]
	parent *rbtree_node[K, V]
	color  Color
	key    K // unique, not empty
	data   *V
}

func (c *rbtree_node[K, V]) resert_key() {
	var key K
	c.key = key
}

func (c *rbtree_node[K, V]) set_black() {
	c.color = BLACK
}

func (c *rbtree_node[K, V]) set_red() {
	c.color = RED
}

func (c *rbtree_node[K, V]) is_black() bool {
	return c.color == BLACK
}

func (c *rbtree_node[K, V]) is_red() bool {
	return c.color == RED
}

func rbt_copy_color[K dao.Comparable[K], V any](n1, n2 *rbtree_node[K, V]) {
	n1.color = n2.color
}

func rbtree_min[K dao.Comparable[K], V any](node *rbtree_node[K, V], sentinel *rbtree_node[K, V]) *rbtree_node[K, V] {
	for node.left != sentinel {
		node = node.left
	}
	return node
}

type RBTree[K dao.Comparable[K], V any] struct {
	key      K // empty key
	length   int
	root     *rbtree_node[K, V]
	sentinel *rbtree_node[K, V]
}

func (c *RBTree[K, V]) is_key_empty(key K) bool {
	return key == c.key
}

func (c *RBTree[K, V]) begin() *rbtree_node[K, V] {
	return c.root
}

func (c *RBTree[K, V]) next(iter *rbtree_node[K, V], key K) *rbtree_node[K, V] {
	if key > iter.key {
		return iter.right
	}
	return iter.left
}

func (c *RBTree[K, V]) end(iter *rbtree_node[K, V]) bool {
	return iter.data == nil
}

func (c *RBTree[K, V]) left_rotate(root **rbtree_node[K, V], sentinel *rbtree_node[K, V], node *rbtree_node[K, V]) {
	var temp *rbtree_node[K, V]
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

func (c *RBTree[K, V]) right_rotate(root **rbtree_node[K, V], sentinel *rbtree_node[K, V], node *rbtree_node[K, V]) {
	var temp *rbtree_node[K, V]
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

func (c *RBTree[K, V]) do_insert(temp *rbtree_node[K, V], node *rbtree_node[K, V], sentinel *rbtree_node[K, V]) {
	var p **rbtree_node[K, V]
	for {
		if node.key < temp.key {
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

func (c *RBTree[K, V]) do_delete(node *rbtree_node[K, V]) {
	var red bool
	var root **rbtree_node[K, V]
	var sentinel, subst, temp, w *rbtree_node[K, V]

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
