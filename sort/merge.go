package sort

// 合并两个有序的区间 arr[l, mid] 和 arr[mid + 1, r]
// 注意这里是闭区间
func merge(arr []int, l int, mid int, r int) {
	temp := make([]int, r - l + 1)
	copy(temp, arr[l:r + 1])

	i := l
	j := mid + 1

	// 每轮循环为arr[k]赋值
	for k := l; k <= r; k++ {
		if i > mid {
			arr[k] = temp[j - l]
			j++
		} else if j > r {
			arr[k] = temp[i - l]
			i++
		} else if temp[i - l] <= temp[j - l] {
			arr[k] = temp[i - l]
			i++
		} else {
			arr[k] = temp[j - l]
			j++
		}
	}
}

func mergeSort(arr []int, l int, r int) {
	if l >= r {
		return
	}

	mid := l + (r - l) / 2
	mergeSort(arr, l , mid)
	mergeSort(arr, mid + 1, r)
	merge(arr, l, mid, r)
}

func MergeSort(arr []int) {
	mergeSort(arr, 0, len(arr) - 1)
}
