package queue

import (
	"fmt"
	"testing"
)

func TestNewCircularQueue(t *testing.T) {
	q := NewCircularQueue(5)
	fmt.Println(q)
}

func TestNewCircularQueue2(t *testing.T) {
	q := NewCircularQueue(8)
	fmt.Println(q)
}

func TestCircularQueue_EnQueue(t *testing.T) {
	q := NewCircularQueue(8)
	for i := 1; i < 9; i++ {
		q.EnQueue(i)
	}
	q.Print()
}

func TestCircularQueue_DeQueue(t *testing.T) {
	q := NewCircularQueue(8)
	for i := 1; i < 9; i++ {
		q.EnQueue(i)
	}
	q.Print()
	fmt.Println("dequeue:" + fmt.Sprintf("%+v", q.DeQueue()))
	q.EnQueue(10)
	q.Print()
}
