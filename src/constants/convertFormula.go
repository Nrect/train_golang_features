package constants

type ScheduleMoneyField string

const (
	DetailsTip        ScheduleMoneyField = "DetailsTip"
	DetailsCash       ScheduleMoneyField = "DetailsCash"
	DetailsCard       ScheduleMoneyField = "DetailsCard"
	DetailsIncome     ScheduleMoneyField = "DetailsIncome"
	DetailsNetCash    ScheduleMoneyField = "DetailsNetCash"
	DetailsNetCard    ScheduleMoneyField = "DetailsNetCard"
	ParkCommission    ScheduleMoneyField = "ParkCommission"
	ServiceCommission ScheduleMoneyField = "ServiceCommission"
	DriverPercent     ScheduleMoneyField = "DriverPercent"
	EarnAmount        ScheduleMoneyField = "EarnAmount"
	CompanyProfit     ScheduleMoneyField = "CompanyProfit"
)

func GetScheduleMoneyFieldNames() []string {
	scheduleMoneyFieldNames := []string{
		DetailsTip.ToString(),
		DetailsCash.ToString(),
		DetailsCard.ToString(),
		DetailsIncome.ToString(),
		DetailsNetCash.ToString(),
		DetailsNetCard.ToString(),
		ParkCommission.ToString(),
		ServiceCommission.ToString(),
		DriverPercent.ToString(),
		EarnAmount.ToString(),
	}
	return scheduleMoneyFieldNames
}

func (s ScheduleMoneyField) ToString() string {
	return string(s)
}
