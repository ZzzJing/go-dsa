package sort

//--------------------------------------------o(n^2)------------------------------------------

//冒泡排序
func BubbleSort(a []int, n int) {
	if n <= 1 {
		return
	}

	for i := 0; i < n; i++ {
		flag := false
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				flag = true
			}
		}

		//flag为false证明这次循环没有调换位置,数组已经排好序
		if !flag {
			break
		}
	}
}

//插入排序
func InsertSort(a []int, n int) {
	if n <= 1 {
		return
	}

	for i := 1; i < n; i++ {
		v := a[i]
		j := i - 1
		for ; j >= 0; j-- {
			if a[j] > v {
				a[j+1] = a[j]
			} else {
				break
			}
		}
		a[j+1] = v
	}
}

//选择排序
func SelectionSort(a []int, n int) {
	if n <= 1 {
		return
	}

	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if a[j] < a[min] {
				min = j
			}
		}

		if i != min {
			a[i], a[min] = a[min], a[i]
		}
	}
}

//--------------------------------------------o(nlogn)------------------------------------------

//归并排序, 其关键点在于merge函数, 一定会运行c*n*logn遍,所以最好情况、最坏情况，还是平均情况，时间复杂度都是 O(nlogn)
//是稳定排序
//空间复杂度O(n),需要额外空间,不是原地排序,无法用在大数组排序
func MergeSort(a []int, n int) {
	if n <= 1 {
		return
	}

	mergeSort(a, 0, n-1)
}

func mergeSort(a []int, start, end int) {
	if start >= end {
		return
	}

	mid := (start + end) / 2
	mergeSort(a, start, mid)
	mergeSort(a, mid+1, end)

	merge(a, start, mid, end)
}

func merge(a []int, start, mid, end int) {
	tmp := make([]int, end-start+1)
	i := start
	j := mid + 1
	k := 0
	for ; i <= mid && j <= end; k++ {
		if a[i] <= a[j] {
			tmp[k] = a[i]
			i++
		} else {
			tmp[k] = a[j]
			j++
		}
	}

	for ; i <= mid; i++ {
		tmp[k] = a[i]
		k++
	}

	for ; j <= end; j++ {
		tmp[k] = a[j]
		k++
	}

	copy(a[start:end+1], tmp)
}

//快排
func QuickSort(a []int, n int) {
	if n <= 1 {
		return
	}

	quickSort(a, 0, n-1)
}

func quickSort(a []int, start, end int) {
	if start >= end {
		return
	}

	pivot := partition2(a, start, end)
	quickSort(a, start, pivot-1)
	quickSort(a, pivot+1, end)
}

func partition(a []int, start, end int) int {
	pivot := a[end]
	i := start
	for j := start; j < end; j++ {
		if a[j] < pivot {
			if i != j {
				a[i], a[j] = a[j], a[i]
			}
			i++
		}
	}
	a[i], a[end] = a[end], a[i]
	return i
}

func partition2(a []int, start, end int) int {
	pivot := a[end]
	i := start
	j := end
	for {
		for a[i] < pivot && i < j {
			i++
		}
		for a[j] >= pivot && i < j {
			j--
		}
		if i < j {
			a[i], a[j] = a[j], a[i]
		} else {
			break
		}
	}

	a[i], a[end] = a[end], a[i]
	return i
}
