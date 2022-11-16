package calculateFromFormula

import (
	"github.com/Knetic/govaluate"
	"reflect"
)

func CalculateFromFormula(formula string, responseDb *FormulaResponseDb) (float64, error) {
	expression, err := govaluate.NewEvaluableExpression(formula)
	if err != nil {
		return 0, err
	}

	parameters, err := GenerateParameters(responseDb)
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

// GenerateParameters generate static parameters for formula
func GenerateParameters(responseDb *FormulaResponseDb) (map[string]interface{}, error) {
	scheduleMoneyFieldNames := GetScheduleMoneyFieldNames()
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
