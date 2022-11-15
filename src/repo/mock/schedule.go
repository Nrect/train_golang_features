package mockRepo

import "testProject/src/utils"

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

func NewMockDbResponse() *responseDb {
	return &responseDb{
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
}

func (r *responseDb) GetMockDbResponseValue(keyName string) float64 {
	value := utils.GetPropertyStruct(r, keyName)
	return value.Float()
}
