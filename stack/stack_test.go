package stack

import (
	"fmt"
	"testing"
)

func TestStack_IsEmpty(t *testing.T) {
	s := NewStack()

	s.Print()
}

func TestStack_Push(t *testing.T) {
	s := NewStack()

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Print()
}

func TestStack_Pop(t *testing.T) {
	s := NewStack()

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Print()
	fmt.Println(s.Pop())
	fmt.Println("pop after:")
	s.Print()
}

func TestStack_Flash(t *testing.T) {
	s := NewStack()

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Print()
	s.Flash()
	s.Print()
}
