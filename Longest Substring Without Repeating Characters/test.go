package main

import "fmt"

func removeDuplicates(input string) string {

	if len(input) == 1 {
		return input
	}

	result := string(input[0])

	for i := 1; i < len(input); i++ {
		if input[i] != input[i-1] {
			result += string(input[i])
		}
	}

	return result
}

func lengthOfLongestSubstring(s string) int {
	output := removeDuplicates(s)
	fmt.Println("After removing duplicates:", output)
	// for i := 0; i < len(s); i++ {
	// 	l := len(s)
	// 	fmt.Printf("%s", s[i:l])
	// 	fmt.Println()
	// }
	// fmt.Println()
	return len(s)
}

func main() {
	str := []string{"abcabcbb", "bbbbb", "pwwkew"}

	for _, v := range str {
		lengthOfLongestSubstring(v)
	}
}
