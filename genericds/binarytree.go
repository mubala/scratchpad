package genericds

import (
	"errors"
)

type binaryTreeNode struct {
	Value       interface{}
	Right, Left *binaryTreeNode
}

type Visitor func(interface{})

type BinaryTree interface {
	InorderTraverse(visitor Visitor)
	PreorderTraverse(visitor Visitor)
	PostorderTraverse(visitor Visitor)

	SetChildren(right, left BinaryTree) error
	RightChild() BinaryTree
	LeftChild() BinaryTree
}

func (b *binaryTreeNode) SetChildren(right, left BinaryTree) error {
	var conversionOK bool
	if right != nil {
		b.Right, conversionOK = right.(*binaryTreeNode)
		if !conversionOK {
			return errors.New("Wrong right argument passed ")
		}
	}
	if left != nil {
		b.Left, conversionOK = left.(*binaryTreeNode)
		if !conversionOK {
			return errors.New("Wrong left argument passed")
		}
	}
	return nil
}

func (b *binaryTreeNode) RightChild() BinaryTree {
	if b.Right != nil {
		return b.Right
	}
	return nil
}

func (b *binaryTreeNode) LeftChild() BinaryTree {
	if b.Left != nil {
		return b.Left
	}
	return nil
}

//Factory method
func NewBinaryTree(value interface{}) BinaryTree {
	return &binaryTreeNode{Value: value}
}

func (b *binaryTreeNode) PreorderTraverse(visitor Visitor) {
	visitor(b.Value)
	if b.Right != nil {
		b.RightChild().PreorderTraverse(visitor)
	}
	if b.LeftChild() != nil {
		b.LeftChild().PreorderTraverse(visitor)
	}
}

func (b *binaryTreeNode) InorderTraverse(visitor Visitor) {

	if b.RightChild() != nil {
		b.RightChild().InorderTraverse(visitor)
	}
	visitor(b.Value)

	if b.LeftChild() != nil {
		b.LeftChild().InorderTraverse(visitor)
	}
}

func (b *binaryTreeNode) PostorderTraverse(visitor Visitor) {

	if b.RightChild() != nil {
		b.RightChild().PostorderTraverse(visitor)
	}
	if b.LeftChild() != nil {
		b.LeftChild().PostorderTraverse(visitor)
	}
	visitor(b.Value)
}
