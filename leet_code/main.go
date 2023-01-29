package main

import (
	"fmt"
	"github.com/pkg/profile"
	"leetCodeProj/easy"
)

func main() {
	defer profile.Start().Stop()

	testData := []int{1, 1, 1, 1, 1}
	result := easy.RunningSum(testData)
	fmt.Println(result)
}
