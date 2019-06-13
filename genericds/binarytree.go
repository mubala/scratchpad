package genericds

import (
	"bytes"
	"errors"
	"fmt"
)

type binaryTreeNode struct {
	Value       interface{}
	Right, Left *binaryTreeNode
}

type Visitor func(interface{})

type BinaryTree interface {
	InorderTraverse(visitor Visitor) string
	PreorderTraverse(visitor Visitor) string
	PostorderTraverse(visitor Visitor) string

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

func (b *binaryTreeNode) PreorderTraverse(visitor Visitor) string {
	buffer := bytes.Buffer{}
	buffer.WriteString(fmt.Sprint(b.Value))
	visitor(b.Value)
	if b.Right != nil {
		buffer.WriteString(", ")
		buffer.WriteString(b.RightChild().PreorderTraverse(visitor))
	}
	if b.LeftChild() != nil {
		buffer.WriteString(", ")
		buffer.WriteString(b.LeftChild().PreorderTraverse(visitor))
	}

	return buffer.String()
}

func (b *binaryTreeNode) InorderTraverse(visitor Visitor) string {
	buffer := bytes.Buffer{}

	if b.RightChild() != nil {
		buffer.WriteString(b.RightChild().InorderTraverse(visitor))
		buffer.WriteString(", ")
	}
	buffer.WriteString(fmt.Sprint(b.Value))
	visitor(b.Value)

	if b.LeftChild() != nil {
		buffer.WriteString(", ")
		buffer.WriteString(b.LeftChild().InorderTraverse(visitor))
	}
	return buffer.String()
}

func (b *binaryTreeNode) PostorderTraverse(visitor Visitor) string {
	buffer := bytes.Buffer{}

	if b.RightChild() != nil {
		buffer.WriteString(b.RightChild().PostorderTraverse(visitor))
		buffer.WriteString(" ")
	}

	if b.LeftChild() != nil {
		buffer.WriteString(" ")
		buffer.WriteString(b.LeftChild().PostorderTraverse(visitor))
	}

	buffer.WriteString(fmt.Sprint(b.Value))
	visitor(b.Value)

	return buffer.String()
}
