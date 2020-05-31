package queue

import "fmt"

type CircularQueue struct {
	capacity int
	data     []interface{}
	head     int
	tail     int
}

func (q *CircularQueue) EnQueue(v interface{}) bool {
	if (q.tail+1)&(q.capacity-1) == q.head { //capacity为2的n次方,使用与操作代替取余
		fmt.Println("queue is full")
		return false
	}

	q.data[q.tail] = v
	q.tail = (q.tail + 1) & (q.capacity - 1)
	return true
}

func (q *CircularQueue) DeQueue() interface{} {
	if q.head == q.tail {
		return nil
	}

	r := q.data[q.head]
	q.head = (q.head + 1) & (q.capacity - 1)
	return r
}

func (q *CircularQueue) Print() {
	if q.head == q.tail {
		fmt.Println("empty queue")
	}
	result := ""
	i := q.head
	for true {
		result += fmt.Sprintf("<-%+v", q.data[i])
		i = (i + 1) & (q.capacity - 1)
		if i == q.tail {
			break
		}
	}

	fmt.Println(result)
}

func NewCircularQueue(n int) *CircularQueue {
	if n <= 0 || n&(n-1) != 0 { //判断capacity是否为2的n次方
		fmt.Println("capacity should be 2^x")
		return nil
	}
	return &CircularQueue{
		capacity: n,
		data:     make([]interface{}, n),
		head:     0,
		tail:     0,
	}
}
