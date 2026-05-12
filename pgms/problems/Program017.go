// Is subsequence
package pgms

import "fmt"

func isSubsequenceV1(s1, s2 string) bool {
	count := 0
	for _, char := range s1 {
		for idx, char2 := range s2 {
			if char == char2 {
				s2 = s2[idx+1:]
				count++
				break
			}
		}
	}
	return count == len(s1)	
}

func isSubsequenceV2(s1, s2 string) bool {
	i := 0

	for _, char := range s2 {
		if i < len(s1) && char == rune(s1[i]) {
			i++
		}
	}

	return i == len(s1)
}

func Program017() {
	s1 := "AXY"
	s2 := "ADXCPY"

	fmt.Println(isSubsequenceV2(s1, s2))
}