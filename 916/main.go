package main

import "fmt"

func wordSubsets(words1 []string, words2 []string) []string {

	strings := make(map[string]int, len(words2))

	for i := 0; i < len(words1); i++ {

		for j := 0; j < len(words2); j++ {

			if IsSubstring(words1[i], words2[j]) {

				strings[words1[i]]++
			}
		}
	}

	ans := make([]string, 0, len(words1))

	for k, v := range strings {
		if v == len(words2) {
			ans = append(ans, k)
		}
	}

	return ans
}

func IsSubstring(str1, str2 string) bool {

	l1, l2 := len(str1), len(str2)

	if l2 == 0 {
		return true
	}

	if l2 > l1 {
		return false
	}

	for i := 0; i <= l1-l2; i++ {

		if str1[i:i+l2] == str2 {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(wordSubsets([]string{"lo", "lo"}, []string{"leetcode", "cods"}))
}
