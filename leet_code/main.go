package main

import (
	"fmt"
	"leetCodeProj/easy"
)

func main() {
	testData := []int{1, 1, 1, 1, 1}
	result := easy.RunningSum(testData)
	fmt.Println(result)
}
