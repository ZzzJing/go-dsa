package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func generateArray(n int) []int {
	rand.Seed(time.Now().Unix())

	var a = make([]int, n)
	for i := 0; i < len(a); i++ {
		a[i] = rand.Intn(1000)
	}
	return a
}

func TestBubbleSort(t *testing.T) {
	a := generateArray(10)
	fmt.Println("pre:", a)
	BubbleSort(a, len(a))
	fmt.Println("after:", a)
}

func TestInsertSort(t *testing.T) {
	a := generateArray(10)
	fmt.Println("pre:", a)
	InsertSort(a, len(a))
	fmt.Println("after:", a)
}

func TestSelectionSort(t *testing.T) {
	a := generateArray(10)
	fmt.Println("pre:", a)
	SelectionSort(a, len(a))
	fmt.Println("after:", a)
}

func TestMergeSort(t *testing.T) {
	a := generateArray(10)
	fmt.Println("pre:", a)
	MergeSort(a, len(a))
	fmt.Println("after:", a)
}

func TestQuickSort(t *testing.T) {
	a := generateArray(10)
	fmt.Println("pre:", a)
	QuickSort(a, len(a))
	fmt.Println("after:", a)
}
