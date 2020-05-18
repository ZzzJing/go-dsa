package stack

import "fmt"

type Stack struct {
	//slice based
	data []interface{}
	top  int64
}

func (s *Stack) IsEmpty() bool {
	return s.top < 0
}

func (s *Stack) Push(v interface{}) {
	s.top++
	s.data = append(s.data, v)
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}

	v := s.data[s.top]
	s.top--
	return v
}

func (s *Stack) Flash() {
	s.top = -1
}

func (s *Stack) Print() {
	if s.IsEmpty() {
		fmt.Println("empty statck")
	} else {
		for i := s.top; i >= 0; i-- {
			fmt.Println(s.data[i])
		}
	}
}

func NewStack() *Stack {
	return &Stack{
		data: make([]interface{}, 0, 64),
		top:  -1,
	}
}
