package lib

import (
	"github.com/Knetic/govaluate"
	"reflect"
)

func CalculateFromFormula(formula string) (float64, error) {
	expression, err := govaluate.NewEvaluableExpression(formula)
	if err != nil {
		return 0, err
	}

	result, err := expression.Evaluate(nil)
	if err != nil {
		return 0, err
	}

	x := reflect.ValueOf(result).Float()

	return x, nil
}
