package main

import (
	"fmt"
	"go-dsal/bst"
)

func main()  {
	t := bst.BST{}
	nums := []int{5, 3, 6, 2, 4, 8}
	for _, v := range nums {
		t.Add(v)
	}
	ret := t.RemoveMax()
	fmt.Println(ret)
	fmt.Println(t.GetSize())
}
