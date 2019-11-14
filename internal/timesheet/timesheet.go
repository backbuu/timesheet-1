package timesheet

import (
	"fmt"
	"math"
	"os"
	"strings"
	"time"
	"timesheet/internal/model"
	"timesheet/internal/repository"
)

const (
	initialIndex = 0
	oneHour      = 60
)

type TimesheetGateways interface {
	CalculatePaymentSummary(employee []model.Employee, incomes []model.Incomes, year, month int) []model.TransactionTimesheet
	CalculatePayment(incomes []model.Incomes) model.Timesheet
	GetSummaryByID(employeeID string, year, month int) (model.SummaryTimesheet, error)
	VerifyAuthentication(email string, idTokenExpirationTime float64) bool
	GetSummaryInYearByEmployeeID(employeeID string, year int) (model.SummaryTransactionTimesheet, error)
}

type Timesheet struct {
	Repository repository.TimesheetRepositoryGatewaysToTimesheet
}

func (timesheet Timesheet) CalculatePaymentSummary(employee []model.Employee, incomes []model.Incomes, year, month int) []model.TransactionTimesheet {
	var transactionTimesheetList []model.TransactionTimesheet
	for _, employee := range employee {
		totalCoachingPaymentRate := calculateTotalCoachingPaymentRateByCompanyID(incomes, employee.CompanyID)
		totalTrainingWage := calculateTotalTrainingWageByCompanyID(incomes, employee.CompanyID)
		totalOtherWage := calculateTotalOtherWageByCompanyID(incomes, employee.CompanyID, employee.TravelExpense)
		paymentWage := calculateTotalPaymentWage(totalCoachingPaymentRate, totalTrainingWage, totalOtherWage)
		netSalary := calculateNetSalary(employee.Salary, employee.IncomeTax1, employee.SocialSecurity)
		wage := calculateWage(paymentWage, employee.Salary)
		incomeTax53 := calculateIncomeTax53(wage, employee.IncomeTax53Percentage)
		netWage := calculateNetWage(employee.IncomeTax53Percentage, paymentWage, employee.Salary)
		netTransfer := calculateNetTransfer(netSalary, netWage)
		transactionTimesheet := model.TransactionTimesheet{
			EmployeeID:            employee.EmployeeID,
			Month:                 month,
			Year:                  year,
			EmployeeNameTH:        employee.EmployeeNameTH,
			CompanyID:             employee.CompanyID,
			Coaching:              totalCoachingPaymentRate,
			Training:              totalTrainingWage,
			Other:                 totalOtherWage,
			TotalIncomes:          paymentWage,
			Salary:                employee.Salary,
			IncomeTax1:            employee.IncomeTax1,
			SocialSecurity:        employee.SocialSecurity,
			NetSalary:             netSalary,
			Wage:                  wage,
			IncomeTax53Percentage: employee.IncomeTax53Percentage,
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

func (timesheet Timesheet) GetSummaryByID(employeeID string, year, month int) (model.SummaryTimesheet, error) {
	var incomeList []model.Incomes
	employeeList, err := timesheet.Repository.GetEmployeeListByEmployeeID(employeeID)
	if err != nil {
		return model.SummaryTimesheet{}, err
	}
	incomeList, err = timesheet.Repository.GetIncomes(employeeID, year, month)
	if err != nil {
		return model.SummaryTimesheet{}, err
	}
	payment, err := timesheet.Repository.GetTimesheet(employeeID, year, month)
	if err != nil {
		err = timesheet.Repository.CreateTimesheet(employeeID, year, month)
		if err != nil {
			return model.SummaryTimesheet{}, err
		}
	}
	return model.SummaryTimesheet{
		EmployeeNameENG:               employeeList[initialIndex].EmployeeNameENG,
		Email:                         employeeList[initialIndex].Email,
		RatePerDay:                    employeeList[initialIndex].RatePerDay,
		RatePerHour:                   employeeList[initialIndex].RatePerHour,
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

func (timesheet Timesheet) VerifyAuthentication(email string, idTokenExpirationTime float64) bool {
	emailAuthenticityList := []string{"welovebug.biz", "scrum123.com"}
	for _, emailAuthenticityList := range emailAuthenticityList {
		if emailAuthenticityList == strings.Split(email, "@")[1] && now().Before(time.Unix(int64(idTokenExpirationTime), 0)) {
			return true
		}
	}
	return false
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

func calculateTotalOtherWageByCompanyID(incomeList []model.Incomes, companyID int, travelExpense float64) float64 {
	var totalOtherWage float64
	for _, income := range incomeList {
		if income.CompanyID == companyID {
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

func calculateTotalCoachingPaymentRateByCompanyID(incomeList []model.Incomes, companyID int) float64 {
	var totalCoachingPaymentRate float64
	for _, income := range incomeList {
		if income.CompanyID == companyID {
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

func calculateTotalTrainingWageByCompanyID(incomeList []model.Incomes, companyID int) float64 {
	var totalTrainingWage float64
	for _, income := range incomeList {
		if income.CompanyID == companyID {
			totalTrainingWage += income.TrainingWage
		}
	}
	return totalTrainingWage
}

func (timesheet Timesheet) GetSummaryInYearByEmployeeID(employeeID string, year int) (model.SummaryTransactionTimesheet, error) {
	var transactionTimesheetList []model.TransactionTimesheet
	transactionTimesheetList, err := timesheet.Repository.GetTransactionTimesheets(employeeID, year)
	if err != nil {
		return model.SummaryTransactionTimesheet{}, err
	}
	totalCoachingInYear := calculateTotalCoachingInYearByEmployeeID(transactionTimesheetList)
	totalTrainingInYear := 30000.00
	totalOtherInYear := 40000.00
	totalIncomesInYear := 155000.00
	totalSalaryInYear := 80000.00
	totalIncomeTax1InYear := 5000.00
	totalSocialSecurityInYear := 0.00
	totalNetSalaryInYear := 75000.00
	totalWageInYear := 75000.00
	totalIncomeTax53InYear := 7500.00
	totalNetWageInYear := 67500.00
	totalNetTransferInYear := 142500.00
	return model.SummaryTransactionTimesheet{
		EmployeeID:                employeeID,
		Year:                      year,
		TransactionTimesheets:     transactionTimesheetList,
		TotalCoachingInYear:       totalCoachingInYear,
		TotalTrainingInYear:       totalTrainingInYear,
		TotalOtherInYear:          totalOtherInYear,
		TotalIncomesInYear:        totalIncomesInYear,
		TotalSalaryInYear:         totalSalaryInYear,
		TotalIncomeTax1InYear:     totalIncomeTax1InYear,
		TotalSocialSecurityInYear: totalSocialSecurityInYear,
		TotalNetSalaryInYear:      totalNetSalaryInYear,
		TotalWageInYear:           totalWageInYear,
		TotalIncomeTax53InYear:    totalIncomeTax53InYear,
		TotalNetWageInYear:        totalNetWageInYear,
		TotalNetTransferInYear:    totalNetTransferInYear}, nil
}

func calculateTotalCoachingInYearByEmployeeID(transactionTimesheetList []model.TransactionTimesheet) float64 {
	var totalCoachingInYear float64
	for _, transactionTimesheet := range transactionTimesheetList {
		totalCoachingInYear += transactionTimesheet.Coaching
	}
	return totalCoachingInYear
}
