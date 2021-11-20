package search

func BinarySearch(data []int, target int) int {
	l := 0
	r := len(data) - 1

	for l <= r {
		mid := l + (r - l) / 2
		if data[mid] == target {
			return mid
		}
		if data[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}
