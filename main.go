package main

import (
	"fmt"
	"testProject/packages/calculateFromFormula"
)

func main() {
	testStr := "555 * 111 + 5"

	repo := calculateFromFormula.NewMockDbResponse()
	result, err := calculateFromFormula.CalculateFromFormula(testStr, repo)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)

	}
}
