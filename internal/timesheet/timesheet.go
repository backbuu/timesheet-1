package timesheet

import (
	"fmt"
	"math"
	"os"
	"time"
	"timesheet/internal/model"
	"timesheet/internal/repository"
)

const (
	initialIndex = 0
	oneHour      = 60
)

type TimesheetGateways interface {
	CalculatePaymentSummary(member []model.Member, incomes []model.Incomes, year, month int) []model.TransactionTimesheet
	CalculatePayment(incomes []model.Incomes) model.Timesheet
	GetSummaryByID(memberID string, year, month int) (model.SummaryTimesheet, error)
	VerifyAuthentication(email string, expiry float64, memberID string) string
	GetSummaryInYearByID(memberID string, year int) (model.SummaryTransactionTimesheet, error)
}

type Timesheet struct {
	Repository repository.TimesheetRepositoryGatewaysToTimesheet
}

func (timesheet Timesheet) CalculatePaymentSummary(member []model.Member, incomes []model.Incomes, year, month int) []model.TransactionTimesheet {
	var transactionTimesheetList []model.TransactionTimesheet
	for _, member := range member {
		totalCoachingPaymentRate := calculateTotalCoachingPaymentRateByCompany(incomes, member.Company)
		totalTrainingWage := calculateTotalTrainingWageByCompany(incomes, member.Company)
		totalOtherWage := calculateTotalOtherWageByCompany(incomes, member.Company, member.TravelExpense)
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

func (timesheet Timesheet) CalculatePayment(incomeList []model.Incomes) model.Timesheet {
	totalHour := calculateTotalHours(incomeList)
	totalCoachingCustomerCharging := calculateTotalCoachingCustomerCharging(incomeList)
	totalCoachingPaymentRate := calculateTotalCoachingPaymentRate(incomeList)
	totalTrainingWage := calculateTotalTrainingWage(incomeList)
	totalOtherWage := calculateTotalOtherWage(incomeList)
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
	var incomeList []model.Incomes
	memberList, err := timesheet.Repository.GetMemberListByMemberID(memberID)
	if err != nil {
		return model.SummaryTimesheet{}, err
	}
	incomeList, err = timesheet.Repository.GetIncomes(memberID, year, month)
	if err != nil {
		return model.SummaryTimesheet{}, err
	}
	payment, err := timesheet.Repository.GetTimesheet(memberID, year, month)
	if err != nil {
		err = timesheet.Repository.CreateTimesheet(memberID, year, month)
		if err != nil {
			return model.SummaryTimesheet{}, err
		}
	}
	return model.SummaryTimesheet{
		MemberNameENG:                 memberList[initialIndex].MemberNameENG,
		Email:                         memberList[initialIndex].Email,
		RatePerDay:                    memberList[initialIndex].RatePerDay,
		RatePerHour:                   memberList[initialIndex].RatePerHour,
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

func (timesheet Timesheet) VerifyAuthentication(email string, expiry float64, memberIDRequest string) string {
	memberIDByEmail, err := timesheet.Repository.GetMemberIDByEmail(email)
	if err != nil {
		return err.Error()
	}
	if memberIDByEmail != memberIDRequest || time.Unix(int64(expiry), 0).Before(now()) {
		return "Unauthorized"
	}
	return "Success"
}

func now() time.Time {
	if os.Getenv("FIX_TIME") != "" {
		fixedTime, _ := time.Parse("20060102150405", os.Getenv("FIX_TIME"))
		return fixedTime
	}
	return time.Now()
}

func calculateTotalHours(incomeList []model.Incomes) string {
	var toltalHours time.Duration
	var overtime int
	for _, income := range incomeList {
		toltalHours += income.EndTimeAM.Sub(income.StartTimeAM)
		toltalHours += income.EndTimePM.Sub(income.StartTimePM)
	}
	hours := toltalHours.Hours() + float64(overtime)
	minutes := math.Mod(toltalHours.Minutes(), oneHour)
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
	return wage - incomeTax53
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

func calculateTotalCoachingCustomerCharging(incomeList []model.Incomes) float64 {
	var totalCoachingCustomerCharging float64
	for _, income := range incomeList {
		totalCoachingCustomerCharging += income.CoachingCustomerCharging
	}
	return totalCoachingCustomerCharging
}

func calculateTotalOtherWage(incomeList []model.Incomes) float64 {
	var totalOtherWage float64
	for _, income := range incomeList {
		totalOtherWage += income.OtherWage
	}
	return totalOtherWage
}

func calculateTotalOtherWageByCompany(incomeList []model.Incomes, company string, travelExpense float64) float64 {
	var totalOtherWage float64
	for _, income := range incomeList {
		if income.Company == company {
			totalOtherWage += income.OtherWage
		}
	}
	return totalOtherWage + travelExpense
}

func calculateTotalCoachingPaymentRate(incomeList []model.Incomes) float64 {
	var totalCoachingPaymentRate float64
	for _, income := range incomeList {
		totalCoachingPaymentRate += income.CoachingPaymentRate
	}
	return totalCoachingPaymentRate
}

func calculateTotalCoachingPaymentRateByCompany(incomeList []model.Incomes, company string) float64 {
	var totalCoachingPaymentRate float64
	for _, income := range incomeList {
		if income.Company == company {
			totalCoachingPaymentRate += income.CoachingPaymentRate
		}
	}
	return totalCoachingPaymentRate
}

func calculateTotalTrainingWage(incomeList []model.Incomes) float64 {
	var totalTrainingWage float64
	for _, income := range incomeList {
		totalTrainingWage += income.TrainingWage
	}
	return totalTrainingWage
}

func calculateTotalTrainingWageByCompany(incomeList []model.Incomes, company string) float64 {
	var totalTrainingWage float64
	for _, income := range incomeList {
		if income.Company == company {
			totalTrainingWage += income.TrainingWage
		}
	}
	return totalTrainingWage
}

func (timesheet Timesheet) GetSummaryInYearByID(memberID string, year int) (model.SummaryTransactionTimesheet, error) {
	var transactionTimesheetList []model.TransactionTimesheet
	transactionTimesheetList, err := timesheet.Repository.GetTransactionTimesheets(memberID, year)
	if err != nil {
		return model.SummaryTransactionTimesheet{}, err
	}
	return model.SummaryTransactionTimesheet{
		MemberID:               memberID,
		Year:                   year,
		TransactionTimesheets:  transactionTimesheetList,
		TotalCoachingInYear:    0.00,
		TotalTrainingInYear:    0.00,
		TotalOtherInYear:       0.00,
		TotalIncomesInYear:     0.00,
		TotalSalaryInYear:      0.00,
		TotalNetSalaryInYear:   0.00,
		TotalWageInYear:        0.00,
		TotalNetWageInYear:     0.00,
		TotalNetTransferInYear: 0.00}, nil
}
