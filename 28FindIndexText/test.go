package main

func strStr(haystack string, needle string) int {
	for i := range haystack {
		if i+len(needle) > len(haystack) {
			return -1
		}

		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

func main() {
	println(strStr("sadbutsad", "sad"))
	println(strStr("leetcode", "leeto"))
	println(strStr("hello", "ll"))
	println(strStr("mississippi", "a"))
	println(strStr("a", "a"))
	println(strStr("abc", "c"))
}
