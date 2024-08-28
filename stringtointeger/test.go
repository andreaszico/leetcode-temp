package main

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

func myAtoi(s string) int {
	// Trim leading whitespaces
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}

	// Initialize variables
	sign := 1
	startIndex := 0
	if s[0] == '-' {
		sign = -1
		startIndex = 1
	} else if s[0] == '+' {
		startIndex = 1
	}

	result := 0
	for i := startIndex; i < len(s); i++ {
		if !unicode.IsDigit(rune(s[i])) {
			break
		}
		digit := int(s[i] - '0')

		// Check for overflow and underflow
		if result > (math.MaxInt32-digit)/10 {
			if sign == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}

		result = result*10 + digit
	}

	return result * sign
}

func main() {
	fmt.Println(myAtoi("42"))              // Output: 42
	fmt.Println(myAtoi("   -42"))          // Output: -42
	fmt.Println(myAtoi("4193 with words")) // Output: 4193
	fmt.Println(myAtoi("words and 987"))   // Output: 0
	fmt.Println(myAtoi("-91283472332"))    // Output: -2147483648
}
