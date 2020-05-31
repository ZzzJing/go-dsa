package queue

import "fmt"

type Queue struct {
	capacity int
	data     []interface{}
	head     int
	tail     int
}

func (q *Queue) EnQueue(v interface{}) bool {
	if q.tail == q.capacity {
		if q.head == 0 {
			return false
		}

		for i := q.head; i < q.tail; i++ {
			fmt.Println(i)
			q.data[i-q.head] = q.data[i]
		}
		q.tail -= q.head
		q.head = 0
	}

	//			q.data[i-q.head] = q.data[i]
	q.data[q.tail] = v
	q.tail++
	return true
}

func (q *Queue) DeQueue() interface{} {
	if q.head == q.tail {
		return nil
	}

	v := q.data[q.head]
	q.head++
	return v
}

func (q *Queue) Print() {
	if q.head == q.tail {
		fmt.Println("empty queue")
	}
	result := ""
	for i := q.head; i < q.tail; i++ {
		result += fmt.Sprintf("<-%+v", q.data[i])
	}
	fmt.Println(result)
}

func NewQueue(n int) *Queue {
	return &Queue{
		data:     make([]interface{}, n),
		capacity: n,
		head:     0,
		tail:     0,
	}
}
