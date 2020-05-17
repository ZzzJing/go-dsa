package linkedlist

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

type List struct {
	head *Node
	tail *Node
	size uint64
}

func (node *Node) GetNext() *Node {
	return node.next
}

func (node *Node) GetValue() interface{} {
	return node.data
}

func (list *List) InsertAfter(n *Node, d interface{}) bool {
	if d == nil {
		return false
	}
	newNode := NewNode(d)
	// Do insert
	oldNext := n.next
	n.next = newNode
	newNode.next = oldNext
	if oldNext == nil {
		list.tail = newNode
	}
	list.size++
	return true
}

func (list *List) InsertBefore(n *Node, d interface{}) bool {
	if d == nil || n == list.head {
		return false
	}
	cur := list.head.next
	pre := list.head
	for nil != cur {
		if cur == n {
			break
		}
		pre = cur
		cur = cur.next
	}
	if cur == nil {
		return false
	}
	newNode := NewNode(d)
	// Do insert
	pre.next = newNode
	newNode.next = cur
	list.size++
	return true
}

func (list *List) InsertAfterHead(d interface{}) bool {
	return list.InsertAfter(list.head, d)
}

func (list *List) InsertToTail(d interface{}) bool {
	return list.InsertAfter(list.tail, d)
}

func (list *List) Insert(index uint64, d interface{}) bool {
	if index == 0 || index > list.size {
		return false
	}

	pre := list.head
	var i uint64 = 1
	for ; i < index; i++ {
		pre = pre.next
	}
	return list.InsertAfter(pre, d)
}

func (list *List) FindByIndex(index uint64) *Node {
	if index > list.size {
		return nil
	}
	cur := list.head
	var i uint64 = 0
	for ; i < index; i++ {
		cur = cur.next
	}
	return cur
}

func (list *List) DeleteNode(n *Node) bool {
	if n == nil || n == list.head {
		return false
	}
	pre := list.head
	cur := pre.next
	for cur != nil {
		if cur == n {
			break
		}
		pre = cur
		cur = cur.next
	}
	if cur == nil {
		return false
	}
	if cur.next == nil {
		pre.next = nil
		list.tail = pre
	} else {
		pre.next = cur.next
	}
	n = nil
	return true
}

func (list *List) Print() {
	cur := list.head.next
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}

func NewNode(d interface{}) *Node {
	return &Node{data: d}
}

func NewList() *List {
	sentry := NewNode(0)
	return &List{
		size: 1,
		head: sentry,
		tail: sentry,
	}
}
