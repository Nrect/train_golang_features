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

func GetScheduleMoneyFieldNames() []ScheduleMoneyField {
	scheduleMoneyFieldNames := []ScheduleMoneyField{
		DetailsTip,
		DetailsCash,
		DetailsCard,
		DetailsIncome,
		DetailsNetCash,
		DetailsNetCard,
		ParkCommission,
		ServiceCommission,
		DriverPercent,
		EarnAmount,
		CompanyProfit,
	}
	return scheduleMoneyFieldNames
}

func (s ScheduleMoneyField) ToString() string {
	return string(s)
}
