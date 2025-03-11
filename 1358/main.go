package main

import (
	"fmt"
)

func main() {
	fmt.Println(numberOfSubstrings("aaacb"))
}
func numberOfSubstrings(s string) int {
	n := len(s)
	if n < 3 {
		return 0
	}

	count := 0
	charCount := [3]int{0, 0, 0}

	left := 0
	for right := 0; right < n; right++ {
		switch s[right] {
		case 'a':
			charCount[0]++
		case 'b':
			charCount[1]++
		case 'c':
			charCount[2]++
		}

		for charCount[0] > 0 && charCount[1] > 0 && charCount[2] > 0 {

			count += n - right

			switch s[left] {
			case 'a':
				charCount[0]--
			case 'b':
				charCount[1]--
			case 'c':
				charCount[2]--
			}
			left++
		}
	}

	return count
}
