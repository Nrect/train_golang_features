package lib

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"reflect"
	"testProject/src/constants"
	mockRepo "testProject/src/repo/mock"
)

func CalculateFromFormula(formula string, repo *mockRepo.ResponseDb) (float64, error) {
	expression, err := govaluate.NewEvaluableExpression(formula)
	if err != nil {
		return 0, err
	}

	scheduleMoneyFieldNames := constants.GetScheduleMoneyFieldNames()
	parameters := make(map[string]interface{}, len(scheduleMoneyFieldNames))

	for _, field := range scheduleMoneyFieldNames {
		parameters[string(field)], err = repo.GetMockDbResponseValue(string(field))

		if err != nil {
			return 0, err
		}
	}

	result, err := expression.Evaluate(parameters)
	if err != nil {
		return 0, err
	}

	fmt.Println(parameters)

	formulaSum := reflect.ValueOf(result).Float()
	return formulaSum, nil
}
