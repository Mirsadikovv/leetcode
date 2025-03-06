package main

import "fmt"

// func finalPrices(prices []int) []int {
// 	for i := 0; i < len(prices)-1; i++ {
// 		for j := i + 1; j < len(prices); j++ {
// 			if prices[i] >= prices[j] {
// 				prices[i] -= prices[j]
// 				break
// 			}
// 		}
// 	}
// 	return prices
// }

func finalPrices(prices []int) []int {
	index := 0
	for i := 1; i < len(prices); i++ {

		if prices[index] > prices[i] {
			prices[index] -= prices[i]
			index++

			break
		}

	}
	return prices
}

func main() {
	fmt.Println(finalPrices([]int{8, 4, 6, 2, 3}))
}
