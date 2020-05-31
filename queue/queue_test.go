package queue

import (
	"fmt"
	"testing"
)

func TestQueue_EnQueue(t *testing.T) {
	q := NewQueue(3)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)

	q.Print()
}

func TestQueue_DeQueue(t *testing.T) {
	q := NewQueue(3)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)

	q.Print()

	fmt.Println("dequeue:" + fmt.Sprintf("%+v", q.DeQueue()))
	fmt.Println("dequeue:" + fmt.Sprintf("%+v", q.DeQueue()))
	q.Print()
}

func TestQueue_EnQueue2(t *testing.T) {
	q := NewQueue(3)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	fmt.Println(q.EnQueue(4))
	q.Print()
	fmt.Println(q.head, q.tail)
	fmt.Println("dequeue:" + fmt.Sprintf("%+v", q.DeQueue()))
	q.Print()
	fmt.Println(q.head, q.tail)
	fmt.Println(q.EnQueue(4))
	q.Print()
}
