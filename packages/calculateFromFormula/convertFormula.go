package calculateFromFormula

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"github.com/samber/lo"
	"reflect"
	"strconv"
	"strings"
)

func CalculateFromFormula(formula string, responseDb *FormulaResponseDb) (float64, error) {
	err := validateCharacter(formula)
	if err != nil {
		return 0, err
	}

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

// validateCharacter validate characters for formula
func validateCharacter(formula string) error {
	convertTestStr := strings.Split(formula, " ")
	isValid := false

	for _, character := range convertTestStr {
		if _, err := strconv.Atoi(character); err == nil {
			continue
		}
		isValid = lo.Contains[string](GetAvailableFormulaCharacters(), character)
	}

	if isValid != true {
		return fmt.Errorf(" Formula is not correct")
	}

	return nil
}
