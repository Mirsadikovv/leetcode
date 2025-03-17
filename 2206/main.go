package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(divideArray([]int{1, 2, 3, 3, 4, 1}))
}

//with bitwice
// func divideArray(nums []int) bool {
// 	p, cnt := make([]int, 501), 0
// 	for _, n := range nums {
// 		p[n]++
// 		cnt += (p[n]&1)*2 - 1
// 	}
// 	return cnt == 0
// }

func divideArray(nums []int) bool {

	var seen = make(map[int]bool, len(nums))
	var arr []int
	for _, e := range nums {
		if _, ok := seen[e]; !ok {
			seen[e] = false
			arr = append(arr, e)
		} else if !seen[e] {
			seen[e] = true
			arr = slices.DeleteFunc(arr, func(v int) bool {
				return v == e
			})
		} else if seen[e] {
			seen[e] = false
			arr = append(arr, e)
		}
	}

	if len(arr) == 0 {
		return true
	}

	return false
}
