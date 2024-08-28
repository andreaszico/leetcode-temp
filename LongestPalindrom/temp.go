package main

// import "fmt"

// func isPalindroms(s string) bool {

// 	temp := ""

// 	for i := len(s) - 1; i >= 0; i-- {
// 		temp = fmt.Sprintf("%s%c", temp, s[i])
// 	}

// 	return temp == s
// }

// func longestPalindrome(s string) string {

// 	longest := ""

// 	if len(s) == 1 {
// 		return s
// 	}

// 	for i := 0; i < len(s); i++ {
// 		// fmt.Printf("%c\n", s[i])
// 		temp := fmt.Sprintf("%c", s[i])
// 		for j := i + 1; j < len(s); j++ {
// 			temp = fmt.Sprintf("%s%c", temp, s[j])
// 			// fmt.Printf("%c => %c | %s | %t\n", s[i], s[j], temp, isPalindrom(temp))
// 			if isPalindroms(temp) {
// 				// println(len(longest), len(temp))
// 				if len(longest) < len(temp) {
// 					longest = temp
// 				}
// 			}
// 			// println()
// 		}
// 		// println("====================")
// 	}

// 	println("longest", longest)

// 	return ""
// }

// func main() {
// 	test := "babad"
// 	test2 := "cbbd"
// 	test3 := "a"
// 	test4 := "ac"
// 	test5 := "civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth"

// 	fmt.Println(longestPalindrome(test))
// 	fmt.Println(longestPalindrome(test2))
// 	fmt.Println(longestPalindrome(test3))
// 	fmt.Println(longestPalindrome(test4))
// 	fmt.Println(longestPalindrome(test5))
// }
