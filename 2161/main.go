package main

import "fmt"

func pivotArray(nums []int, pivot int) []int {
	smallerArr, biggerArr := []int{}, []int{}
	pivotCount := 0

	for _, num := range nums {
		if num < pivot {
			smallerArr = append(smallerArr, num)
		} else if num > pivot {
			biggerArr = append(biggerArr, num)
		} else {
			pivotCount++
		}
	}

	for i := 0; i < pivotCount; i++ {
		smallerArr = append(smallerArr, pivot)
	}

	return append(smallerArr, biggerArr...)
}
func main() {
	fmt.Println(pivotArray([]int{9, 12, 5, 10, 14, 3, 10}, 10))
}
