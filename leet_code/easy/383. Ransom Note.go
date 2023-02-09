package easy

import "strings"

//https://leetcode.com/problems/ransom-note/

// CanConstruct My solution
func CanConstruct(ransomNote string, magazine string) bool {
	magazineLetters := make(map[int32]uint32, len(magazine))

	for _, char := range magazine {
		magazineLetters[char] += 1
	}

	for _, char := range ransomNote {

		if magazineLetters[char] == 0 {
			return false
		}

		magazineLetters[char]--
	}

	return true
}

// other interest solution
func canConstruct(ransomNote string, magazine string) bool {
	for _, v := range ransomNote {
		if strings.Count(ransomNote, string(v)) > strings.Count(magazine, string(v)) {
			return false
		}
	}
	return true
}
