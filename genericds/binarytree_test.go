package genericds

import (
	"reflect"
	"testing"
)

// Builds below tree
//         1
//       /   \
//      2     3
//     /    /   \
//    4    5     6
//                \
//                 7
func buildTrees() []BinaryTree {
	nodes := make([]BinaryTree, 0, 7)
	for i := 1; i < 8; i++ {
		nodes = append(nodes, NewBinaryTree(i))

	}

	nodes[0].SetChildren(nodes[1], nodes[2])
	nodes[1].SetChildren(nodes[3], nil)
	nodes[2].SetChildren(nodes[4], nodes[5])
	nodes[5].SetChildren(nil, nodes[6])
	return nodes
}

func TestBinaryTree_Traverse(t *testing.T) {
	trees := buildTrees()
	collected := make([]int, 0, 7)
	testVisitor := func(value interface{}) {
		if intValue, ok := value.(int); ok {
			collected = append(collected, intValue)
		}
	}
	tests := []struct {
		name      string
		tree      BinaryTree
		inorder   []int
		preorder  []int
		postorder []int
	}{
		{
			name:      " Test whole tree ",
			tree:      trees[0],
			preorder:  []int{1, 2, 4, 3, 5, 6, 7},
			inorder:   []int{4, 2, 1, 5, 3, 6, 7},
			postorder: []int{4, 2, 5, 7, 6, 3, 1},
		},
		{
			name:      " Test Tree with depth 1 ",
			tree:      trees[3],
			preorder:  []int{4},
			inorder:   []int{4},
			postorder: []int{4},
		},
		{
			name:      " Test Tree with depth 2 , one child ",
			tree:      trees[1],
			preorder:  []int{2, 4},
			inorder:   []int{4, 2},
			postorder: []int{4, 2},
		},
		{
			name:      " Test Tree with depth 2 , one child ",
			tree:      trees[5],
			preorder:  []int{6, 7},
			inorder:   []int{6, 7},
			postorder: []int{7, 6},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			collected = collected[:0]
			t.Log("Running inorder ", test.name)
			if got := test.tree.InorderTraverse(testVisitor); !reflect.DeepEqual(collected, test.inorder) {
				t.Errorf(" %s Failed  Response: %v  Collected: %v != inorder  %v", test.name, got, collected, test.inorder)
			}
			collected = collected[:0]
			t.Log("Running preorder", test.name)
			if got := test.tree.PreorderTraverse(testVisitor); !reflect.DeepEqual(collected, test.preorder) {
				t.Errorf(" %s Failed  Response: %v  Collected: %v != inorder  %v", test.name, got, collected, test.preorder)
			}
			collected = collected[:0]
			t.Log("Running postorder", test.name)
			if got := test.tree.PostorderTraverse(testVisitor); !reflect.DeepEqual(collected, test.postorder) {
				t.Errorf(" %s Failed  Response: %v  Collected: %v != inorder  %v", test.name, got, collected, test.postorder)
			}
		})
	}
}

type incompatibleTree struct{}

func (b *incompatibleTree) SetChildren(right, left BinaryTree) error {
	return nil
}

func (b *incompatibleTree) RightChild() BinaryTree {
	return nil
}

func (b *incompatibleTree) LeftChild() BinaryTree {
	return nil
}

//Factory method
func NewIncompatibleBinaryTree() BinaryTree {
	return &incompatibleTree{}
}

func (b *incompatibleTree) PreorderTraverse(visitor Visitor) string {
	return ""
}

func (b *incompatibleTree) InorderTraverse(visitor Visitor) string {
	return ""
}

func (b *incompatibleTree) PostorderTraverse(visitor Visitor) string {
	return ""
}

func TestBinaryTree_SetChildren(t *testing.T) {
	tree := NewBinaryTree(0)
	t.Run(" Test Error case", func(t *testing.T) {
		if err := tree.SetChildren(nil, NewIncompatibleBinaryTree()); err == nil {
			t.Errorf(" Failed incompatability check ... ")
		}

		if err := tree.SetChildren(NewIncompatibleBinaryTree(), nil); err == nil {
			t.Errorf(" Failed incompatability check ... ")
		}

	})
}
