package main

import "strings"

func clearDigits(s string) string {
	res := ""

	for _, c := range s {
		if IsDigit(string(c)) {
			res = res[:len(res)-1]
		} else {
			res += string(c)
		}
	}

	return res
}

func IsDigit(s string) bool {
	return len(s) > 0 && s[0] >= '0' && s[0] <= '9'
}

func removeOccurrences(s string, part string) string {
	s = Reverse(s)
	part = Reverse(part)
	res := strings.ReplaceAll(s, part, "")
	for {
		if res == strings.ReplaceAll(res, part, "") {
			break
		} else {
			res = strings.ReplaceAll(res, part, "")
		}
	}

	return Reverse(res)
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
