package genericds

import (
	"errors"
)

//
//         (value)
//          root
//         /    \
//      Left    Right
//
type BstNode struct {
	value       interface{}
	right, left *BstNode
}

type Comparator func(interface{}, interface{}) (int, error)

type BstTree struct {
	root       *BstNode
	comparator Comparator
}

type BinarySearchTree interface {
	Add(value interface{}) (BinarySearchTree, error)
	Find(value interface{}) (bool, error)

	PostorderTraverse(visitor func(interface{}))
	InorderTraverse(visitor func(interface{}))
	Depth() int
	Balance()
}

func NewBinarySearchTree(comparatorFunction Comparator) BinarySearchTree {
	return &BstTree{comparator: comparatorFunction, root: nil}
}

func (tree *BstTree) seekParent(value interface{}) (*BstNode, error) {
	if tree.root == nil {
		return nil, nil
	}
	if tree.comparator == nil {
		return nil, errors.New(" empty comparator ")
	}

	current := tree.root
	parent := current

	for current != nil {
		parent = current
		compareResponse, err := tree.comparator(value, current.value)
		if err != nil {
			return nil, err
		}
		if compareResponse < 0 {
			current = current.left
		} else {
			current = current.right
		}
	}
	return parent, nil
}

func (tree *BstTree) Add(value interface{}) (BinarySearchTree, error) {
	parent, err := tree.seekParent(value)
	if err != nil {
		return nil, err
	}
	node := &BstNode{value: value}
	if parent == nil {
		tree.root = node
		return tree, nil
	}

	compare, _ := tree.comparator(value, parent.value)
	if compare < 0 {
		parent.left = node
	} else {
		parent.right = node
	}
	return tree, nil
}

func (tree *BstTree) Find(value interface{}) (bool, error) {
	if tree.root == nil || tree.comparator == nil {
		return false, errors.New(" Invalid type ")
	}

	current := tree.root

	for current != nil {

		compareResponse, err := tree.comparator(value, current.value)
		if err != nil {
			return false, err
		}
		if compareResponse == 0 {
			return true, nil
		}
		if compareResponse < 0 {
			current = current.left
		} else {
			current = current.right
		}
	}
	return false, nil

}

func buildBalanceBST(array []interface{}) *BstNode {
	if len(array) == 0 {
		return nil
	}

	mid := len(array) / 2
	node := &BstNode{value: array[mid]}
	if mid > 0 {
		node.left = buildBalanceBST(array[0:mid])
		node.right = buildBalanceBST(array[mid+1:])
	}
	return node
}

func (t *BstTree) InorderTraverse(visitor func(interface{})) {
	if t.root != nil {
		//	visitor(  fmt.Sprintf( "Depth is %d ", t.root.Depth()))
		t.root.InorderTraverse(visitor)
	} else {
		visitor(" Empty Tree")
	}
}

func (t *BstTree) Balance() {
	if t.root == nil {
		return
	}
	collected := []interface{}{}
	t.PostorderTraverse(func(value interface{}) {
		collected = append(collected, value)
	})
	t.root = buildBalanceBST(collected)

}

func (t *BstTree) PostorderTraverse(visitor func(interface{})) {
	if t.root != nil {
		t.root.PostorderTraverse(visitor)
	} else {
		visitor(" Empty Tree")
	}
}

func (t *BstNode) InorderTraverse(visitor func(interface{})) {
	visitor(t.value)

	if t.left != nil {

		t.left.InorderTraverse(visitor)
	}
	if t.right != nil {

		t.right.InorderTraverse(visitor)
	}
}

func (t *BstNode) PostorderTraverse(visitor func(interface{})) {
	if t.left != nil {
		t.left.PostorderTraverse(visitor)
	}
	visitor(t.value)
	if t.right != nil {
		t.right.PostorderTraverse(visitor)
	}
}

func (t *BstTree) Depth() int {
	if t.root == nil {
		return 0
	}
	return t.root.Depth()
}

func (t *BstNode) Depth() int {
	l, r := 0, 0
	if t.left != nil {
		l = t.left.Depth()
	}
	if t.right != nil {
		r = t.right.Depth()
	}
	if l < r {
		return r + 1
	} else {
		return l + 1
	}
}
