package timesheet

import (
	"testing"
	"timesheet/internal/model"

	"github.com/stretchr/testify/assert"
)

func Test_CalculateTotalHour_Input_Incomes_Should_Be_Time_18_0_0(t *testing.T) {
	expected := model.Time{
		Hours:   18,
		Minutes: 0,
		Seconds: 0,
	}
	incomes := []model.Incomes{
		{
			StartTimeAMHours:   9,
			StartTimeAMMinutes: 0,
			StartTimeAMSeconds: 0,
			EndTimeAMHours:     12,
			EndTimeAMMinutes:   0,
			EndTimeAMSeconds:   0,
			StartTimePMHours:   13,
			StartTimePMMinutes: 0,
			StartTimePMSeconds: 0,
			EndTimePMHours:     18,
			EndTimePMMinutes:   30,
			EndTimePMSeconds:   0,
			Overtime:           0,
		},
		{
			StartTimeAMHours:   9,
			StartTimeAMMinutes: 0,
			StartTimeAMSeconds: 0,
			EndTimeAMHours:     12,
			EndTimeAMMinutes:   0,
			EndTimeAMSeconds:   0,
			StartTimePMHours:   13,
			StartTimePMMinutes: 0,
			StartTimePMSeconds: 0,
			EndTimePMHours:     18,
			EndTimePMMinutes:   30,
			EndTimePMSeconds:   0,
			Overtime:           1,
		},
	}

	actual := calculateTotalHour(incomes)

	assert.Equal(t, expected, actual)
}

func Test_CalculateNetSalary_Input_PaymentWage_155000_Salary_80000_IncomeTax1_5000_SocialSecurity_0_Should_Be_NetSalary_75000(t *testing.T) {
	expectedSalary := 80000.00
	expectedIncomeTax1 := 5000.00
	expectedSocialSecurity := 0.00
	expectedNetSalary := 75000.00

	paymentWage := 155000.00
	salary := 80000.00
	incomeTax1 := 5000.00
	socialSecurity := 0.00

	actualSalary, actualIncomeTax1, actualSocialSecurity, actualNetSalary := calculateNetSalary(paymentWage, salary, incomeTax1, socialSecurity)

	assert.Equal(t, expectedSalary, actualSalary)
	assert.Equal(t, expectedIncomeTax1, actualIncomeTax1)
	assert.Equal(t, expectedSocialSecurity, actualSocialSecurity)
	assert.Equal(t, expectedNetSalary, actualNetSalary)
}

func Test_CalculateNetSalary_Input_PaymentWage_10000_Salary_40000_IncomeTax1_5000_SocialSecurity_750_Should_Be_NetSalary_0(t *testing.T) {
	expectedSalary := 0.00
	expectedIncomeTax1 := 0.00
	expectedSocialSecurity := 0.00
	expectedNetSalary := 0.00
	paymentWage := 10000.00
	salary := 40000.00
	incomeTax1 := 5000.00
	socialSecurity := 750.00

	actualSalary, actualIncomeTax1, actualSocialSecurity, actualNetSalary := calculateNetSalary(paymentWage, salary, incomeTax1, socialSecurity)

	assert.Equal(t, expectedSalary, actualSalary)
	assert.Equal(t, expectedIncomeTax1, actualIncomeTax1)
	assert.Equal(t, expectedSocialSecurity, actualSocialSecurity)
	assert.Equal(t, expectedNetSalary, actualNetSalary)
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

func Test_CalculateWage_Input_PaymentWage_155000_Salary_80000_Should_Be_750000(t *testing.T) {
	expected := 75000.00
	paymentWage := 155000.00
	salary := 80000.00

	actual := calculateWage(paymentWage, salary)

	assert.Equal(t, expected, actual)
}

func Test_CalculateWage_Input_PaymentWage_155000_Salary_80000_Should_Be_155000(t *testing.T) {
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

func Test_CalculateTotalCoachingCustomerCharging_Input_Incomes_Should_Be_60000(t *testing.T) {
	expected := 60000.00
	incomes := []model.Incomes{
		{
			Day:                      28,
			StartTimeAMHours:         9,
			StartTimeAMMinutes:       0,
			StartTimeAMSeconds:       0,
			EndTimeAMHours:           12,
			EndTimeAMMinutes:         0,
			EndTimeAMSeconds:         0,
			StartTimePMHours:         13,
			StartTimePMMinutes:       0,
			StartTimePMSeconds:       0,
			EndTimePMHours:           18,
			EndTimePMMinutes:         0,
			EndTimePMSeconds:         0,
			Overtime:                 0,
			TotalHoursHours:          8,
			TotalHoursMinutes:        0,
			TotalHoursSeconds:        0,
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             155000.00,
			OtherWage:                0.00,
			Company:                  "siam_chamnankit",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
		{
			Day:                      29,
			StartTimeAMHours:         9,
			StartTimeAMMinutes:       0,
			StartTimeAMSeconds:       0,
			EndTimeAMHours:           12,
			EndTimeAMMinutes:         0,
			EndTimeAMSeconds:         0,
			StartTimePMHours:         13,
			StartTimePMMinutes:       0,
			StartTimePMSeconds:       0,
			EndTimePMHours:           18,
			EndTimePMMinutes:         0,
			EndTimePMSeconds:         0,
			Overtime:                 0,
			TotalHoursHours:          8,
			TotalHoursMinutes:        0,
			TotalHoursSeconds:        0,
			CoachingCustomerCharging: 15000.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             0.00,
			OtherWage:                0.00,
			Company:                  "siam_chamnankit",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
		{
			Day:                      30,
			StartTimeAMHours:         9,
			StartTimeAMMinutes:       0,
			StartTimeAMSeconds:       0,
			EndTimeAMHours:           12,
			EndTimeAMMinutes:         0,
			EndTimeAMSeconds:         0,
			StartTimePMHours:         13,
			StartTimePMMinutes:       0,
			StartTimePMSeconds:       0,
			EndTimePMHours:           18,
			EndTimePMMinutes:         0,
			EndTimePMSeconds:         0,
			Overtime:                 0,
			TotalHoursHours:          8,
			TotalHoursMinutes:        0,
			TotalHoursSeconds:        0,
			CoachingCustomerCharging: 45000.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             0.00,
			OtherWage:                0.00,
			Company:                  "shuhari",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
	}

	actual := calculateTotalCoachingCustomerCharging(incomes)

	assert.Equal(t, expected, actual)
}

func Test_CalculateTotalOtherWage_Input_Incomes_Company_Shuhari_TravelExpense_0_Should_Be_45000(t *testing.T) {
	expected := 45000.00
	company := "shuhari"
	travelExpense := 0.00
	incomes := []model.Incomes{
		{
			Day:                      28,
			StartTimeAMHours:         9,
			StartTimeAMMinutes:       0,
			StartTimeAMSeconds:       0,
			EndTimeAMHours:           12,
			EndTimeAMMinutes:         0,
			EndTimeAMSeconds:         0,
			StartTimePMHours:         13,
			StartTimePMMinutes:       0,
			StartTimePMSeconds:       0,
			EndTimePMHours:           18,
			EndTimePMMinutes:         0,
			EndTimePMSeconds:         0,
			Overtime:                 0,
			TotalHoursHours:          8,
			TotalHoursMinutes:        0,
			TotalHoursSeconds:        0,
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             155000.00,
			OtherWage:                0.00,
			Company:                  "siam_chamnankit",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
		{
			Day:                      29,
			StartTimeAMHours:         9,
			StartTimeAMMinutes:       0,
			StartTimeAMSeconds:       0,
			EndTimeAMHours:           12,
			EndTimeAMMinutes:         0,
			EndTimeAMSeconds:         0,
			StartTimePMHours:         13,
			StartTimePMMinutes:       0,
			StartTimePMSeconds:       0,
			EndTimePMHours:           18,
			EndTimePMMinutes:         0,
			EndTimePMSeconds:         0,
			Overtime:                 0,
			TotalHoursHours:          8,
			TotalHoursMinutes:        0,
			TotalHoursSeconds:        0,
			CoachingCustomerCharging: 15000.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             0.00,
			OtherWage:                25000.00,
			Company:                  "shuhari",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
		{
			Day:                      30,
			StartTimeAMHours:         9,
			StartTimeAMMinutes:       0,
			StartTimeAMSeconds:       0,
			EndTimeAMHours:           12,
			EndTimeAMMinutes:         0,
			EndTimeAMSeconds:         0,
			StartTimePMHours:         13,
			StartTimePMMinutes:       0,
			StartTimePMSeconds:       0,
			EndTimePMHours:           18,
			EndTimePMMinutes:         0,
			EndTimePMSeconds:         0,
			Overtime:                 0,
			TotalHoursHours:          8,
			TotalHoursMinutes:        0,
			TotalHoursSeconds:        0,
			CoachingCustomerCharging: 45000.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             0.00,
			OtherWage:                20000.00,
			Company:                  "shuhari",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
	}

	actual := calculateTotalOtherWage(incomes, company, travelExpense)

	assert.Equal(t, expected, actual)
}

func Test_CalculateTotalOtherWage_Input_Incomes_TravelExpense_0_Should_Be_55000(t *testing.T) {
	expected := 55000.00
	company := ""
	travelExpense := 0.00
	incomes := []model.Incomes{
		{
			Day:                      28,
			StartTimeAMHours:         9,
			StartTimeAMMinutes:       0,
			StartTimeAMSeconds:       0,
			EndTimeAMHours:           12,
			EndTimeAMMinutes:         0,
			EndTimeAMSeconds:         0,
			StartTimePMHours:         13,
			StartTimePMMinutes:       0,
			StartTimePMSeconds:       0,
			EndTimePMHours:           18,
			EndTimePMMinutes:         0,
			EndTimePMSeconds:         0,
			Overtime:                 0,
			TotalHoursHours:          8,
			TotalHoursMinutes:        0,
			TotalHoursSeconds:        0,
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             155000.00,
			OtherWage:                10000.00,
			Company:                  "siam_chamnankit",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
		{
			Day:                      29,
			StartTimeAMHours:         9,
			StartTimeAMMinutes:       0,
			StartTimeAMSeconds:       0,
			EndTimeAMHours:           12,
			EndTimeAMMinutes:         0,
			EndTimeAMSeconds:         0,
			StartTimePMHours:         13,
			StartTimePMMinutes:       0,
			StartTimePMSeconds:       0,
			EndTimePMHours:           18,
			EndTimePMMinutes:         0,
			EndTimePMSeconds:         0,
			Overtime:                 0,
			TotalHoursHours:          8,
			TotalHoursMinutes:        0,
			TotalHoursSeconds:        0,
			CoachingCustomerCharging: 15000.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             0.00,
			OtherWage:                25000.00,
			Company:                  "shuhari",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
		{
			Day:                      30,
			StartTimeAMHours:         9,
			StartTimeAMMinutes:       0,
			StartTimeAMSeconds:       0,
			EndTimeAMHours:           12,
			EndTimeAMMinutes:         0,
			EndTimeAMSeconds:         0,
			StartTimePMHours:         13,
			StartTimePMMinutes:       0,
			StartTimePMSeconds:       0,
			EndTimePMHours:           18,
			EndTimePMMinutes:         0,
			EndTimePMSeconds:         0,
			Overtime:                 0,
			TotalHoursHours:          8,
			TotalHoursMinutes:        0,
			TotalHoursSeconds:        0,
			CoachingCustomerCharging: 45000.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             0.00,
			OtherWage:                20000.00,
			Company:                  "shuhari",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
	}

	actual := calculateTotalOtherWage(incomes, company, travelExpense)

	assert.Equal(t, expected, actual)
}

func Test_CalculateTotalCoachingPaymentRate_Input_Incomes_Company_SiamChamnankit_Should_Be_10000(t *testing.T) {
	expected := 10000.00
	company := "siam_chamnankit"
	incomes := []model.Incomes{
		{
			Day:                      28,
			StartTimeAMHours:         9,
			StartTimeAMMinutes:       0,
			StartTimeAMSeconds:       0,
			EndTimeAMHours:           12,
			EndTimeAMMinutes:         0,
			EndTimeAMSeconds:         0,
			StartTimePMHours:         13,
			StartTimePMMinutes:       0,
			StartTimePMSeconds:       0,
			EndTimePMHours:           18,
			EndTimePMMinutes:         0,
			EndTimePMSeconds:         0,
			Overtime:                 0,
			TotalHoursHours:          8,
			TotalHoursMinutes:        0,
			TotalHoursSeconds:        0,
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             155000.00,
			OtherWage:                0.00,
			Company:                  "siam_chamnankit",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
		{
			Day:                      29,
			StartTimeAMHours:         9,
			StartTimeAMMinutes:       0,
			StartTimeAMSeconds:       0,
			EndTimeAMHours:           12,
			EndTimeAMMinutes:         0,
			EndTimeAMSeconds:         0,
			StartTimePMHours:         13,
			StartTimePMMinutes:       0,
			StartTimePMSeconds:       0,
			EndTimePMHours:           18,
			EndTimePMMinutes:         0,
			EndTimePMSeconds:         0,
			Overtime:                 0,
			TotalHoursHours:          8,
			TotalHoursMinutes:        0,
			TotalHoursSeconds:        0,
			CoachingCustomerCharging: 15000.00,
			CoachingPaymentRate:      10000.00,
			TrainingWage:             0.00,
			OtherWage:                0.00,
			Company:                  "siam_chamnankit",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
		{
			Day:                      30,
			StartTimeAMHours:         9,
			StartTimeAMMinutes:       0,
			StartTimeAMSeconds:       0,
			EndTimeAMHours:           12,
			EndTimeAMMinutes:         0,
			EndTimeAMSeconds:         0,
			StartTimePMHours:         13,
			StartTimePMMinutes:       0,
			StartTimePMSeconds:       0,
			EndTimePMHours:           18,
			EndTimePMMinutes:         0,
			EndTimePMSeconds:         0,
			Overtime:                 0,
			TotalHoursHours:          8,
			TotalHoursMinutes:        0,
			TotalHoursSeconds:        0,
			CoachingCustomerCharging: 45000.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             0.00,
			OtherWage:                0.00,
			Company:                  "shuhari",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
	}

	actual := calculateTotalCoachingPaymentRate(incomes, company)

	assert.Equal(t, expected, actual)
}

func Test_CalculateTotalCoachingPaymentRate_Input_Incomes_Should_Be_20000(t *testing.T) {
	expected := 20000.00
	company := ""
	incomes := []model.Incomes{
		{
			Day:                      28,
			StartTimeAMHours:         9,
			StartTimeAMMinutes:       0,
			StartTimeAMSeconds:       0,
			EndTimeAMHours:           12,
			EndTimeAMMinutes:         0,
			EndTimeAMSeconds:         0,
			StartTimePMHours:         13,
			StartTimePMMinutes:       0,
			StartTimePMSeconds:       0,
			EndTimePMHours:           18,
			EndTimePMMinutes:         0,
			EndTimePMSeconds:         0,
			Overtime:                 0,
			TotalHoursHours:          8,
			TotalHoursMinutes:        0,
			TotalHoursSeconds:        0,
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             155000.00,
			OtherWage:                0.00,
			Company:                  "siam_chamnankit",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
		{
			Day:                      29,
			StartTimeAMHours:         9,
			StartTimeAMMinutes:       0,
			StartTimeAMSeconds:       0,
			EndTimeAMHours:           12,
			EndTimeAMMinutes:         0,
			EndTimeAMSeconds:         0,
			StartTimePMHours:         13,
			StartTimePMMinutes:       0,
			StartTimePMSeconds:       0,
			EndTimePMHours:           18,
			EndTimePMMinutes:         0,
			EndTimePMSeconds:         0,
			Overtime:                 0,
			TotalHoursHours:          8,
			TotalHoursMinutes:        0,
			TotalHoursSeconds:        0,
			CoachingCustomerCharging: 15000.00,
			CoachingPaymentRate:      10000.00,
			TrainingWage:             0.00,
			OtherWage:                0.00,
			Company:                  "siam_chamnankit",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
		{
			Day:                      30,
			StartTimeAMHours:         9,
			StartTimeAMMinutes:       0,
			StartTimeAMSeconds:       0,
			EndTimeAMHours:           12,
			EndTimeAMMinutes:         0,
			EndTimeAMSeconds:         0,
			StartTimePMHours:         13,
			StartTimePMMinutes:       0,
			StartTimePMSeconds:       0,
			EndTimePMHours:           18,
			EndTimePMMinutes:         0,
			EndTimePMSeconds:         0,
			Overtime:                 0,
			TotalHoursHours:          8,
			TotalHoursMinutes:        0,
			TotalHoursSeconds:        0,
			CoachingCustomerCharging: 45000.00,
			CoachingPaymentRate:      10000.00,
			TrainingWage:             0.00,
			OtherWage:                0.00,
			Company:                  "shuhari",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
	}

	actual := calculateTotalCoachingPaymentRate(incomes, company)

	assert.Equal(t, expected, actual)
}

func Test_CalculateTotalTrainingWage_Input_Incomes_Company_Shuhari_Should_Be_20000(t *testing.T) {
	expected := 20000.00
	company := "shuhari"
	incomes := []model.Incomes{
		{
			Day:                      28,
			StartTimeAMHours:         9,
			StartTimeAMMinutes:       0,
			StartTimeAMSeconds:       0,
			EndTimeAMHours:           12,
			EndTimeAMMinutes:         0,
			EndTimeAMSeconds:         0,
			StartTimePMHours:         13,
			StartTimePMMinutes:       0,
			StartTimePMSeconds:       0,
			EndTimePMHours:           18,
			EndTimePMMinutes:         0,
			EndTimePMSeconds:         0,
			Overtime:                 0,
			TotalHoursHours:          8,
			TotalHoursMinutes:        0,
			TotalHoursSeconds:        0,
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             10000.00,
			OtherWage:                0.00,
			Company:                  "siam_chamnankit",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
		{
			Day:                      29,
			StartTimeAMHours:         9,
			StartTimeAMMinutes:       0,
			StartTimeAMSeconds:       0,
			EndTimeAMHours:           12,
			EndTimeAMMinutes:         0,
			EndTimeAMSeconds:         0,
			StartTimePMHours:         13,
			StartTimePMMinutes:       0,
			StartTimePMSeconds:       0,
			EndTimePMHours:           18,
			EndTimePMMinutes:         0,
			EndTimePMSeconds:         0,
			Overtime:                 0,
			TotalHoursHours:          8,
			TotalHoursMinutes:        0,
			TotalHoursSeconds:        0,
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             10000.00,
			OtherWage:                0.00,
			Company:                  "shuhari",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
		{
			Day:                      30,
			StartTimeAMHours:         9,
			StartTimeAMMinutes:       0,
			StartTimeAMSeconds:       0,
			EndTimeAMHours:           12,
			EndTimeAMMinutes:         0,
			EndTimeAMSeconds:         0,
			StartTimePMHours:         13,
			StartTimePMMinutes:       0,
			StartTimePMSeconds:       0,
			EndTimePMHours:           18,
			EndTimePMMinutes:         0,
			EndTimePMSeconds:         0,
			Overtime:                 0,
			TotalHoursHours:          8,
			TotalHoursMinutes:        0,
			TotalHoursSeconds:        0,
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             10000.00,
			OtherWage:                0.00,
			Company:                  "shuhari",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
	}

	actual := calculateTotalTrainingWage(incomes, company)

	assert.Equal(t, expected, actual)
}

func Test_CalculateTotalTrainingWage_Input_Incomes_Should_Be_30000(t *testing.T) {
	expected := 30000.00
	company := ""
	incomes := []model.Incomes{
		{
			Day:                      28,
			StartTimeAMHours:         9,
			StartTimeAMMinutes:       0,
			StartTimeAMSeconds:       0,
			EndTimeAMHours:           12,
			EndTimeAMMinutes:         0,
			EndTimeAMSeconds:         0,
			StartTimePMHours:         13,
			StartTimePMMinutes:       0,
			StartTimePMSeconds:       0,
			EndTimePMHours:           18,
			EndTimePMMinutes:         0,
			EndTimePMSeconds:         0,
			Overtime:                 0,
			TotalHoursHours:          8,
			TotalHoursMinutes:        0,
			TotalHoursSeconds:        0,
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             10000.00,
			OtherWage:                0.00,
			Company:                  "siam_chamnankit",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
		{
			Day:                      29,
			StartTimeAMHours:         9,
			StartTimeAMMinutes:       0,
			StartTimeAMSeconds:       0,
			EndTimeAMHours:           12,
			EndTimeAMMinutes:         0,
			EndTimeAMSeconds:         0,
			StartTimePMHours:         13,
			StartTimePMMinutes:       0,
			StartTimePMSeconds:       0,
			EndTimePMHours:           18,
			EndTimePMMinutes:         0,
			EndTimePMSeconds:         0,
			Overtime:                 0,
			TotalHoursHours:          8,
			TotalHoursMinutes:        0,
			TotalHoursSeconds:        0,
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             10000.00,
			OtherWage:                0.00,
			Company:                  "shuhari",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
		{
			Day:                      30,
			StartTimeAMHours:         9,
			StartTimeAMMinutes:       0,
			StartTimeAMSeconds:       0,
			EndTimeAMHours:           12,
			EndTimeAMMinutes:         0,
			EndTimeAMSeconds:         0,
			StartTimePMHours:         13,
			StartTimePMMinutes:       0,
			StartTimePMSeconds:       0,
			EndTimePMHours:           18,
			EndTimePMMinutes:         0,
			EndTimePMSeconds:         0,
			Overtime:                 0,
			TotalHoursHours:          8,
			TotalHoursMinutes:        0,
			TotalHoursSeconds:        0,
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             10000.00,
			OtherWage:                0.00,
			Company:                  "shuhari",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
	}

	actual := calculateTotalTrainingWage(incomes, company)

	assert.Equal(t, expected, actual)
}
