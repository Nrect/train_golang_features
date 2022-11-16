package main

import (
	"fmt"
	"github.com/samber/lo"
	"strconv"
	"strings"
	"testProject/packages/calculateFromFormula"
)

func main() {
	testStr := "555 * 111 + 5"
	convertTestStr := strings.Split(testStr, " ")

	for _, character := range convertTestStr {
		if _, err := strconv.Atoi(character); err == nil {
			continue
		}
		validateCharacter(character)
	}

	repo := calculateFromFormula.NewMockDbResponse()
	result, err := calculateFromFormula.CalculateFromFormula(testStr, repo)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}

func validateCharacter(character string) {
	isValid := lo.Contains[string](calculateFromFormula.GetAvailableFormulaCharacters(), character)
	fmt.Println(isValid)
}
