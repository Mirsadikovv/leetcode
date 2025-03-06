package main

import "fmt"

func checkPowersOfThree(n int) bool {
	count := 0
	for {
		if n%3 != 0 {

		}

		if toThreePower(count) == n {
			return true
		}

		count++
	}
}

func main() {
	fmt.Println(checkPowersOfThree(27))
}
func toThreePower(n int) int {
	res := 1
	for i := 0; i < n; i++ {
		res *= 3
	}
	return res
}
