// Palindrome check
package pgms

import "fmt"

func palindromeV1(s string) {
	left := 0
	right := len(s) - 1

	for left < right {
		if s[left] != s[right] {
			fmt.Println("Not palindrome")
			return
		}
		left++
		right--
	}
	fmt.Println("Palindrome")
}

func palindromeV2(s string) {
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-i-1] {
			fmt.Println("Not palindrome")
			return
		}
	}
	fmt.Println("Palindrome")
}

func Program003() {
	s := "abba"

	palindromeV2(s)
}
