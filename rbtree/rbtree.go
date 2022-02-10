package rbtree

import (
	"github.com/lxzan/dao"
	"testing"
)

type Color uint8

const (
	BLACK Color = 0
	RED   Color = 1
)

type Iterator[K dao.Comparable[K], V any] struct {
	Key  K // unique
	Val  V
	next bool
}

func (this *Iterator[K, V]) Break() {
	this.next = false
}

type rbtree_node[K dao.Comparable[K], V any] struct {
	left   *rbtree_node[K, V]
	right  *rbtree_node[K, V]
	parent *rbtree_node[K, V]
	color  Color
	data   *Iterator[K, V]
}

func (this *rbtree_node[K, V]) resert_key() {
	var data Iterator[K, V]
	this.data = &data
}

func (this *rbtree_node[K, V]) set_black() {
	this.color = BLACK
}

func (this *rbtree_node[K, V]) set_red() {
	this.color = RED
}

func (this *rbtree_node[K, V]) is_black() bool {
	return this.color == BLACK
}

func (this *rbtree_node[K, V]) is_red() bool {
	return this.color == RED
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
	length   int
	root     *rbtree_node[K, V]
	sentinel *rbtree_node[K, V]
}

func New[K dao.Comparable[K], V any]() *RBTree[K, V] {
	var node rbtree_node[K, V]
	return &RBTree[K, V]{root: &node, sentinel: &node, length: 0}
}

func (this *RBTree[K, V]) Len() int {
	return this.length
}

func (this *RBTree[K, V]) is_key_empty(d *Iterator[K, V]) bool {
	return d == nil
}

func (this *RBTree[K, V]) begin() *rbtree_node[K, V] {
	return this.root
}

func (this *RBTree[K, V]) next(iter *rbtree_node[K, V], ele *Iterator[K, V]) *rbtree_node[K, V] {
	if ele.Key > iter.data.Key {
		return iter.right
	}
	return iter.left
}

func (this *RBTree[K, V]) end(iter *rbtree_node[K, V]) bool {
	return iter.data == nil
}

func (this *RBTree[K, V]) left_rotate(root **rbtree_node[K, V], sentinel *rbtree_node[K, V], node *rbtree_node[K, V]) {
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

func (this *RBTree[K, V]) right_rotate(root **rbtree_node[K, V], sentinel *rbtree_node[K, V], node *rbtree_node[K, V]) {
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

func (this *RBTree[K, V]) do_insert(temp *rbtree_node[K, V], node *rbtree_node[K, V], sentinel *rbtree_node[K, V]) {
	var p **rbtree_node[K, V]
	for {
		if node.data.Key < temp.data.Key {
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

func (this *RBTree[K, V]) do_delete(node *rbtree_node[K, V]) {
	var red bool
	var root **rbtree_node[K, V]
	var sentinel, subst, temp, w *rbtree_node[K, V]

	/* a binary tree delete */

	root = &this.root
	sentinel = this.sentinel
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
				this.left_rotate(root, sentinel, temp.parent)
				w = temp.parent.right
			}
			if w.left.is_black() && w.right.is_black() {
				w.set_red()
				temp = temp.parent

			} else {
				if w.right.is_black() {
					w.left.set_black()
					w.set_red()
					this.right_rotate(root, sentinel, w)
					w = temp.parent.right
				}

				rbt_copy_color(w, temp.parent)
				temp.parent.set_black()
				w.right.set_black()
				this.left_rotate(root, sentinel, temp.parent)
				temp = *root
			}
		} else {
			w = temp.parent.left
			if w.is_red() {
				w.set_black()
				temp.parent.set_red()
				this.right_rotate(root, sentinel, temp.parent)
				w = temp.parent.left
			}
			if w.left.is_black() && w.right.is_black() {
				w.set_red()
				temp = temp.parent

			} else {
				if w.left.is_black() {
					w.right.set_black()
					w.set_red()
					this.left_rotate(root, sentinel, w)
					w = temp.parent.left
				}
				rbt_copy_color(w, temp.parent)
				temp.parent.set_black()
				w.left.set_black()
				this.right_rotate(root, sentinel, temp.parent)
				temp = *root
			}
		}
	}
	temp.set_black()
}

func (this *RBTree[K, V]) validate(t *testing.T, node *rbtree_node[K, V]) {
	if node == nil {
		return
	}
	if node.left != nil {
		if !this.is_key_empty(node.left.data) && node.data.Key < node.left.data.Key {
			t.Error("left node error!")
		}
		this.validate(t, node.left)
	}

	if node.right != nil {
		if !this.is_key_empty(node.right.data) && node.data.Key > node.right.data.Key {
			t.Error("right node error!")
		}
		this.validate(t, node.right)
	}
}

// insert with unique check
func (this *RBTree[K, V]) Insert(data *Iterator[K, V]) (success bool) {
	for i := this.begin(); !this.end(i); i = this.next(i, data) {
		if data.Key == i.data.Key {
			return false
		}
	}

	this.length++
	var node = &rbtree_node[K, V]{data: data}
	var root = &this.root
	var temp, sentinel *rbtree_node[K, V]

	/* a binary tree insert */

	sentinel = this.sentinel
	if *root == sentinel {
		node.parent = nil
		node.left = sentinel
		node.right = sentinel
		node.set_black()
		*root = node
		return
	}
	this.do_insert(*root, node, sentinel)

	/* re-balance tree */

	for node != *root && node.parent.is_red() {
		if node.parent == node.parent.parent.left {
			temp = node.parent.parent.right
			if temp.is_red() {
				node.parent.set_black()
				temp.set_black()
				node.parent.parent.set_red()
				node = node.parent.parent
			} else {
				if node == node.parent.right {
					node = node.parent
					this.left_rotate(root, sentinel, node)
				}
				node.parent.set_black()
				node.parent.parent.set_red()
				this.right_rotate(root, sentinel, node.parent.parent)
			}
		} else {
			temp = node.parent.parent.left

			if temp.is_red() {
				node.parent.set_black()
				temp.set_black()
				node.parent.parent.set_red()
				node = node.parent.parent
			} else {
				if node == node.parent.left {
					node = node.parent
					this.right_rotate(root, sentinel, node)
				}
				node.parent.set_black()
				node.parent.parent.set_red()
				this.left_rotate(root, sentinel, node.parent.parent)
			}
		}
	}
	(*root).set_black()
	return true
}

func (this *RBTree[K, V]) Delete(key K) (success bool) {
	var data = &Iterator[K, V]{Key: key}
	for i := this.begin(); !this.end(i); i = this.next(i, data) {
		if key == i.data.Key {
			this.length--
			this.do_delete(i)
			return true
		}
	}
	return false
}
