package linkedlist

import (
	"fmt"
	"testing"
)

func TestList_HasCircle(t *testing.T) {
	l := NewList()
	for i := 1; i < 10; i++ {
		l.InsertToTail(i)
	}
	fmt.Println(l.HasCircle())
	l.tail.next = l.FindByIndex(4)
	fmt.Println(l.HasCircle())
}
