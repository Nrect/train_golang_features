package lib

import (
	"github.com/Knetic/govaluate"
	"reflect"
	"testProject/src/constants"
	mockRepo "testProject/src/repo/mock"
)

func CalculateFromFormula(formula string, responseDb *mockRepo.FormulaResponseDb) (float64, error) {
	expression, err := govaluate.NewEvaluableExpression(formula)
	if err != nil {
		return 0, err
	}

	parameters, err := generateParameters(responseDb)
	if err != nil {
		return 0, err
	}

	result, err := expression.Evaluate(parameters)
	if err != nil {
		return 0, err
	}

	formulaSum := reflect.ValueOf(result).Float()
	return formulaSum, nil
}

// generateParameters generate static parameters for formula
func generateParameters(responseDb *mockRepo.FormulaResponseDb) (map[string]interface{}, error) {
	scheduleMoneyFieldNames := constants.GetScheduleMoneyFieldNames()
	parameters := make(map[string]interface{}, len(scheduleMoneyFieldNames))

	for _, field := range scheduleMoneyFieldNames {
		dbFieldValue, err := responseDb.GetMockDbResponseValue(field)
		parameters[field] = dbFieldValue

		if err != nil {
			return nil, err
		}
	}

	return parameters, nil
}
