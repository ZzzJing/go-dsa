package sort

func getMax(a []int) int {
	max := a[0]
	for i := 1; i < len(a); i++ {
		if max < a[i] {
			max = a[i]
		}
	}

	return max
}

//计数排序
func CountSort(a []int) {
	aLen := len(a)
	if aLen <= 1 {
		return
	}

	vList := make([]int, getMax(a)+1)
	for i := 0; i < aLen; i++ {
		vList[a[i]]++
	}
	rList := make([]int, 0)
	for j := 0; j < len(vList); j++ {
		for vList[j] > 0 {
			rList = append(rList, j)
			vList[j]--
		}
	}

	copy(a, rList)
}

//桶排序
func BucketSort(a []int) {
	aLen := len(a)
	if aLen <= 1 {
		return
	}

	max := getMax(a)
	//这里将桶的个数设置为aLen,时间复杂度为m*(n/m)log(n/m),这里n=m=aLen,所以为o(n)
	//实际算法中桶的个数选择非常重要
	var index int
	bucket := make([][]int, aLen)
	for i := 0; i < aLen; i++ {
		index = a[i] * (aLen - 1) / max // 计算桶的序号
		bucket[index] = append(bucket[index], a[i])
	}

	tmpIndex := 0
	for j := 0; j < aLen; j++ {
		bucketLen := len(bucket[j])
		if bucketLen > 0 {
			QuickSort(bucket[j], bucketLen)
			copy(a[tmpIndex:], bucket[j])
			tmpIndex += bucketLen
		}
	}
}
