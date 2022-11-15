package mockRepo

import "testProject/src/utils"

type ResponseDb struct {
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
}

func NewMockDbResponse() *ResponseDb {
	return &ResponseDb{
		DetailsTip:        utils.GetRandomMoneyValue(),
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

func (r *ResponseDb) GetMockDbResponseValue(keyName string) (float64, error) {
	value, err := utils.GetPropertyStruct(r, keyName)
	if err != nil {
		return 0, err
	}
	return value.Float(), nil
}
