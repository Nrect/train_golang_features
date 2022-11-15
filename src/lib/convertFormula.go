package lib

import (
	"github.com/Knetic/govaluate"
	"reflect"
	"testProject/src/constants"
	mockRepo "testProject/src/repo/mock"
)

func CalculateFromFormula(formula string) (float64, error) {
	expression, err := govaluate.NewEvaluableExpression(formula)
	if err != nil {
		return 0, err
	}

	scheduleMoneyFieldNames := constants.GetScheduleMoneyFieldNames()
	parameters := make(map[string]interface{}, len(scheduleMoneyFieldNames))

	repo := mockRepo.NewMockDbResponse()

	for _, field := range scheduleMoneyFieldNames {
		parameters[string(field)] = repo.GetMockDbResponseValue(string(field))
	}

	result, err := expression.Evaluate(parameters)
	if err != nil {
		return 0, err
	}

	formulaSum := reflect.ValueOf(result).Float()
	return formulaSum, nil
}
