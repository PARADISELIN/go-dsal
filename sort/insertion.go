package sort

func InsertionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i; j - 1 >= 0 && arr[j] < arr[j - 1]; j-- {
			arr[j], arr[j - 1] = arr[j - 1], arr[j]
		}
	}
}

// optimization: reduce the cost of each exchange operation
func insertionSortOptimized(arr []int) {
	for i := 0; i < len(arr); i++ {
		cur := arr[i]
		var j int
		for j = i; j - 1 >= 0 && cur < arr[j - 1]; j-- {
			arr[j] = arr[j - 1]
		}
		arr[j] = cur
	}
}