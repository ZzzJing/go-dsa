package linkedlist

import (
	"fmt"
	"testing"
)

func TestList_IsPalindrome(t *testing.T) {
	l := NewList()
	l.InsertToTail(1)
	l.InsertToTail(3)
	l.InsertToTail(2)
	l.InsertToTail(2)
	l.InsertToTail(3)
	l.InsertToTail(1)
	l.Print()

	fmt.Println(l.IsPalindrome())
	l.Print()
}

func TestList_IsPalindrome2(t *testing.T) {
	l := NewList()
	l.InsertToTail(1)
	l.InsertToTail(3)
	l.InsertToTail(2)
	l.InsertToTail(3)
	l.InsertToTail(1)
	l.Print()

	fmt.Println(l.IsPalindrome())
	l.Print()
}
