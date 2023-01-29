package main

import (
	"fmt"
	"leetCodeProj/easy"
)

func main() {
	testData := [][]int{{1, 2, 30}, {3, 2, 1}, {}}
	result := easy.MaximumWealth(testData)
	fmt.Println(result)
}
