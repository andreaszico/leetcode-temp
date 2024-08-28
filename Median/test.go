package main

import (
	"math"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	merged := append(nums1, nums2...)

	sort.Slice(merged, func(i, j int) bool {
		return merged[i] < merged[j]
	})
	length := len(merged)
	result := length % 2

	middle := float64(merged[int(math.Ceil(float64(length/2)))])
	if result == 1 {
		return middle
	}

	return float64(middle+float64(merged[int(math.Ceil(float64(length/2)))-1])) / 2
}

func main() {

	arr1 := []int{1, 2, 5}
	arr2 := []int{4, 3, 6, 10}
	println(findMedianSortedArrays(arr1, arr2))

}
