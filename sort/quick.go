package sort

import (
	"math/rand"
	"time"
)

func randomInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max - min + 1) + min
}

func partition(arr []int, l int, r int) int {
	p := randomInt(l, r)
	arr[l], arr[p] = arr[p], arr[l]

	// arr[l+1...j] < v; arr[j+1...i] >= v
	j := l
	for i := l + 1; i <= r; i++ {
		if arr[i] < arr[l] {
			j++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[l], arr[j] = arr[j], arr[l]
	return j
}

func quickSort(arr []int, l int, r int) {
	if l >= r {
		return
	}

	pivot := partition(arr, l, r)
	quickSort(arr, l, pivot - 1)
	quickSort(arr, pivot + 1, r)
}

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr) - 1)
}