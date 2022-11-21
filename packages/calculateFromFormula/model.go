package calculateFromFormula

import "testProject/packages/calculateFromFormula/utils"

type FormulaResponseDb struct {
	CashIncome        float64
	DetailsTip        float64
	DetailsCash       float64
	DetailsCard       float64
	DetailsIncome     float64
	DetailsNetCash    float64
	DetailsNetCard    float64
	ParkCommission    float64
	ServiceCommission float64
	DriverPercent     float64
	EarnAmount        float64
	CompanyProfit     float64
	Tip               float64
}

func NewMockDbResponse() *FormulaResponseDb {
	return &FormulaResponseDb{
		DetailsTip:        utils.GetRandomMoneyValue(),
		CashIncome:        utils.GetRandomMoneyValue(),
		DetailsCash:       utils.GetRandomMoneyValue(),
		DetailsCard:       utils.GetRandomMoneyValue(),
		DetailsIncome:     utils.GetRandomMoneyValue(),
		DetailsNetCash:    utils.GetRandomMoneyValue(),
		DetailsNetCard:    utils.GetRandomMoneyValue(),
		ParkCommission:    utils.GetRandomMoneyValue(),
		ServiceCommission: utils.GetRandomMoneyValue(),
		DriverPercent:     40,
		EarnAmount:        utils.GetRandomMoneyValue(),
		CompanyProfit:     utils.GetRandomMoneyValue(),
	}
}

func (r *FormulaResponseDb) GetMockDbResponseValue(keyName string) (float64, error) {
	value, err := utils.GetPropertyStruct(r, keyName)
	if err != nil {
		return 0, err
	}
	return value.Float(), nil
}
