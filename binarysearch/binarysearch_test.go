package binarysearch

import (
	"fmt"
	"go-dsa/sort"
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

func TestBinarySearch(t *testing.T) {
	a := generateArray(100)
	sort.CountSort(a)
	k := BinarySearch(a, a[19])
	fmt.Println(k)
}

func TestBinarySearchFirst(t *testing.T) {
	a := []int{1, 2, 4, 5, 5, 5, 5, 5, 6, 8, 9}
	k := BinarySearchFirst(a, 5)
	fmt.Println(k)
}

func TestBinarySearchFirstGT(t *testing.T) {
	a := generateArray(10000)
	sort.CountSort(a)
	k := BinarySearchFirstGT(a, a[100])
	fmt.Println(k)
	for k, v := range a {
		if v == a[100] {
			fmt.Println(k)
		}
	}
}
