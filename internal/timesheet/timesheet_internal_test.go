package timesheet

import (
	"testing"
	"time"
	"timesheet/internal/model"

	"github.com/stretchr/testify/assert"
)

func Test_CalculateTotalHour_Input_IncomeList_Should_Be_Time_18_00_00(t *testing.T) {
	expected := "17:00:00"
	startTimeAM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 09:00:00")
	endTimeAM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 12:00:00")
	startTimePM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 13:00:00")
	endTimePM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 18:30:00")
	incomeList := []model.Incomes{
		{
			StartTimeAM: startTimeAM,
			EndTimeAM:   endTimeAM,
			StartTimePM: startTimePM,
			EndTimePM:   endTimePM,
		},
		{
			StartTimeAM: startTimeAM,
			EndTimeAM:   endTimeAM,
			StartTimePM: startTimePM,
			EndTimePM:   endTimePM,
		},
	}

	actual := calculateTotalHours(incomeList)

	assert.Equal(t, expected, actual)
}

func Test_CalculateNetSalary_Input_Salary_80000_IncomeTax1_5000_SocialSecurity_0_Should_Be_75000(t *testing.T) {
	expected := 75000.00
	salary := 80000.00
	incomeTax1 := 5000.00
	socialSecurity := 0.00

	actual := calculateNetSalary(salary, incomeTax1, socialSecurity)

	assert.Equal(t, expected, actual)
}

func Test_CalculateNetSalary_Input_Salary_40000_IncomeTax1_5000_SocialSecurity_750_Should_Be_34250(t *testing.T) {
	expected := 34250.00
	salary := 40000.00
	incomeTax1 := 5000.00
	socialSecurity := 750.00

	actual := calculateNetSalary(salary, incomeTax1, socialSecurity)

	assert.Equal(t, expected, actual)
}

func Test_CalculateNetSalary_Input_Salary_25000_IncomeTax1_0_SocialSecurity_750_Should_Be_24250(t *testing.T) {
	expected := 24250.00
	salary := 25000.00
	incomeTax1 := 0.00
	socialSecurity := 750.00

	actual := calculateNetSalary(salary, incomeTax1, socialSecurity)

	assert.Equal(t, expected, actual)
}

func Test_CalculateNetWage_Input_IncomeTax53Percentage_10_PaymentWage_155000_Salary_80000_Should_Be_67500(t *testing.T) {
	expected := 67500.00
	incomeTax53Percentage := 10
	salary := 80000.00
	paymentWage := 155000.00

	actual := calculateNetWage(incomeTax53Percentage, paymentWage, salary)

	assert.Equal(t, expected, actual)
}

func Test_CalculateIncomeTax53_Input_IncomeTax53Percentage_10_Wage_40000_Should_Be_4000(t *testing.T) {
	expected := 4000.00
	incomeTax53Percentage := 10
	wage := 40000.00

	actual := calculateIncomeTax53(wage, incomeTax53Percentage)

	assert.Equal(t, expected, actual)
}

func Test_CalculateWage_Input_PaymentWage_155000_Salary_80000_Should_Be_75000(t *testing.T) {
	expected := 75000.00
	paymentWage := 155000.00
	salary := 80000.00

	actual := calculateWage(paymentWage, salary)

	assert.Equal(t, expected, actual)
}

func Test_CalculateNetTransfer_Input_NetSalary_75000_NetWage_67500_Should_Be_142500(t *testing.T) {
	expected := 142500.00
	netSalary := 75000.00
	netWage := 67500.00

	actual := calculateNetTransfer(netSalary, netWage)

	assert.Equal(t, expected, actual)
}

func Test_CalculateTotalPaymentWage_Input_CoachingPaymentRate_85000_TrainingWage_70000_OtherWage_40000_Should_Be_195000(t *testing.T) {
	expected := 195000.00
	coachingPaymentRate := 85000.00
	trainingWage := 70000.00
	otherWage := 40000.00

	actual := calculateTotalPaymentWage(coachingPaymentRate, trainingWage, otherWage)

	assert.Equal(t, expected, actual)
}

func Test_CalculateTotalCoachingCustomerCharging_Input_IncomeList_Should_Be_60000(t *testing.T) {
	expected := 60000.00
	startTimeAM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 09:00:00")
	endTimeAM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 12:00:00")
	startTimePM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 13:00:00")
	endTimePM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 18:00:00")
	totalHours, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 08:00:00")
	incomeList := []model.Incomes{
		{
			Day:                      28,
			StartTimeAM:              startTimeAM,
			EndTimeAM:                endTimeAM,
			StartTimePM:              startTimePM,
			EndTimePM:                endTimePM,
			TotalHours:               totalHours,
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             155000.00,
			OtherWage:                0.00,
			CompanyID:                1,
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
		{
			Day:                      29,
			StartTimeAM:              startTimeAM,
			EndTimeAM:                endTimeAM,
			StartTimePM:              startTimePM,
			EndTimePM:                endTimePM,
			TotalHours:               totalHours,
			CoachingCustomerCharging: 15000.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             0.00,
			OtherWage:                0.00,
			CompanyID:                1,
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
		{
			Day:                      30,
			StartTimeAM:              startTimeAM,
			EndTimeAM:                endTimeAM,
			StartTimePM:              startTimePM,
			EndTimePM:                endTimePM,
			TotalHours:               totalHours,
			CoachingCustomerCharging: 45000.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             0.00,
			OtherWage:                0.00,
			CompanyID:                2,
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
	}

	actual := calculateTotalCoachingCustomerCharging(incomeList)

	assert.Equal(t, expected, actual)
}

func Test_CalculateTotalOtherWageBycompany_id_Input_IncomeList_companyID_2_TravelExpense_0_Should_Be_45000(t *testing.T) {
	expected := 45000.00
	companyID := 2
	travelExpense := 0.00
	startTimeAM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 09:00:00")
	endTimeAM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 12:00:00")
	startTimePM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 13:00:00")
	endTimePM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 18:00:00")
	totalHours, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 08:00:00")
	incomeList := []model.Incomes{
		{
			Day:                      28,
			StartTimeAM:              startTimeAM,
			EndTimeAM:                endTimeAM,
			StartTimePM:              startTimePM,
			EndTimePM:                endTimePM,
			TotalHours:               totalHours,
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             155000.00,
			OtherWage:                0.00,
			CompanyID:                1,
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
		{
			Day:                      29,
			StartTimeAM:              startTimeAM,
			EndTimeAM:                endTimeAM,
			StartTimePM:              startTimePM,
			EndTimePM:                endTimePM,
			TotalHours:               totalHours,
			CoachingCustomerCharging: 15000.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             0.00,
			OtherWage:                25000.00,
			CompanyID:                2,
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
		{
			Day:                      30,
			StartTimeAM:              startTimeAM,
			EndTimeAM:                endTimeAM,
			StartTimePM:              startTimePM,
			EndTimePM:                endTimePM,
			TotalHours:               totalHours,
			CoachingCustomerCharging: 45000.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             0.00,
			OtherWage:                20000.00,
			CompanyID:                2,
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
	}

	actual := calculateTotalOtherWageByCompanyID(incomeList, companyID, travelExpense)

	assert.Equal(t, expected, actual)
}

func Test_CalculateTotalOtherWage_Input_IncomeList_Should_Be_55000(t *testing.T) {
	expected := 55000.00
	startTimeAM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 09:00:00")
	endTimeAM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 12:00:00")
	startTimePM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 13:00:00")
	endTimePM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 18:00:00")
	totalHours, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 08:00:00")
	incomeList := []model.Incomes{
		{
			Day:                      28,
			StartTimeAM:              startTimeAM,
			EndTimeAM:                endTimeAM,
			StartTimePM:              startTimePM,
			EndTimePM:                endTimePM,
			TotalHours:               totalHours,
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             155000.00,
			OtherWage:                10000.00,
			CompanyID:                1,
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
		{
			Day:                      29,
			StartTimeAM:              startTimeAM,
			EndTimeAM:                endTimeAM,
			StartTimePM:              startTimePM,
			EndTimePM:                endTimePM,
			TotalHours:               totalHours,
			CoachingCustomerCharging: 15000.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             0.00,
			OtherWage:                25000.00,
			CompanyID:                2,
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
		{
			Day:                      30,
			StartTimeAM:              startTimeAM,
			EndTimeAM:                endTimeAM,
			StartTimePM:              startTimePM,
			EndTimePM:                endTimePM,
			TotalHours:               totalHours,
			CoachingCustomerCharging: 45000.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             0.00,
			OtherWage:                20000.00,
			CompanyID:                2,
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
	}

	actual := calculateTotalOtherWage(incomeList)

	assert.Equal(t, expected, actual)
}

func Test_CalculateTotalCoachingPaymentRateBycompany_id_Input_IncomeList_companyID_1_Should_Be_10000(t *testing.T) {
	expected := 10000.00
	companyID := 1
	startTimeAM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 09:00:00")
	endTimeAM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 12:00:00")
	startTimePM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 13:00:00")
	endTimePM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 18:00:00")
	totalHours, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 08:00:00")
	incomeList := []model.Incomes{
		{
			Day:                      28,
			StartTimeAM:              startTimeAM,
			EndTimeAM:                endTimeAM,
			StartTimePM:              startTimePM,
			EndTimePM:                endTimePM,
			TotalHours:               totalHours,
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             155000.00,
			OtherWage:                0.00,
			CompanyID:                1,
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
		{
			Day:                      29,
			StartTimeAM:              startTimeAM,
			EndTimeAM:                endTimeAM,
			StartTimePM:              startTimePM,
			EndTimePM:                endTimePM,
			TotalHours:               totalHours,
			CoachingCustomerCharging: 15000.00,
			CoachingPaymentRate:      10000.00,
			TrainingWage:             0.00,
			OtherWage:                0.00,
			CompanyID:                1,
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
		{
			Day:                      30,
			StartTimeAM:              startTimeAM,
			EndTimeAM:                endTimeAM,
			StartTimePM:              startTimePM,
			EndTimePM:                endTimePM,
			TotalHours:               totalHours,
			CoachingCustomerCharging: 45000.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             0.00,
			OtherWage:                0.00,
			CompanyID:                2,
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
	}

	actual := calculateTotalCoachingPaymentRateByCompanyID(incomeList, companyID)

	assert.Equal(t, expected, actual)
}

func Test_CalculateTotalCoachingPaymentRate_Input_IncomeList_Should_Be_20000(t *testing.T) {
	expected := 20000.00
	startTimeAM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 09:00:00")
	endTimeAM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 12:00:00")
	startTimePM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 13:00:00")
	endTimePM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 18:00:00")
	totalHours, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 08:00:00")
	incomeList := []model.Incomes{
		{
			Day:                      28,
			StartTimeAM:              startTimeAM,
			EndTimeAM:                endTimeAM,
			StartTimePM:              startTimePM,
			EndTimePM:                endTimePM,
			TotalHours:               totalHours,
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             155000.00,
			OtherWage:                0.00,
			CompanyID:                1,
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
		{
			Day:                      29,
			StartTimeAM:              startTimeAM,
			EndTimeAM:                endTimeAM,
			StartTimePM:              startTimePM,
			EndTimePM:                endTimePM,
			TotalHours:               totalHours,
			CoachingCustomerCharging: 15000.00,
			CoachingPaymentRate:      10000.00,
			TrainingWage:             0.00,
			OtherWage:                0.00,
			CompanyID:                1,
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
		{
			Day:                      30,
			StartTimeAM:              startTimeAM,
			EndTimeAM:                endTimeAM,
			StartTimePM:              startTimePM,
			EndTimePM:                endTimePM,
			TotalHours:               totalHours,
			CoachingCustomerCharging: 45000.00,
			CoachingPaymentRate:      10000.00,
			TrainingWage:             0.00,
			OtherWage:                0.00,
			CompanyID:                2,
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
	}

	actual := calculateTotalCoachingPaymentRate(incomeList)

	assert.Equal(t, expected, actual)
}

func Test_CalculateTotalTrainingWageBycompany_id_Input_IncomeList_companyID_2_Should_Be_20000(t *testing.T) {
	expected := 20000.00
	companyID := 2
	startTimeAM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 09:00:00")
	endTimeAM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 12:00:00")
	startTimePM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 13:00:00")
	endTimePM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 18:00:00")
	totalHours, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 08:00:00")
	incomeList := []model.Incomes{
		{
			Day:                      28,
			StartTimeAM:              startTimeAM,
			EndTimeAM:                endTimeAM,
			StartTimePM:              startTimePM,
			EndTimePM:                endTimePM,
			TotalHours:               totalHours,
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             10000.00,
			OtherWage:                0.00,
			CompanyID:                1,
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
		{
			Day:                      29,
			StartTimeAM:              startTimeAM,
			EndTimeAM:                endTimeAM,
			StartTimePM:              startTimePM,
			EndTimePM:                endTimePM,
			TotalHours:               totalHours,
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             10000.00,
			OtherWage:                0.00,
			CompanyID:                2,
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
		{
			Day:                      30,
			StartTimeAM:              startTimeAM,
			EndTimeAM:                endTimeAM,
			StartTimePM:              startTimePM,
			EndTimePM:                endTimePM,
			TotalHours:               totalHours,
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             10000.00,
			OtherWage:                0.00,
			CompanyID:                2,
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
	}

	actual := calculateTotalTrainingWageByCompanyID(incomeList, companyID)

	assert.Equal(t, expected, actual)
}

func Test_CalculateTotalTrainingWage_Input_IncomeList_Should_Be_30000(t *testing.T) {
	expected := 30000.00
	startTimeAM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 09:00:00")
	endTimeAM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 12:00:00")
	startTimePM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 13:00:00")
	endTimePM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 18:00:00")
	totalHours, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 08:00:00")
	incomeList := []model.Incomes{
		{
			Day:                      28,
			StartTimeAM:              startTimeAM,
			EndTimeAM:                endTimeAM,
			StartTimePM:              startTimePM,
			EndTimePM:                endTimePM,
			TotalHours:               totalHours,
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             10000.00,
			OtherWage:                0.00,
			CompanyID:                1,
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
		{
			Day:                      29,
			StartTimeAM:              startTimeAM,
			EndTimeAM:                endTimeAM,
			StartTimePM:              startTimePM,
			EndTimePM:                endTimePM,
			TotalHours:               totalHours,
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             10000.00,
			OtherWage:                0.00,
			CompanyID:                2,
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
		{
			Day:                      30,
			StartTimeAM:              startTimeAM,
			EndTimeAM:                endTimeAM,
			StartTimePM:              startTimePM,
			EndTimePM:                endTimePM,
			TotalHours:               totalHours,
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             10000.00,
			OtherWage:                0.00,
			CompanyID:                2,
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
	}

	actual := calculateTotalTrainingWage(incomeList)

	assert.Equal(t, expected, actual)
}

func Test_CalculateTotalCoachingInYearByEmployeeID_Input_TransactionTimesheetList_Should_Be_340000(t *testing.T) {
	expected := 340000.00
	transactionTimesheetList := []model.TransactionTimesheet{
		{
			ID:                     "00120171001",
			EmployeeID:             "001",
			EmployeeNameTH:         "ประธาน ด่านสกุลเจริญกิจ",
			Month:                  10,
			Year:                   2017,
			CompanyID:              1,
			Coaching:               85000.00,
			Training:               30000.00,
			Other:                  40000.00,
			TotalIncomes:           155000.00,
			Salary:                 80000.00,
			IncomeTax1:             5000.00,
			SocialSecurity:         0.00,
			NetSalary:              75000.00,
			Wage:                   75000.00,
			IncomeTax53Percentage:  10,
			IncomeTax53:            7500.00,
			NetWage:                67500.00,
			NetTransfer:            142500.00,
			StatusCheckingTransfer: "รอการตรวจสอบ",
			DateTransfer:           "",
			Comment:                "",
		},
		{
			ID:                     "00120171101",
			EmployeeID:             "001",
			EmployeeNameTH:         "ประธาน ด่านสกุลเจริญกิจ",
			Month:                  11,
			Year:                   2017,
			CompanyID:              1,
			Coaching:               85000.00,
			Training:               30000.00,
			Other:                  40000.00,
			TotalIncomes:           155000.00,
			Salary:                 80000.00,
			IncomeTax1:             5000.00,
			SocialSecurity:         0.00,
			NetSalary:              75000.00,
			Wage:                   75000.00,
			IncomeTax53Percentage:  10,
			IncomeTax53:            7500.00,
			NetWage:                67500.00,
			NetTransfer:            142500.00,
			StatusCheckingTransfer: "รอการตรวจสอบ",
			DateTransfer:           "",
			Comment:                "",
		},
		{
			ID:                     "00120171201",
			EmployeeID:             "001",
			EmployeeNameTH:         "ประธาน ด่านสกุลเจริญกิจ",
			Month:                  12,
			Year:                   2017,
			CompanyID:              1,
			Coaching:               85000.00,
			Training:               30000.00,
			Other:                  40000.00,
			TotalIncomes:           155000.00,
			Salary:                 80000.00,
			IncomeTax1:             5000.00,
			SocialSecurity:         0.00,
			NetSalary:              75000.00,
			Wage:                   75000.00,
			IncomeTax53Percentage:  10,
			IncomeTax53:            7500.00,
			NetWage:                67500.00,
			NetTransfer:            142500.00,
			StatusCheckingTransfer: "รอการตรวจสอบ",
			DateTransfer:           "",
			Comment:                "",
		},
		{
			ID:                     "00120171201",
			EmployeeID:             "001",
			EmployeeNameTH:         "ประธาน ด่านสกุลเจริญกิจ",
			Month:                  12,
			Year:                   2017,
			CompanyID:              2,
			Coaching:               85000.00,
			Training:               30000.00,
			Other:                  40000.00,
			TotalIncomes:           155000.00,
			Salary:                 80000.00,
			IncomeTax1:             5000.00,
			SocialSecurity:         0.00,
			NetSalary:              75000.00,
			Wage:                   75000.00,
			IncomeTax53Percentage:  10,
			IncomeTax53:            7500.00,
			NetWage:                67500.00,
			NetTransfer:            142500.00,
			StatusCheckingTransfer: "รอการตรวจสอบ",
			DateTransfer:           "",
			Comment:                "",
		},
	}

	actual := calculateTotalCoachingInYearByEmployeeID(transactionTimesheetList)

	assert.Equal(t, expected, actual)
}

func Test_CalculateTotalTrainingInYearByEmployeeID_Input_TransactionTimesheetList_Should_Be_120000(t *testing.T) {
	expected := 120000.00
	transactionTimesheetList := []model.TransactionTimesheet{
		{
			ID:                     "00120171001",
			EmployeeID:             "001",
			EmployeeNameTH:         "ประธาน ด่านสกุลเจริญกิจ",
			Month:                  10,
			Year:                   2017,
			CompanyID:              1,
			Coaching:               85000.00,
			Training:               30000.00,
			Other:                  40000.00,
			TotalIncomes:           155000.00,
			Salary:                 80000.00,
			IncomeTax1:             5000.00,
			SocialSecurity:         0.00,
			NetSalary:              75000.00,
			Wage:                   75000.00,
			IncomeTax53Percentage:  10,
			IncomeTax53:            7500.00,
			NetWage:                67500.00,
			NetTransfer:            142500.00,
			StatusCheckingTransfer: "รอการตรวจสอบ",
			DateTransfer:           "",
			Comment:                "",
		},
		{
			ID:                     "00120171101",
			EmployeeID:             "001",
			EmployeeNameTH:         "ประธาน ด่านสกุลเจริญกิจ",
			Month:                  11,
			Year:                   2017,
			CompanyID:              1,
			Coaching:               85000.00,
			Training:               30000.00,
			Other:                  40000.00,
			TotalIncomes:           155000.00,
			Salary:                 80000.00,
			IncomeTax1:             5000.00,
			SocialSecurity:         0.00,
			NetSalary:              75000.00,
			Wage:                   75000.00,
			IncomeTax53Percentage:  10,
			IncomeTax53:            7500.00,
			NetWage:                67500.00,
			NetTransfer:            142500.00,
			StatusCheckingTransfer: "รอการตรวจสอบ",
			DateTransfer:           "",
			Comment:                "",
		},
		{
			ID:                     "00120171201",
			EmployeeID:             "001",
			EmployeeNameTH:         "ประธาน ด่านสกุลเจริญกิจ",
			Month:                  12,
			Year:                   2017,
			CompanyID:              1,
			Coaching:               85000.00,
			Training:               30000.00,
			Other:                  40000.00,
			TotalIncomes:           155000.00,
			Salary:                 80000.00,
			IncomeTax1:             5000.00,
			SocialSecurity:         0.00,
			NetSalary:              75000.00,
			Wage:                   75000.00,
			IncomeTax53Percentage:  10,
			IncomeTax53:            7500.00,
			NetWage:                67500.00,
			NetTransfer:            142500.00,
			StatusCheckingTransfer: "รอการตรวจสอบ",
			DateTransfer:           "",
			Comment:                "",
		},
		{
			ID:                     "00120171201",
			EmployeeID:             "001",
			EmployeeNameTH:         "ประธาน ด่านสกุลเจริญกิจ",
			Month:                  12,
			Year:                   2017,
			CompanyID:              2,
			Coaching:               85000.00,
			Training:               30000.00,
			Other:                  40000.00,
			TotalIncomes:           155000.00,
			Salary:                 80000.00,
			IncomeTax1:             5000.00,
			SocialSecurity:         0.00,
			NetSalary:              75000.00,
			Wage:                   75000.00,
			IncomeTax53Percentage:  10,
			IncomeTax53:            7500.00,
			NetWage:                67500.00,
			NetTransfer:            142500.00,
			StatusCheckingTransfer: "รอการตรวจสอบ",
			DateTransfer:           "",
			Comment:                "",
		},
	}

	actual := calculateTotalTrainingInYearByEmployeeID(transactionTimesheetList)

	assert.Equal(t, expected, actual)
}

func Test_CalculateTotalOtherInYearByEmployeeID_Input_TransactionTimesheetList_Should_Be_160000(t *testing.T) {
	expected := 160000.00
	transactionTimesheetList := []model.TransactionTimesheet{
		{
			ID:                     "00120171001",
			EmployeeID:             "001",
			EmployeeNameTH:         "ประธาน ด่านสกุลเจริญกิจ",
			Month:                  10,
			Year:                   2017,
			CompanyID:              1,
			Coaching:               85000.00,
			Training:               30000.00,
			Other:                  40000.00,
			TotalIncomes:           155000.00,
			Salary:                 80000.00,
			IncomeTax1:             5000.00,
			SocialSecurity:         0.00,
			NetSalary:              75000.00,
			Wage:                   75000.00,
			IncomeTax53Percentage:  10,
			IncomeTax53:            7500.00,
			NetWage:                67500.00,
			NetTransfer:            142500.00,
			StatusCheckingTransfer: "รอการตรวจสอบ",
			DateTransfer:           "",
			Comment:                "",
		},
		{
			ID:                     "00120171101",
			EmployeeID:             "001",
			EmployeeNameTH:         "ประธาน ด่านสกุลเจริญกิจ",
			Month:                  11,
			Year:                   2017,
			CompanyID:              1,
			Coaching:               85000.00,
			Training:               30000.00,
			Other:                  40000.00,
			TotalIncomes:           155000.00,
			Salary:                 80000.00,
			IncomeTax1:             5000.00,
			SocialSecurity:         0.00,
			NetSalary:              75000.00,
			Wage:                   75000.00,
			IncomeTax53Percentage:  10,
			IncomeTax53:            7500.00,
			NetWage:                67500.00,
			NetTransfer:            142500.00,
			StatusCheckingTransfer: "รอการตรวจสอบ",
			DateTransfer:           "",
			Comment:                "",
		},
		{
			ID:                     "00120171201",
			EmployeeID:             "001",
			EmployeeNameTH:         "ประธาน ด่านสกุลเจริญกิจ",
			Month:                  12,
			Year:                   2017,
			CompanyID:              1,
			Coaching:               85000.00,
			Training:               30000.00,
			Other:                  40000.00,
			TotalIncomes:           155000.00,
			Salary:                 80000.00,
			IncomeTax1:             5000.00,
			SocialSecurity:         0.00,
			NetSalary:              75000.00,
			Wage:                   75000.00,
			IncomeTax53Percentage:  10,
			IncomeTax53:            7500.00,
			NetWage:                67500.00,
			NetTransfer:            142500.00,
			StatusCheckingTransfer: "รอการตรวจสอบ",
			DateTransfer:           "",
			Comment:                "",
		},
		{
			ID:                     "00120171201",
			EmployeeID:             "001",
			EmployeeNameTH:         "ประธาน ด่านสกุลเจริญกิจ",
			Month:                  12,
			Year:                   2017,
			CompanyID:              2,
			Coaching:               85000.00,
			Training:               30000.00,
			Other:                  40000.00,
			TotalIncomes:           155000.00,
			Salary:                 80000.00,
			IncomeTax1:             5000.00,
			SocialSecurity:         0.00,
			NetSalary:              75000.00,
			Wage:                   75000.00,
			IncomeTax53Percentage:  10,
			IncomeTax53:            7500.00,
			NetWage:                67500.00,
			NetTransfer:            142500.00,
			StatusCheckingTransfer: "รอการตรวจสอบ",
			DateTransfer:           "",
			Comment:                "",
		},
	}

	actual := calculateTotalOtherInYearByEmployeeID(transactionTimesheetList)

	assert.Equal(t, expected, actual)
}

func Test_CalculateTotalIncomesInYearByEmployeeID_Input_TransactionTimesheetList_Should_Be_620000(t *testing.T) {
	expected := 620000.00
	transactionTimesheetList := []model.TransactionTimesheet{
		{
			ID:                     "00120171001",
			EmployeeID:             "001",
			EmployeeNameTH:         "ประธาน ด่านสกุลเจริญกิจ",
			Month:                  10,
			Year:                   2017,
			CompanyID:              1,
			Coaching:               85000.00,
			Training:               30000.00,
			Other:                  40000.00,
			TotalIncomes:           155000.00,
			Salary:                 80000.00,
			IncomeTax1:             5000.00,
			SocialSecurity:         0.00,
			NetSalary:              75000.00,
			Wage:                   75000.00,
			IncomeTax53Percentage:  10,
			IncomeTax53:            7500.00,
			NetWage:                67500.00,
			NetTransfer:            142500.00,
			StatusCheckingTransfer: "รอการตรวจสอบ",
			DateTransfer:           "",
			Comment:                "",
		},
		{
			ID:                     "00120171101",
			EmployeeID:             "001",
			EmployeeNameTH:         "ประธาน ด่านสกุลเจริญกิจ",
			Month:                  11,
			Year:                   2017,
			CompanyID:              1,
			Coaching:               85000.00,
			Training:               30000.00,
			Other:                  40000.00,
			TotalIncomes:           155000.00,
			Salary:                 80000.00,
			IncomeTax1:             5000.00,
			SocialSecurity:         0.00,
			NetSalary:              75000.00,
			Wage:                   75000.00,
			IncomeTax53Percentage:  10,
			IncomeTax53:            7500.00,
			NetWage:                67500.00,
			NetTransfer:            142500.00,
			StatusCheckingTransfer: "รอการตรวจสอบ",
			DateTransfer:           "",
			Comment:                "",
		},
		{
			ID:                     "00120171201",
			EmployeeID:             "001",
			EmployeeNameTH:         "ประธาน ด่านสกุลเจริญกิจ",
			Month:                  12,
			Year:                   2017,
			CompanyID:              1,
			Coaching:               85000.00,
			Training:               30000.00,
			Other:                  40000.00,
			TotalIncomes:           155000.00,
			Salary:                 80000.00,
			IncomeTax1:             5000.00,
			SocialSecurity:         0.00,
			NetSalary:              75000.00,
			Wage:                   75000.00,
			IncomeTax53Percentage:  10,
			IncomeTax53:            7500.00,
			NetWage:                67500.00,
			NetTransfer:            142500.00,
			StatusCheckingTransfer: "รอการตรวจสอบ",
			DateTransfer:           "",
			Comment:                "",
		},
		{
			ID:                     "00120171201",
			EmployeeID:             "001",
			EmployeeNameTH:         "ประธาน ด่านสกุลเจริญกิจ",
			Month:                  12,
			Year:                   2017,
			CompanyID:              2,
			Coaching:               85000.00,
			Training:               30000.00,
			Other:                  40000.00,
			TotalIncomes:           155000.00,
			Salary:                 80000.00,
			IncomeTax1:             5000.00,
			SocialSecurity:         0.00,
			NetSalary:              75000.00,
			Wage:                   75000.00,
			IncomeTax53Percentage:  10,
			IncomeTax53:            7500.00,
			NetWage:                67500.00,
			NetTransfer:            142500.00,
			StatusCheckingTransfer: "รอการตรวจสอบ",
			DateTransfer:           "",
			Comment:                "",
		},
	}

	actual := calculateTotalIncomesInYearByEmployeeID(transactionTimesheetList)

	assert.Equal(t, expected, actual)
}
