package main

import (
	"slices"
)

func oper(a int, b int) int {
	return a*2 + b
}

func minOperations(nums []int, k int) int {
	slices.Sort(nums)
	if nums[0]*2 >= k {
		return 0
	}

	res := 0
	i, j := 0, len(nums)-1

	for i <= j {
		if nums[i]*2 >= k {
			break
		}
		if nums[i]+nums[j] >= k {
			j--
		} else {
			i++
		}
		res++
	}

	return res
}
