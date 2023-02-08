package main

import (
	"fmt"
	"leetCodeProj/easy"
)

func main() {
	example := make([]easy.ListNode, 10)

	for i := 0; i < 10; i++ {
		example[i] = easy.ListNode{Val: i}

		if i > 0 {
			example[i-1].Next = &example[i]
		}
	}

	result := make([]easy.ListNode, 0)
	for _, head := range example {
		result = append(result, *easy.MiddleNode(&head))
	}

	fmt.Println(result)
}
