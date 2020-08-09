package tree

import (
	"fmt"
	"testing"
)

func compareFunc(v, nodeV interface{}) int {
	var nv, nnodeV int
	switch v.(type) {
	case int:
		nv = v.(int)
	}

	switch nodeV.(type) {
	case int:
		nnodeV = nodeV.(int)
	}
	return nv - nnodeV
}

func TestBST_Find(t *testing.T) {
	bst := NewBST(10, compareFunc)
	bst.Insert(5)
	bst.Insert(8)
	bst.Insert(7)
	bst.Insert(6)
	bst.Insert(15)
	bst.Insert(13)
	bst.Insert(18)
	bst.Insert(11)
	fmt.Println(bst.Find(13))
}

func TestBST_Delete(t *testing.T) {
	bst := NewBST(10, compareFunc)
	bst.Insert(5)
	bst.Insert(8)
	bst.Insert(7)
	bst.Insert(6)
	bst.Insert(15)
	bst.Insert(13)
	bst.Insert(18)
	bst.Insert(11)
	fmt.Println(bst.Find(13))
	bst.InOrderTraverse(bst.root)
	deleteSuccess := bst.Delete(13)
	fmt.Println(deleteSuccess)
	fmt.Println(bst.Find(13))
	bst.InOrderTraverse(bst.root)
	deleteSuccess = bst.Delete(13)
	fmt.Println(deleteSuccess)
}

func TestBST_PreOrderTraverse(t *testing.T) {
	bst := NewBST(10, compareFunc)
	bst.Insert(5)
	bst.Insert(8)
	bst.Insert(7)
	bst.Insert(6)
	bst.Insert(15)
	bst.Insert(13)
	bst.Insert(18)
	bst.Insert(11)
	bst.PreOrderTraverse(bst.root)
}

func TestBST_InOrderTraverse(t *testing.T) {
	bst := NewBST(10, compareFunc)
	bst.Insert(5)
	bst.Insert(8)
	bst.Insert(7)
	bst.Insert(6)
	bst.Insert(15)
	bst.Insert(13)
	bst.Insert(18)
	bst.Insert(11)
	bst.InOrderTraverse(bst.root)
}
