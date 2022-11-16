package main

import (
	"fmt"
	"github.com/samber/lo"
	"strconv"
	"strings"
	"testProject/src/constants"
	"testProject/src/lib"
	mockRepo "testProject/src/repo/mock"
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

	repo := mockRepo.NewMockDbResponse()
	_, err := lib.CalculateFromFormula(testStr, repo)
	if err != nil {
		fmt.Println(err)
	}
}

func validateCharacter(character string) {
	isValid := lo.Contains[string](constants.GetAvailableFormulaCharacters(), character)
	fmt.Println(isValid)
}
