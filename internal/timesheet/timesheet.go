package timesheet

import (
	"fmt"
	"math"
	"time"
	"timesheet/internal/model"
	"timesheet/internal/repository"
)

const (
	SiamChamnankitCompany = "siam_chamnankit"
	ShuhariCompany        = "shuhari"

	OneMinute = 60
	OneHour   = 60
)

type TimesheetGateways interface {
	CalculatePaymentSummary(member []model.Member, incomes []model.Incomes, year, month int) []model.TransactionTimesheet
	CalculatePayment(incomes []model.Incomes) model.Timesheet
	GetSummaryByID(memberID string, year, month int) (model.SummaryTimesheet, error)
}

type Timesheet struct {
	Repository repository.TimesheetRepository
}

func (timesheet Timesheet) CalculatePaymentSummary(member []model.Member, incomes []model.Incomes, year, month int) []model.TransactionTimesheet {
	var transactionTimesheetList []model.TransactionTimesheet
	for _, member := range member {
		totalCoachingPaymentRate := calculateTotalCoachingPaymentRate(incomes, member.Company)
		totalTrainingWage := calculateTotalTrainingWage(incomes, member.Company)
		totalOtherWage := calculateTotalOtherWage(incomes, member.Company, member.TravelExpense)
		paymentWage := calculateTotalPaymentWage(totalCoachingPaymentRate, totalTrainingWage, totalOtherWage)
		netSalary := calculateNetSalary(member.Salary, member.IncomeTax1, member.SocialSecurity)
		wage := calculateWage(paymentWage, member.Salary)
		incomeTax53 := calculateIncomeTax53(wage, member.IncomeTax53Percentage)
		netWage := calculateNetWage(member.IncomeTax53Percentage, paymentWage, member.Salary)
		netTransfer := calculateNetTransfer(netSalary, netWage)
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
			Salary:                member.Salary,
			IncomeTax1:            member.IncomeTax1,
			SocialSecurity:        member.SocialSecurity,
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

func (timesheet Timesheet) CalculatePayment(incomes []model.Incomes) model.Timesheet {
	totalHour := calculateTotalHours(incomes)
	totalCoachingCustomerCharging := calculateTotalCoachingCustomerCharging(incomes)
	totalCoachingPaymentRate := calculateTotalCoachingPaymentRate(incomes, "")
	totalTrainingWage := calculateTotalTrainingWage(incomes, "")
	totalOtherWage := calculateTotalOtherWage(incomes, "", 0.00)
	paymentWage := calculateTotalPaymentWage(totalCoachingPaymentRate, totalTrainingWage, totalOtherWage)
	return model.Timesheet{
		TotalHours:                    totalHour,
		TotalCoachingCustomerCharging: totalCoachingCustomerCharging,
		TotalCoachingPaymentRate:      totalCoachingPaymentRate,
		TotalTrainigWage:              totalTrainingWage,
		TotalOtherWage:                totalOtherWage,
		PaymentWage:                   paymentWage,
	}
}

func (timesheet Timesheet) GetSummaryByID(memberID string, year, month int) (model.SummaryTimesheet, error) {
	var payment model.Timesheet
	memberList, err := timesheet.Repository.GetMemberByID(memberID)
	if err != nil {
		return model.SummaryTimesheet{}, err
	}
	incomeList, err := timesheet.Repository.GetIncomes(memberID, year, month)
	if err != nil {
		return model.SummaryTimesheet{}, err
	}
	payment, err = timesheet.Repository.GetTimesheet(memberID, year, month)
	if err != nil {
		err = timesheet.Repository.CreateTimesheet(memberID, year, month)
		if err != nil {
			return model.SummaryTimesheet{}, err
		}
	}
	return model.SummaryTimesheet{
		MemberNameENG:                 memberList[0].MemberNameENG,
		Email:                         memberList[0].Email,
		OvertimeRate:                  memberList[0].OvertimeRate,
		RatePerDay:                    memberList[0].RatePerDay,
		RatePerHour:                   memberList[0].RatePerHour,
		Year:                          year,
		Month:                         month,
		Incomes:                       incomeList,
		TimesheetID:                   payment.ID,
		TotalHours:                    payment.TotalHours,
		TotalCoachingCustomerCharging: payment.TotalCoachingCustomerCharging,
		TotalCoachingPaymentRate:      payment.TotalCoachingPaymentRate,
		TotalTrainigWage:              payment.TotalTrainigWage,
		TotalOtherWage:                payment.TotalOtherWage,
		PaymentWage:                   payment.PaymentWage}, nil
}

func calculateTotalHours(incomes []model.Incomes) string {
	var toltalHours time.Duration
	var overtime int
	for _, income := range incomes {
		toltalHours += income.EndTimeAM.Sub(income.StartTimeAM)
		toltalHours += income.EndTimePM.Sub(income.StartTimePM)
		overtime += income.Overtime
	}
	hours := toltalHours.Hours() + float64(overtime)
	minutes := math.Mod(toltalHours.Minutes(), OneHour)
	if minutes != 0 {
		return fmt.Sprintf("%.0f:%2.0f:00", hours, minutes)
	}
	return fmt.Sprintf("%.0f:00:00", hours)
}

func calculateTotalPaymentWage(coachingPaymentRate, trainingWage, otherWage float64) float64 {
	return coachingPaymentRate + trainingWage + otherWage
}

func calculateNetSalary(salary, incomeTax1, socialSecurity float64) float64 {
	return salary - incomeTax1 - socialSecurity
}

func calculateNetWage(incomeTax53Percentage int, paymentWage, salary float64) float64 {
	wage := calculateWage(paymentWage, salary)
	incomeTax53 := calculateIncomeTax53(wage, incomeTax53Percentage)
	netWage := wage - incomeTax53
	return netWage
}

func calculateWage(paymentWage, salary float64) float64 {
	if paymentWage >= salary {
		return paymentWage - salary
	}
	return paymentWage
}

func calculateIncomeTax53(wage float64, incomeTax53Percentag int) float64 {
	return wage * (float64(incomeTax53Percentag) / 100)
}

func calculateNetTransfer(netSalary, netWage float64) float64 {
	return netSalary + netWage
}

func calculateTotalCoachingCustomerCharging(incomes []model.Incomes) float64 {
	var totalCoachingCustomerCharging float64
	for _, income := range incomes {
		totalCoachingCustomerCharging += income.CoachingCustomerCharging
	}
	return totalCoachingCustomerCharging
}

func calculateTotalOtherWage(incomes []model.Incomes, company string, travelExpense float64) float64 {
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

func calculateTotalCoachingPaymentRate(incomes []model.Incomes, company string) float64 {
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

func calculateTotalTrainingWage(incomes []model.Incomes, company string) float64 {
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
