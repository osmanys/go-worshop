package exercises_test

import (
	"testing"

	"../exercises"
)

func TestAddTree(t *testing.T) {
	var bst = buildTree()

	var result []int
	bst.InOrderTraverse(func(i int) {
		result = append(result, i)
	})
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if !isEqualSlice(result, expected) {
		t.Errorf("Traversal order incorrect, got %v", result)
	}
}

func isEqualSlice(a, b []int) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func buildTree() *exercises.BSTree {
	var bst = &exercises.BSTree{}
	bst.Insert(8)
	bst.Insert(4)
	bst.Insert(10)
	bst.Insert(2)
	bst.Insert(6)
	bst.Insert(1)
	bst.Insert(3)
	bst.Insert(5)
	bst.Insert(7)
	bst.Insert(9)
	return bst
}
