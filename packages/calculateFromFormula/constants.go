package calculateFromFormula

const (
	CashIncome    ScheduleMoneyField = "CashIncome"
	DetailsCard   ScheduleMoneyField = "DetailsCard"
	DetailsIncome ScheduleMoneyField = "DetailsIncome"
	Tip           ScheduleMoneyField = "Tip"
)

type ScheduleMoneyField string

func (s ScheduleMoneyField) ToString() string {
	return string(s)
}

func GetScheduleMoneyFieldNames() []string {
	scheduleMoneyFieldNames := []string{
		CashIncome.ToString(),
		DetailsCard.ToString(),
		DetailsIncome.ToString(),
		Tip.ToString(),
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
