package main

import (
	"fmt"
	"math"
	"strconv"
)

func reverse(x int) int {
	// Determine if the number is negative
	isNegative := x < 0

	// Convert the absolute value of the integer to a string
	str := strconv.Itoa(int(math.Abs(float64(x))))

	// Reverse the string
	reversedStr := ""
	for _, v := range str {
		reversedStr = string(v) + reversedStr
	}

	// Convert the reversed string back to an integer
	result, err := strconv.Atoi(reversedStr)
	if err != nil {
		panic(err)
	}

	// If the original number was negative, make the result negative
	if isNegative {
		result = -result
	}

	// Handle integer overflow cases
	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}

	return result
}

func main() {
	fmt.Println(reverse(123))  // Output: 321
	fmt.Println(reverse(-123)) // Output: -321
	fmt.Println(reverse(120))  // Output: 21
	fmt.Println(reverse(0))    // Output: 0
}
