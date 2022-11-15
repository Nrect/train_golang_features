package constants

type ScheduleMoneyField string

const (
	detailsTip        ScheduleMoneyField = "DetailsTip"
	detailsCash       ScheduleMoneyField = "DetailsCash"
	detailsCard       ScheduleMoneyField = "DetailsCard"
	detailsIncome     ScheduleMoneyField = "DetailsIncome"
	detailsNetCash    ScheduleMoneyField = "DetailsNetCash"
	detailsNetCard    ScheduleMoneyField = "DetailsNetCard"
	parkCommission    ScheduleMoneyField = "ParkCommission"
	serviceCommission ScheduleMoneyField = "ServiceCommission"
	driverPercent     ScheduleMoneyField = "DriverPercent"
	earnAmount        ScheduleMoneyField = "EarnAmount"
	companyProfit     ScheduleMoneyField = "CompanyProfit"
)

func GetScheduleMoneyFieldNames() []ScheduleMoneyField {
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
