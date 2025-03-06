package main

import "fmt"

func findTargetSumWays(nums []int, target int) int {
	tar, count := 0, 0
	for i := 0; i < len(nums); i++ {

		fmt.Println(nums)

		if tar == target {
			count++
		}
		tar = 0

	}
	return count
}
