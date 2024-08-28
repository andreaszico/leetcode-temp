package main

import "fmt"

// Helper function to expand around the center and find the longest palindrome
func expandAroundCenter(s string, left int, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	fmt.Print(s[left+1:right], " ")
	return s[left+1 : right]
}

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	longest := ""

	for i := 0; i < len(s); i++ {
		// Check for odd length palindromes (single character center)
		oddPalindrome := expandAroundCenter(s, i, i)
		if len(oddPalindrome) > len(longest) {
			longest = oddPalindrome
		}

		// Check for even length palindromes (two character center)
		evenPalindrome := expandAroundCenter(s, i, i+1)
		if len(evenPalindrome) > len(longest) {
			longest = evenPalindrome
		}
	}

	return longest
}

func main() {
	// test := "babad"
	// test2 := "cbbd"
	test3 := "docnoteidissentafastneverpreventsafatnessidietoncod"
	test4 := "aaaaaaaa"
	test5 := "civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth"

	// fmt.Println(longestPalindrome(test))
	// fmt.Println(longestPalindrome(test2))
	fmt.Println(longestPalindrome(test3))
	fmt.Println()
	fmt.Println(longestPalindrome(test4))
	fmt.Println()
	fmt.Println(longestPalindrome(test5))
}
