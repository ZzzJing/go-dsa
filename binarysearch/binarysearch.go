package binarysearch

func BinarySearch(a []int, v int) int {
	alen := len(a)
	low := 0
	high := alen - 1

	for low <= high {
		mid := low + ((high - low) >> 1)
		if v < a[mid] {
			high = mid - 1
			continue
		} else if v > a[mid] {
			low = mid + 1
			continue
		} else {
			return mid
		}
	}

	return -1
}

func BinarySearchFirst(a []int, v int) int {
	alen := len(a)
	low := 0
	high := alen - 1

	for low <= high {
		mid := low + ((high - low) >> 1)
		if v < a[mid] {
			high = mid - 1
			continue
		} else if v > a[mid] {
			low = mid + 1
			continue
		} else {
			if a[mid-1] < v || mid == 0 {
				return mid
			} else {
				high = mid - 1
			}
		}
	}

	return -1
}

func BinarySearchFirstGT(a []int, v int) int {
	alen := len(a)
	low := 0
	high := alen - 1

	for low <= high {
		mid := low + ((high - low) >> 1)
		if v <= a[mid] {
			if a[mid-1] < v || mid == 0 {
				return mid
			} else {
				high = mid - 1
			}
		} else {
			low = mid + 1
		}
	}

	return -1
}
