package main

import (
	"fmt"
	"testProject/packages/calculateFromFormula"
)

func main() {
	detailsIncome := "CashIncome + DetailsCard + Tip"

	repo := calculateFromFormula.NewMockDbResponse()
	result, err := calculateFromFormula.CalculateFromFormula(detailsIncome, repo)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func testRequestCase() {
	/*
		details_income  общий накат
		cash_income нал, без чаевых
		details_card безнал, без чаевых
		tip безнал чаевые
	*/
}
