package main

import (
	"fmt"
	"testProject/src/lib"
)

func main() {
	result, err := lib.CalculateFromFormula("5 + 5")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("result", result)
}
