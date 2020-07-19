package genericds

import (
	"errors"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func createIntBST() BinarySearchTree {
	tree := NewBinarySearchTree(func(v1, v2 interface{}) (int, error) {
		a, a_is_invalid := v1.(int)
		b, b_is_invalid := v2.(int)
		if a_is_invalid && b_is_invalid {
			return a - b, nil
		}
		return -1, errors.New(" Wrong type is passed , it has to be int")
	})
	return tree
}

func Test_BSTTreeBalance(t *testing.T) {
	tree := createIntBST()
	count := 20
	t.Log(" In Testc_BSTTreeBalance ")
	for i := 0; i < count; i++ {
		tree.Add(i)
	}
	depthBeforeBalance := tree.Depth()
	if count != depthBeforeBalance {
		t.Errorf(" The tree depth has to be 20, but found %d ", depthBeforeBalance)
	}
	tree.Balance()
	depthAfterBalance := tree.Depth()

	if depthAfterBalance > depthBeforeBalance {
		t.Errorf(" The tree has not been balanced %d < %d", depthAfterBalance, depthBeforeBalance)
	}

}

func Test_BSTTree(t *testing.T) {
	tree := createIntBST()
	count := 20
	collected := []interface{}{}
	t.Log(" In testcaseBalance ")
	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < count; i++ {
		tree.Add(r1.Intn(count))
	}

	var lastValue interface{} = nil
	tree.PostorderTraverse(func(value interface{}) {
		collected = append(collected, value)
		if lastValue != nil {
			a := lastValue.(int)
			b := value.(int)
			if a > b {
				t.Errorf(" The value is not in ordered list %d < %d", a, b)
			}
		}
		lastValue = value
	})
	t.Log(" Sorted Array is ", collected)
}
func Test_EmptyBSTree(t *testing.T) {
	tree := NewBinarySearchTree(nil)
	tree.Balance()
	if 0 != tree.Depth() {
		t.Errorf(" The depth has to be 0 for the empty tree")
	}
	emptyCheckVisitor := func(value interface{}) {
		if value.(string) != " Empty Tree" {
			t.Errorf(" Empty checked on visitor function ")
		}
	}
	tree.InorderTraverse(emptyCheckVisitor)
	tree.PostorderTraverse(emptyCheckVisitor)
	_, err := tree.Find("anything")
	if err == nil {
		t.Error(" `anything` is found in Empty tree , broken functionality ")
	}
}

func Test_BstTreeFind(t *testing.T) {
	tree := createIntBST()
	tree.Add(20)
	tree.Add(30)
	tree.Add(5)
	tree.Add(40)

	//        20
	//     5     30
	//               40
	// Inorder traversal is 20 5 30 40

	collected := []interface{}{}

	expected := []interface{}{20, 5, 30, 40}
	tree.InorderTraverse(func(value interface{}) {
		collected = append(collected, value)
	})
	if !reflect.DeepEqual(expected, collected) {
		t.Errorf(" InorderTraverse failed  expected %v, collected %v", expected, collected)
	}
	found, err := tree.Find(40)
	if err != nil {
		t.Error(" Find method failed with err ", err)
	}
	if found != true {
		t.Error(" Find method failed to find value 40 ")
	}

	found, err = tree.Find(5)
	if err != nil {
		t.Error(" Find method failed with err ", err)
	}
	if found != true {
		t.Error(" Find method failed to find value 5 ")
	}

	found, err = tree.Find(500)
	if err != nil {
		t.Error(" Find method failed with err ", err)
	}
	if found == true {
		t.Error(" Find method failed and found unstored  value 500 ")
	}

	// Check error case
	found, err = tree.Find("incompatible value")
	if err == nil {
		t.Error(" Find is not checking `incomptible value`")
	}

	_, err = tree.Add("incompatible value")
	if err == nil {
		t.Error(" Add is not checking `incomptible value`")
	}
	t1 := &BstTree{root: &BstNode{value: 10}}
	_, err = t1.Add(" broken tree ")
	if err == nil {
		t.Error(" Add is not checking `broken tree`")
	}
}
