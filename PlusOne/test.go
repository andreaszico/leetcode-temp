package main

func plusOne(digits []int) []int {

	last := len(digits) - 1

	for last != -2 {
		if last < 0 {
			digits = append([]int{1}, digits...)
			return digits
		}

		sum := digits[last] + 1

		if sum > 9 {
			digits[last] = 0
			last--
			continue
		}

		digits[last] = sum

		return digits
	}

	return digits
}

func printArr(nums []int) {
	for _, v := range nums {
		print(v)
	}
	println()
}

func main() {
	printArr(plusOne([]int{1, 2, 3}))
	printArr(plusOne([]int{9}))
	printArr(plusOne([]int{9, 9}))
	printArr(plusOne([]int{1, 9, 9, 9}))
}
