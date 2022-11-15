package main

import (
	"fmt"
	"testProject/src/lib"
	mockRepo "testProject/src/repo/mock"
)

func main() {
	repo := mockRepo.NewMockDbResponse()
	result, err := lib.CalculateFromFormula("5 + 5", repo)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("result:", result)
}
