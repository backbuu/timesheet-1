package timesheet

import (
	"timesheet/internal/model"
)

const (
	SiamChamnankitCompany = "siam_chamnankit"
	ShuhariCompany        = "shuhari"

	NoSalary         = 0.00
	NoIncomeTax1     = 0.00
	NoSocialSecurity = 0.00

	OneMinute = 60
	OneHour   = 60
)

type TimesheetGateways interface {
	CalculatePaymentSummary(member []model.Member, incomes []model.Incomes, year, month int) []model.TransactionTimesheet
	CalculatePayment(incomes []model.Incomes) model.Payment
}

type Timesheet struct {
	TransactionTimesheet model.TransactionTimesheet
	Payment              model.Payment
}

func (timesheet Timesheet) CalculatePaymentSummary(member []model.Member, incomes []model.Incomes, year, month int) []model.TransactionTimesheet {
	var transactionTimesheetList []model.TransactionTimesheet
	for _, member := range member {
		totalCoachingPaymentRate := CalculateTotalCoachingPaymentRate(incomes, member.Company)
		totalTrainingWage := CalculateTotalTrainingWage(incomes, member.Company)
		totalOtherWage := CalculateTotalOtherWage(incomes, member.Company, member.TravelExpense)
		paymentWage := CalculateTotalPaymentWage(totalCoachingPaymentRate, totalTrainingWage, totalOtherWage)
		salary, incomeTax1, socialSecurity, netSalary := CalculateNetSalary(paymentWage, member.Salary, member.IncomeTax1, member.SocialSecurity)
		wage := CalculateWage(paymentWage, member.Salary)
		incomeTax53 := CalculateIncomeTax53(wage, member.IncomeTax53Percentage)
		netWage := CalculateNetWage(member.IncomeTax53Percentage, paymentWage, member.Salary)
		netTransfer := CalculateNetTransfer(netSalary, netWage)
		transactionTimesheet := model.TransactionTimesheet{
			MemberID:              member.MemberID,
			Month:                 month,
			Year:                  year,
			MemberNameTH:          member.MemberNameTH,
			Company:               member.Company,
			Coaching:              totalCoachingPaymentRate,
			Training:              totalTrainingWage,
			Other:                 totalOtherWage,
			TotalIncomes:          paymentWage,
			Salary:                salary,
			IncomeTax1:            incomeTax1,
			SocialSecurity:        socialSecurity,
			NetSalary:             netSalary,
			Wage:                  wage,
			IncomeTax53Percentage: member.IncomeTax53Percentage,
			IncomeTax53:           incomeTax53,
			NetWage:               netWage,
			NetTransfer:           netTransfer,
		}
		if netTransfer > 0 {
			transactionTimesheetList = append(transactionTimesheetList, transactionTimesheet)
		}
	}
	return transactionTimesheetList
}

func (timesheet Timesheet) CalculatePayment(incomes []model.Incomes) model.Payment {
	totalHour := CalculateTotalHour(incomes)
	totalCoachingCustomerCharging := CalculateTotalCoachingCustomerCharging(incomes)
	totalCoachingPaymentRate := CalculateTotalCoachingPaymentRate(incomes, "")
	totalTrainingWage := CalculateTotalTrainingWage(incomes, "")
	totalOtherWage := CalculateTotalOtherWage(incomes, "", 0.00)
	paymentWage := CalculateTotalPaymentWage(totalCoachingPaymentRate, totalTrainingWage, totalOtherWage)
	return model.Payment{
		TotalHoursHours:               totalHour.Hours,
		TotalHoursMinutes:             totalHour.Minutes,
		TotalHoursSeconds:             totalHour.Seconds,
		TotalCoachingCustomerCharging: totalCoachingCustomerCharging,
		TotalCoachingPaymentRate:      totalCoachingPaymentRate,
		TotalTrainigWage:              totalTrainingWage,
		TotalOtherWage:                totalOtherWage,
		PaymentWage:                   paymentWage,
	}
}

func CalculateTotalHour(incomes []model.Incomes) model.Time {
	var hours int
	for _, income := range incomes {
		hours += int(income.EndTimeAM.Sub(income.StartTimeAM))
		hours += int(income.EndTimePM.Sub(income.StartTimePM))
		hours += income.Overtime
	}

	return model.Time{
		Hours: hours,
	}
}

func CalculateTotalPaymentWage(coachingPaymentRate, trainingWage, otherWage float64) float64 {
	return coachingPaymentRate + trainingWage + otherWage
}

func CalculateNetSalary(paymentWage, salary, incomeTax1, socialSecurity float64) (float64, float64, float64, float64) {
	netSalary := salary - incomeTax1 - socialSecurity
	if paymentWage >= salary {
		return salary, incomeTax1, socialSecurity, netSalary
	}
	return NoSalary, NoIncomeTax1, NoSocialSecurity, NoSalary
}

func CalculateNetWage(incomeTax53Percentage int, paymentWage, salary float64) float64 {
	wage := CalculateWage(paymentWage, salary)
	incomeTax53 := CalculateIncomeTax53(wage, incomeTax53Percentage)
	netWage := wage - incomeTax53
	return netWage
}

func CalculateWage(paymentWage, salary float64) float64 {
	if paymentWage >= salary {
		return paymentWage - salary
	}
	return paymentWage
}

func CalculateIncomeTax53(wage float64, incomeTax53Percentag int) float64 {
	return wage * (float64(incomeTax53Percentag) / 100)
}

func CalculateNetTransfer(netSalary, netWage float64) float64 {
	return netSalary + netWage
}

func CalculateTotalCoachingCustomerCharging(incomes []model.Incomes) float64 {
	var totalCoachingCustomerCharging float64
	for _, income := range incomes {
		totalCoachingCustomerCharging += income.CoachingCustomerCharging
	}
	return totalCoachingCustomerCharging
}

func CalculateTotalOtherWage(incomes []model.Incomes, company string, travelExpense float64) float64 {
	var totalOtherWage float64
	if company == SiamChamnankitCompany {
		for _, income := range incomes {
			if income.Company == SiamChamnankitCompany {
				totalOtherWage += income.OtherWage
			}
		}
		return totalOtherWage + travelExpense
	}
	if company == ShuhariCompany {
		for _, income := range incomes {
			if income.Company == ShuhariCompany {
				totalOtherWage += income.OtherWage
			}
		}
		return totalOtherWage + travelExpense
	}
	for _, income := range incomes {
		totalOtherWage += income.OtherWage
	}
	return totalOtherWage + travelExpense
}

func CalculateTotalCoachingPaymentRate(incomes []model.Incomes, company string) float64 {
	var totalCoachingPaymentRate float64
	if company == SiamChamnankitCompany {
		for _, income := range incomes {
			if income.Company == SiamChamnankitCompany {
				totalCoachingPaymentRate += income.CoachingPaymentRate
			}
		}
		return totalCoachingPaymentRate
	}
	if company == ShuhariCompany {
		for _, income := range incomes {
			if income.Company == ShuhariCompany {
				totalCoachingPaymentRate += income.CoachingPaymentRate
			}
		}
		return totalCoachingPaymentRate
	}
	for _, income := range incomes {
		totalCoachingPaymentRate += income.CoachingPaymentRate
	}
	return totalCoachingPaymentRate
}

func CalculateTotalTrainingWage(incomes []model.Incomes, company string) float64 {
	var totalCoachingTrainingWage float64
	if company == SiamChamnankitCompany {
		for _, income := range incomes {
			if income.Company == SiamChamnankitCompany {
				totalCoachingTrainingWage += income.TrainingWage
			}
		}
		return totalCoachingTrainingWage
	}
	if company == ShuhariCompany {
		for _, income := range incomes {
			if income.Company == ShuhariCompany {
				totalCoachingTrainingWage += income.TrainingWage
			}
		}
		return totalCoachingTrainingWage
	}
	for _, income := range incomes {
		totalCoachingTrainingWage += income.TrainingWage
	}
	return totalCoachingTrainingWage
}
