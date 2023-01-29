package easy

import (
	"fmt"
)

// https://leetcode.com/problems/fizz-buzz/

// my answer

func FizzBuzz(n int) []string {
	result := make([]string, n)

	for i := 1; i <= n; i++ {
		str := ""

		if i%3 == 0 {
			str += "Fizz"
		}
		if i%5 == 0 {
			str += "Buzz"
		}
		if str == "" {
			str = fmt.Sprintf("%d", i)
		}
		result[i-1] = str
	}

	return result
}
