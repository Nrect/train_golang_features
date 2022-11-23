package main

import (
	"fmt"
	"testProject/packages/calculateFromFormula"
)

func main() {
	detailsIncome := "DetailsIncome + CommonTenantsIncome + EarnAmount"

	repo := calculateFromFormula.NewMockDbResponse()
	result, err := calculateFromFormula.CalculateFromFormula(detailsIncome, repo)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
