package main

import (
	"fmt"
	"github.com/samber/lo"
	"testProject/src/constants"
	"testProject/src/lib"
	mockRepo "testProject/src/repo/mock"
)

func main() {
	ll := lo.Contains[string](constants.GetScheduleMoneyFieldNames(), "DetailsTip")

	repo := mockRepo.NewMockDbResponse()
	result, err := lib.CalculateFromFormula("5 + 5", repo)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("result:", result)
}
