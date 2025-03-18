package main

import (
	"fmt"
)

func longestNiceSubarray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	maxLen := 1
	currOr := 0
	left := 0

	for right := 0; right < n; right++ {
		for currOr&nums[right] != 0 {
			currOr ^= nums[left]
			left++
		}
		currOr |= nums[right]
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}

	return maxLen
}

func main() {
	fmt.Println(longestNiceSubarray([]int{1, 3, 8, 48, 10})) // Ожидаемый результат: 3
	fmt.Println(longestNiceSubarray([]int{3, 1, 5, 11, 13})) // Ожидаемый результат: 1
}
