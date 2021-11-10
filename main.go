package main

import (
	"fmt"
	"go-dsal/sort"
)

func main()  {
	nums := []int{5, 3, 6, 2, 4, 8}
	sort.SelectionSort(nums)
	fmt.Println(nums)
}
