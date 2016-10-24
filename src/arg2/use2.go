package main

import (
	"fmt"
	"quicksort"
)

func main() {
	array := []int{32, 41, 56, 11, 77, 32, 45, 23, 67, 12, 16, 21}
	fmt.Println("before:",array)
	num:=quicksort.Quicksort(array, 1, len(array))
	fmt.Println("after:",num)
}
