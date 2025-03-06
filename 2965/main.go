package main

import (
	"fmt"
)

func main() {

	fmt.Println(findMissingAndRepeatedValues([][]int{{1, 2}, {1, 4}}))
}

func findMissingAndRepeatedValues(grid [][]int) []int {
	var res []int

	seen := make(map[int]bool)
	for _, row := range grid {
		for _, val := range row {
			if seen[val] {
				res = append(res, val)
			}
			seen[val] = true
		}
	}
	var c int = len(seen) + 1
	for i := 1; i <= len(seen); i++ {
		if seen[i] == false {
			c = i
		}
	}
	if c != 0 {
		res = append(res, c)
	}
	return res
}
