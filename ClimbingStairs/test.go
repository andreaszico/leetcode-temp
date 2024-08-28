package main

import "fmt"

func climbStairs(n int) int {
	memo := make(map[int]int)
	return climbStairsHelper(n, memo)
}

func climbStairsHelper(n int, memo map[int]int) int {
	if n == 0 || n == 1 {
		return 1
	}
	if val, found := memo[n]; found {
		return val
	}
	memo[n] = climbStairsHelper(n-1, memo) + climbStairsHelper(n-2, memo)
	return memo[n]
}

func main() {
	fmt.Println(climbStairs(44))
}
