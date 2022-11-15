package constants

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
