package timesheet_test

import (
	"errors"
	"os"
	"testing"
	"time"
	"timesheet/internal/mockinternal"
	"timesheet/internal/model"
	. "timesheet/internal/timesheet"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CalculatePaymentSummary_Input_Employee_EmployeeID_001_Should_Be_TransactionTimesheet(t *testing.T) {
	expected := []model.TransactionTimesheet{
		{
			EmployeeID:            "001",
			EmployeeNameTH:        "ประธาน ด่านสกุลเจริญกิจ",
			Year:                  2018,
			Month:                 12,
			CompanyID:             1,
			Coaching:              0.00,
			Training:              155000.00,
			Other:                 0.00,
			TotalIncomes:          155000.00,
			Salary:                80000.00,
			IncomeTax1:            5000.00,
			SocialSecurity:        0.00,
			NetSalary:             75000.00,
			Wage:                  75000.00,
			IncomeTax53Percentage: 10,
			IncomeTax53:           7500.00,
			NetWage:               67500.00,
			NetTransfer:           142500.00,
		},
		{
			EmployeeID:            "001",
			EmployeeNameTH:        "ประธาน ด่านสกุลเจริญกิจ",
			Year:                  2018,
			Month:                 12,
			CompanyID:             2,
			Coaching:              0.00,
			Training:              20000.00,
			Other:                 0.00,
			TotalIncomes:          20000.00,
			Salary:                0.00,
			IncomeTax1:            0.00,
			SocialSecurity:        0.00,
			NetSalary:             0.00,
			Wage:                  20000.00,
			IncomeTax53Percentage: 10,
			IncomeTax53:           2000.00,
			NetWage:               18000.00,
			NetTransfer:           18000.00,
		},
	}
	employee := []model.Employee{
		{
			EmployeeID:            "001",
			CompanyID:             1,
			EmployeeNameTH:        "ประธาน ด่านสกุลเจริญกิจ",
			EmployeeNameENG:       "Prathan Dansakulcharoenkit",
			Email:                 "prathan@scrum123.com",
			RatePerDay:            15000.00,
			RatePerHour:           1875.00,
			Salary:                80000.00,
			IncomeTax1:            5000.00,
			SocialSecurity:        0.00,
			IncomeTax53Percentage: 10,
			TravelExpense:         0.00,
		},
		{
			EmployeeID:            "001",
			CompanyID:             2,
			EmployeeNameTH:        "ประธาน ด่านสกุลเจริญกิจ",
			EmployeeNameENG:       "Prathan Dansakulcharoenkit",
			Email:                 "prathan@scrum123.com",
			RatePerDay:            15000.00,
			RatePerHour:           1875.00,
			Salary:                0.00,
			IncomeTax1:            0.00,
			SocialSecurity:        0.00,
			IncomeTax53Percentage: 10,
			TravelExpense:         0.00,
		},
	}
	startTimeAM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 09:00:00")
	endTimeAM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 12:00:00")
	startTimePM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 13:00:00")
	endTimePM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 18:00:00")
	totalHours, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 08:00:00")
	incomes := []model.Incomes{
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
			Day:                      3,
			StartTimeAM:              startTimeAM,
			EndTimeAM:                endTimeAM,
			StartTimePM:              startTimePM,
			EndTimePM:                endTimePM,
			TotalHours:               totalHours,
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             20000.00,
			OtherWage:                0.00,
			CompanyID:                2,
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
	}
	year := 2018
	month := 12

	timesheet := Timesheet{}
	actual := timesheet.CalculatePaymentSummary(employee, incomes, year, month)

	assert.Equal(t, expected, actual)
}

func Test_CalculatePaymentSummary_Input_Employee_EmployeeID_001_Should_Be_Append_One_TransactionTimesheet(t *testing.T) {
	expected := []model.TransactionTimesheet{
		{
			EmployeeID:            "001",
			EmployeeNameTH:        "ประธาน ด่านสกุลเจริญกิจ",
			Year:                  2018,
			Month:                 12,
			CompanyID:             1,
			Coaching:              0.00,
			Training:              155000.00,
			Other:                 0.00,
			TotalIncomes:          155000.00,
			Salary:                80000.00,
			IncomeTax1:            5000.00,
			SocialSecurity:        0.00,
			NetSalary:             75000.00,
			Wage:                  75000.00,
			IncomeTax53Percentage: 10,
			IncomeTax53:           7500.00,
			NetWage:               67500.00,
			NetTransfer:           142500.00,
		},
	}
	employee := []model.Employee{
		{
			EmployeeID:            "001",
			CompanyID:             1,
			EmployeeNameTH:        "ประธาน ด่านสกุลเจริญกิจ",
			EmployeeNameENG:       "Prathan Dansakulcharoenkit",
			Email:                 "prathan@scrum123.com",
			RatePerDay:            15000.00,
			RatePerHour:           1875.00,
			Salary:                80000.00,
			IncomeTax1:            5000.00,
			SocialSecurity:        0.00,
			IncomeTax53Percentage: 10,
			TravelExpense:         0.00,
		},
		{
			EmployeeID:            "001",
			CompanyID:             2,
			EmployeeNameTH:        "ประธาน ด่านสกุลเจริญกิจ",
			EmployeeNameENG:       "Prathan Dansakulcharoenkit",
			Email:                 "prathan@scrum123.com",
			RatePerDay:            15000.00,
			RatePerHour:           1875.00,
			Salary:                0.00,
			IncomeTax1:            0.00,
			SocialSecurity:        0.00,
			IncomeTax53Percentage: 10,
			TravelExpense:         0.00,
		},
	}
	startTimeAM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 09:00:00")
	endTimeAM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 12:00:00")
	startTimePM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 13:00:00")
	endTimePM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 18:00:00")
	totalHours, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 08:00:00")
	incomes := []model.Incomes{
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
	}
	year := 2018
	month := 12

	timesheet := Timesheet{}
	actual := timesheet.CalculatePaymentSummary(employee, incomes, year, month)

	assert.Equal(t, expected, actual)
}

func Test_CalculatePaymentSummary_Input_Employee_EmployeeID_002_Should_Be_TransactionTimesheet(t *testing.T) {
	expected := []model.TransactionTimesheet{
		{
			EmployeeID:            "002",
			EmployeeNameTH:        "นารีนารถ เนรัญชร",
			Year:                  2018,
			Month:                 12,
			CompanyID:             2,
			Coaching:              0.00,
			Training:              0.00,
			Other:                 6500.00,
			TotalIncomes:          6500.00,
			Salary:                25000.00,
			IncomeTax1:            0.00,
			SocialSecurity:        750.00,
			NetSalary:             24250.00,
			Wage:                  6500.00,
			IncomeTax53Percentage: 5,
			IncomeTax53:           325.00,
			NetWage:               6175.00,
			NetTransfer:           30425.00,
		},
	}
	employee := []model.Employee{
		{
			EmployeeID:            "002",
			CompanyID:             2,
			EmployeeNameTH:        "นารีนารถ เนรัญชร",
			EmployeeNameENG:       "Nareenart Narunchon",
			Email:                 "nareenart@scrum123.com",
			RatePerDay:            0.00,
			RatePerHour:           0.00,
			Salary:                25000.00,
			IncomeTax1:            0.00,
			SocialSecurity:        750.00,
			IncomeTax53Percentage: 5,
			TravelExpense:         1500.00,
			Status:                "salary",
		},
	}
	startTimeAM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 09:00:00")
	endTimeAM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 12:00:00")
	startTimePM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 13:00:00")
	endTimePM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 18:00:00")
	totalHours, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 08:00:00")
	incomes := []model.Incomes{
		{
			Day:                      1,
			StartTimeAM:              startTimeAM,
			EndTimeAM:                endTimeAM,
			StartTimePM:              startTimePM,
			EndTimePM:                endTimePM,
			TotalHours:               totalHours,
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             0.00,
			OtherWage:                5000.00,
			CompanyID:                2,
			Description:              "work at TN",
		},
	}
	year := 2018
	month := 12

	timesheet := Timesheet{}
	actual := timesheet.CalculatePaymentSummary(employee, incomes, year, month)

	assert.Equal(t, expected, actual)
}

func Test_CalculatePayment_Input_Income_CoachingCustomerCharging_15000_CoachingPaymentRate_10000_TrainigWage_20000_Should_Be_Payment(t *testing.T) {
	expected := model.Timesheet{
		TotalHours:                    "16:00:00",
		TotalCoachingCustomerCharging: 15000.00,
		TotalCoachingPaymentRate:      10000.00,
		TotalTrainigWage:              20000.00,
		TotalOtherWage:                0.00,
		PaymentWage:                   30000.00,
	}
	startTimeAM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 09:00:00")
	endTimeAM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 12:00:00")
	startTimePM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 13:00:00")
	endTimePM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 18:00:00")
	totalHours, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 08:00:00")
	incomes := []model.Incomes{
		{
			Day:                      28,
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
			Day:                      3,
			StartTimeAM:              startTimeAM,
			EndTimeAM:                endTimeAM,
			StartTimePM:              startTimePM,
			EndTimePM:                endTimePM,
			TotalHours:               totalHours,
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             20000.00,
			OtherWage:                0.00,
			CompanyID:                2,
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
	}
	timesheet := Timesheet{}

	actual := timesheet.CalculatePayment(incomes)

	assert.Equal(t, expected, actual)
}

func Test_GetSummaryByID_Input_EmployeeID_003_Year_2019_Month_12_Should_Be_SummaryTimesheet(t *testing.T) {
	startTimeAM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 09:00:00")
	endTimeAM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 12:00:00")
	startTimePM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 13:00:00")
	endTimePM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 18:00:00")
	totalHours, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 08:00:00")
	expected := model.SummaryTimesheet{
		EmployeeNameENG: "Somkiat Puisungnoen",
		Email:           "somkiat@scrum123.com",
		RatePerDay:      15000.00,
		RatePerHour:     1875.00,
		Year:            2019,
		Month:           12,
		Incomes: []model.Incomes{
			{
				ID:                       61,
				EmployeeID:               "003",
				Month:                    12,
				Year:                     2019,
				Day:                      1,
				StartTimeAM:              startTimeAM,
				EndTimeAM:                endTimeAM,
				StartTimePM:              startTimePM,
				EndTimePM:                endTimePM,
				TotalHours:               totalHours,
				CoachingCustomerCharging: 0.00,
				CoachingPaymentRate:      0.00,
				TrainingWage:             40000.00,
				OtherWage:                0.00,
				CompanyID:                2,
				Description:              "Technical Excellence at Khonkean",
			},
			{
				ID:                       62,
				EmployeeID:               "003",
				Month:                    12,
				Year:                     2019,
				Day:                      2,
				StartTimeAM:              startTimeAM,
				EndTimeAM:                endTimeAM,
				StartTimePM:              startTimePM,
				EndTimePM:                endTimePM,
				TotalHours:               totalHours,
				CoachingCustomerCharging: 0.00,
				CoachingPaymentRate:      0.00,
				TrainingWage:             40000.00,
				OtherWage:                0.00,
				CompanyID:                2,
				Description:              "Technical Excellence at Khonkean",
			},
		},
		TimesheetID:                   "003201912",
		TotalHours:                    "16:00:00",
		TotalCoachingCustomerCharging: 0.00,
		TotalCoachingPaymentRate:      0.00,
		TotalTrainigWage:              80000.00,
		TotalOtherWage:                0.00,
		PaymentWage:                   80000.00,
	}
	mockRepositoryToTimesheet := new(mockinternal.MockRepositoryToTimesheet)
	mockRepositoryToTimesheet.On("GetEmployeeListByEmployeeID", "003").Return([]model.Employee{
		{
			ID:                    4,
			EmployeeID:            "003",
			CompanyID:             1,
			EmployeeNameTH:        "สมเกียรติ ปุ๋ยสูงเนิน",
			EmployeeNameENG:       "Somkiat Puisungnoen",
			Email:                 "somkiat@scrum123.com",
			RatePerDay:            15000.00,
			RatePerHour:           1875.00,
			Salary:                15000.00,
			IncomeTax1:            0.00,
			SocialSecurity:        750.00,
			IncomeTax53Percentage: 10,
			Status:                "wage",
			TravelExpense:         0.00,
		},
		{
			ID:                    5,
			EmployeeID:            "003",
			CompanyID:             2,
			EmployeeNameTH:        "สมเกียรติ ปุ๋ยสูงเนิน",
			EmployeeNameENG:       "Somkiat Puisungnoen",
			Email:                 "somkiat@scrum123.com",
			RatePerDay:            15000.00,
			RatePerHour:           1875.00,
			Salary:                40000.00,
			IncomeTax1:            5000.00,
			SocialSecurity:        0.00,
			IncomeTax53Percentage: 10,
			Status:                "wage",
			TravelExpense:         0.00,
		},
	}, nil)

	mockRepositoryToTimesheet.On("GetIncomes", "003", 2019, 12).Return([]model.Incomes{
		{
			ID:                       61,
			EmployeeID:               "003",
			Month:                    12,
			Year:                     2019,
			Day:                      1,
			StartTimeAM:              startTimeAM,
			EndTimeAM:                endTimeAM,
			StartTimePM:              startTimePM,
			EndTimePM:                endTimePM,
			TotalHours:               totalHours,
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             40000.00,
			OtherWage:                0.00,
			CompanyID:                2,
			Description:              "Technical Excellence at Khonkean",
		},
		{
			ID:                       62,
			EmployeeID:               "003",
			Month:                    12,
			Year:                     2019,
			Day:                      2,
			StartTimeAM:              startTimeAM,
			EndTimeAM:                endTimeAM,
			StartTimePM:              startTimePM,
			EndTimePM:                endTimePM,
			TotalHours:               totalHours,
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             40000.00,
			OtherWage:                0.00,
			CompanyID:                2,
			Description:              "Technical Excellence at Khonkean",
		},
	}, nil)

	mockRepositoryToTimesheet.On("GetTimesheet", "003", 2019, 12).Return(model.Timesheet{
		ID:                            "003201912",
		EmployeeID:                    "003",
		Month:                         12,
		Year:                          2019,
		RatePerDay:                    15000.00,
		RatePerHour:                   1875.00,
		TotalHours:                    "16:00:00",
		TotalCoachingCustomerCharging: 0.00,
		TotalCoachingPaymentRate:      0.00,
		TotalTrainigWage:              80000.00,
		TotalOtherWage:                0.00,
		PaymentWage:                   80000.00,
	}, nil)

	mockRepositoryToTimesheet.On("CreateTimesheet", "003", 2019, 12).Return(nil)

	timesheet := Timesheet{
		Repository: mockRepositoryToTimesheet,
	}
	employeeID := "003"
	year := 2019
	month := 12

	actual, err := timesheet.GetSummaryByID(employeeID, year, month)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func Test_GetSummaryByID_Input_EmployeeID_002_Year_2019_Month_12_Should_Be_SummaryTimesheet_No_Incomes_And_Created_Timesheet(t *testing.T) {
	expected := model.SummaryTimesheet{
		EmployeeNameENG:               "Nareenart Narunchon",
		Email:                         "nareenart@scrum123.com",
		RatePerDay:                    0.00,
		RatePerHour:                   0.00,
		Year:                          2019,
		Month:                         12,
		Incomes:                       nil,
		TimesheetID:                   "",
		TotalHours:                    "",
		TotalCoachingCustomerCharging: 0.00,
		TotalCoachingPaymentRate:      0.00,
		TotalTrainigWage:              0.00,
		TotalOtherWage:                0.00,
		PaymentWage:                   0.00,
	}

	mockRepositoryToTimesheet := new(mockinternal.MockRepositoryToTimesheet)
	mockRepositoryToTimesheet.On("GetEmployeeListByEmployeeID", "002").Return([]model.Employee{
		{
			ID:                    3,
			EmployeeID:            "002",
			CompanyID:             2,
			EmployeeNameTH:        "นารีนารถ เนรัญชร",
			EmployeeNameENG:       "Nareenart Narunchon",
			Email:                 "nareenart@scrum123.com",
			RatePerDay:            0.00,
			RatePerHour:           0.00,
			Salary:                25000.00,
			IncomeTax1:            0.00,
			SocialSecurity:        0.00,
			IncomeTax53Percentage: 5,
			Status:                "salary",
			TravelExpense:         0.00,
		},
	}, nil)

	mockRepositoryToTimesheet.On("GetIncomes", "002", 2019, 12).Return([]model.Incomes(nil), nil)
	mockRepositoryToTimesheet.On("GetTimesheet", "002", 2019, 12).Return(model.Timesheet{}, errors.New("sql: no rows in result set"))
	mockRepositoryToTimesheet.On("CreateTimesheet", "002", 2019, 12).Return(nil)
	timesheet := Timesheet{
		Repository: mockRepositoryToTimesheet,
	}
	employeeID := "002"
	year := 2019
	month := 12

	actual, err := timesheet.GetSummaryByID(employeeID, year, month)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func Test_VerifyAuthentication_Input_Email_nareenart_gmail_com_Expiry_1569920400_Should_Be_False(t *testing.T) {
	expected := false
	email := "nareenart@gmail.com"
	expiry := 1569920400.00
	os.Setenv("FIX_TIME", "20181201120000")
	mockRepositoryToTimesheet := new(mockinternal.MockRepositoryToTimesheet)
	mockRepositoryToTimesheet.On("GetEmployeeIDByEmail", mock.Anything).Return("071", nil)
	timesheet := Timesheet{
		Repository: mockRepositoryToTimesheet,
	}

	actual := timesheet.VerifyAuthentication(email, expiry)

	assert.Equal(t, expected, actual)
}

func Test_VerifyAuthentication_Input_Email_nareenart_scrum123_com_ID_Token_Expiration_Time_1569920400_Should_Be_True(t *testing.T) {
	expected := true
	email := "nareenart@scrum123.com"
	expiry := 1569920400.00
	os.Setenv("FIX_TIME", "20181201120000")
	mockRepositoryToTimesheet := new(mockinternal.MockRepositoryToTimesheet)
	mockRepositoryToTimesheet.On("GetEmployeeIDByEmail", mock.Anything).Return("071", nil)
	timesheet := Timesheet{
		Repository: mockRepositoryToTimesheet,
	}

	actual := timesheet.VerifyAuthentication(email, expiry)

	assert.Equal(t, expected, actual)
}

func Test_VerifyAuthentication_Input_Email_nuttaya_c_welovebug_biz_ID_Token_Expiration_Time_1538384400_Should_Be_False_By_Expired(t *testing.T) {
	expected := false
	email := "nuttaya.c@welovebug.biz"
	expiry := 1538384400.00
	os.Setenv("FIX_TIME", "20181201120000")
	mockRepositoryToTimesheet := new(mockinternal.MockRepositoryToTimesheet)
	mockRepositoryToTimesheet.On("GetEmployeeIDByEmail", mock.Anything).Return("071", nil)
	timesheet := Timesheet{
		Repository: mockRepositoryToTimesheet,
	}

	actual := timesheet.VerifyAuthentication(email, expiry)

	assert.Equal(t, expected, actual)
}

func Test_GetSummaryInYearByEmployeeID_Input_EmployeeID_001_Year_2017_Should_Be_SummaryTransactionTimesheet(t *testing.T) {
	expected := model.SummaryTransactionTimesheet{
		EmployeeID: "001",
		Year:       2017,
		TransactionTimesheets: []model.TransactionTimesheet{
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
		},
		TotalCoachingInYear:       85000.00,
		TotalTrainingInYear:       30000.00,
		TotalOtherInYear:          40000.00,
		TotalIncomesInYear:        155000.00,
		TotalSalaryInYear:         80000.00,
		TotalIncomeTax1InYear:     5000.00,
		TotalSocialSecurityInYear: 0.00,
		TotalNetSalaryInYear:      75000.00,
		TotalWageInYear:           75000.00,
		TotalIncomeTax53InYear:    7500.00,
		TotalNetWageInYear:        67500.00,
		TotalNetTransferInYear:    142500.00,
	}

	mockRepositoryToTimesheet := new(mockinternal.MockRepositoryToTimesheet)
	mockRepositoryToTimesheet.On("GetTransactionTimesheets", "001", 2017).Return([]model.TransactionTimesheet{
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
	}, nil)

	timesheet := Timesheet{
		Repository: mockRepositoryToTimesheet,
	}
	employeeID := "001"
	year := 2017

	actual, err := timesheet.GetSummaryInYearByEmployeeID(employeeID, year)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}
