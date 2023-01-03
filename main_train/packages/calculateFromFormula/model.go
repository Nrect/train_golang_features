package calculateFromFormula

import "testProject/packages/calculateFromFormula/utils"

/*
		DetailsIncome - общий накат (бэк)
		CommonTenantsIncome - общий доход арендатора(на фронте)
		EarnAmount - чистый доход арендатора (бэк)
		GiveOutAmount - к выплате арендатору (на фронте)
		CompanyProfit - прибыль компании (бэк)

		DetailsCash - наличные, без чаевых
		DetailsCard - безналичные, без чаевых
		DetailsTip - безналичные чаевые
		ServiceCommission -  комиссия агрегатора
		ParkCommission - комиссия партнера
		RetentionSum -  удержания
		SurchargeSum -  доплаты
	    FuelSum -  топливо
	    DriverPercent - % водителя
	    CompanyPercent - % компании
*/
type ScheduleMoneyField string
type CacheKey string

// TODO: write test for get keys from cache
const (
	formulaDetailsIncome       CacheKey = "formula_details_income"
	formulaCommonTenantsIncome CacheKey = "formula_common_tenants_income"
	formulaEarnAmount          CacheKey = "formula_earn_amount"
	formulaGiveOutAmount       CacheKey = "formula_give_out_amount"
	formulaCompanyProfit       CacheKey = "formula_company_profit"
)

const (
	DetailsIncome       ScheduleMoneyField = "DetailsIncome"
	CommonTenantsIncome ScheduleMoneyField = "CommonTenantsIncome"
	EarnAmount          ScheduleMoneyField = "EarnAmount"
	GiveOutAmount       ScheduleMoneyField = "GiveOutAmount"
	CompanyProfit       ScheduleMoneyField = "CompanyProfit"

	DetailsCash       ScheduleMoneyField = "DetailsCash"
	DetailsCard       ScheduleMoneyField = "DetailsCard"
	DetailsTip        ScheduleMoneyField = "DetailsTip"
	ServiceCommission ScheduleMoneyField = "ServiceCommission"
	ParkCommission    ScheduleMoneyField = "ParkCommission"
	RetentionSum      ScheduleMoneyField = "RetentionSum"
	SurchargeSum      ScheduleMoneyField = "SurchargeSum"
	FuelSum           ScheduleMoneyField = "FuelSum"
	DriverPercent     ScheduleMoneyField = "DriverPercent"
	CompanyPercent    ScheduleMoneyField = "CompanyPercent"
)

func (s ScheduleMoneyField) ToString() string {
	return string(s)
}

type FormulaResponseDb struct {
	DetailsIncome       float64
	CommonTenantsIncome float64
	EarnAmount          float64
	GiveOutAmount       float64
	CompanyProfit       float64
	DetailsCash         float64
	DetailsCard         float64
	DetailsTip          float64
	ServiceCommission   float64
	ParkCommission      float64
	RetentionSum        float64
	SurchargeSum        float64
	FuelSum             float64
	DriverPercent       float64
	CompanyPercent      float64
}

func NewMockDbResponse() *FormulaResponseDb {
	return &FormulaResponseDb{
		DetailsIncome:       utils.GetRandomMoneyValue(),
		CommonTenantsIncome: utils.GetRandomMoneyValue(),
		EarnAmount:          utils.GetRandomMoneyValue(),
		GiveOutAmount:       utils.GetRandomMoneyValue(),
		CompanyProfit:       utils.GetRandomMoneyValue(),
		DetailsCash:         utils.GetRandomMoneyValue(),
		DetailsCard:         utils.GetRandomMoneyValue(),
		DetailsTip:          utils.GetRandomMoneyValue(),
		ServiceCommission:   utils.GetRandomMoneyValue(),
		ParkCommission:      utils.GetRandomMoneyValue(),
		RetentionSum:        utils.GetRandomMoneyValue(),
		SurchargeSum:        utils.GetRandomMoneyValue(),
		FuelSum:             utils.GetRandomMoneyValue(),
		DriverPercent:       40,
	}
}

func (r *FormulaResponseDb) GetMockDbResponseValue(keyName string) (float64, error) {
	value, err := utils.GetPropertyStruct(r, keyName)
	if err != nil {
		return 0, err
	}
	return value.Float(), nil
}

func GetScheduleMoneyFieldNames() []string {
	scheduleMoneyFieldNames := []string{
		DetailsIncome.ToString(),
		CommonTenantsIncome.ToString(),
		EarnAmount.ToString(),
		GiveOutAmount.ToString(),
		CompanyProfit.ToString(),
		DetailsCash.ToString(),
		DetailsCard.ToString(),
		DetailsTip.ToString(),
		ServiceCommission.ToString(),
		ParkCommission.ToString(),
		RetentionSum.ToString(),
		SurchargeSum.ToString(),
		FuelSum.ToString(),
		DriverPercent.ToString(),
		CompanyPercent.ToString(),
	}
	return scheduleMoneyFieldNames
}

func GetAvailableFormulaCharacters() []string {
	availableFormulaCharacters := []string{
		"+",
		"-",
		"*",
		"/",
		"(",
		")",
	}
	availableFormulaCharacters = append(availableFormulaCharacters, GetScheduleMoneyFieldNames()...)

	return availableFormulaCharacters
}
