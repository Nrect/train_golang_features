package lib

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"reflect"
	"testProject/src/utils"
)

type ScheduleMoneyField string

const (
	detailsTip        ScheduleMoneyField = "detailsTip"
	detailsCash       ScheduleMoneyField = "detailsCash"
	detailsCard       ScheduleMoneyField = "detailsCard"
	detailsIncome     ScheduleMoneyField = "detailsIncome"
	detailsNetCash    ScheduleMoneyField = "detailsNetCash"
	detailsNetCard    ScheduleMoneyField = "detailsNetCard"
	parkCommission    ScheduleMoneyField = "parkCommission"
	serviceCommission ScheduleMoneyField = "serviceCommission"
	driverPercent     ScheduleMoneyField = "driverPercent"
	earnAmount        ScheduleMoneyField = "earnAmount"
	companyProfit     ScheduleMoneyField = "companyProfit"
)

func getScheduleMoneyFieldNames() []ScheduleMoneyField {
	scheduleMoneyFieldNames := []ScheduleMoneyField{
		detailsTip,
		detailsCash,
		detailsCard,
		detailsIncome,
		detailsNetCash,
		detailsNetCard,
		parkCommission,
		serviceCommission,
		driverPercent,
		earnAmount,
		companyProfit,
	}
	return scheduleMoneyFieldNames
}

type responseDb struct {
	detailsTip        float64
	detailsCash       float64
	detailsCard       float64
	detailsIncome     float64
	detailsNetCash    float64
	detailsNetCard    float64
	parkCommission    float64
	serviceCommission float64
	driverPercent     float64
	earnAmount        float64
	companyProfit     float64
}

func getMockValueFromRepo(keyName string) float64 {
	randomDbValues := &responseDb{
		detailsTip:        utils.GetRandomMoneyValue(),
		detailsCash:       utils.GetRandomMoneyValue(),
		detailsCard:       utils.GetRandomMoneyValue(),
		detailsIncome:     utils.GetRandomMoneyValue(),
		detailsNetCash:    utils.GetRandomMoneyValue(),
		detailsNetCard:    utils.GetRandomMoneyValue(),
		parkCommission:    utils.GetRandomMoneyValue(),
		serviceCommission: utils.GetRandomMoneyValue(),
		driverPercent:     utils.GetRandomMoneyValue(),
		earnAmount:        utils.GetRandomMoneyValue(),
		companyProfit:     utils.GetRandomMoneyValue(),
	}

	value := utils.GetPropertyStruct(randomDbValues, keyName)
	return value.Float()
}

func CalculateFromFormula(formula string) (float64, error) {
	expression, err := govaluate.NewEvaluableExpression(formula)
	if err != nil {
		return 0, err
	}

	scheduleMoneyFieldNames := getScheduleMoneyFieldNames()

	parameters := make(map[string]interface{}, len(scheduleMoneyFieldNames))

	for _, field := range scheduleMoneyFieldNames {
		parameters[string(field)] = getMockValueFromRepo(string(field))
	}

	fmt.Println(parameters)

	result, err := expression.Evaluate(parameters)
	if err != nil {
		return 0, err
	}

	x := reflect.ValueOf(result).Float()

	return x, nil
}
