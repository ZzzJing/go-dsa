package linkedlist

import (
	"fmt"
	"testing"
)

func TestList_InsertAfter(t *testing.T) {
	l := NewList()
	for i := 10; i > 0; i-- {
		l.InsertAfter(l.head, i)
	}

	l.Print()
}

func TestList_InsertBefore(t *testing.T) {
	l := NewList()
	l.InsertAfter(l.head, 10)

	for i := 1; i < 10; i++ {
		l.InsertBefore(l.tail, i)
	}

	l.Print()
}

func TestList_InsertAfterHead(t *testing.T) {
	l := NewList()
	for i := 10; i > 0; i-- {
		l.InsertAfterHead(i)
	}

	l.Print()
}

func TestList_InsertToTail(t *testing.T) {
	l := NewList()
	for i := 1; i <= 10; i++ {
		l.InsertToTail(i)
	}

	l.Print()
}

func TestList_Insert(t *testing.T) {
	l := NewList()
	var i uint64 = 10
	for ; i > 0; i-- {
		l.Insert(1, i)
	}

	l.Print()
}

func TestList_FindByIndex(t *testing.T) {
	l := NewList()
	for i := 1; i <= 10; i++ {
		l.InsertToTail(i)
	}

	l.Print()

	var j uint64 = 1
	format := ""
	for ; j <= 10; j++ {
		cur := l.FindByIndex(j)
		format += fmt.Sprintf("%+v", cur.GetValue())
		if nil != cur.next {
			format += "->"
		}
	}
	fmt.Println(format)
}

func TestList_DeleteNode(t *testing.T) {
	l := NewList()
	for i := 1; i <= 10; i++ {
		l.InsertToTail(i)
	}

	l.Print()

	deleteList := []*Node{}
	var j uint64 = 1
	for ; j < 10; j += 2 {
		deleteList = append(deleteList, l.FindByIndex(j))
	}

	for _, v := range deleteList {
		l.DeleteNode(v)
	}
	l.Print()
}
