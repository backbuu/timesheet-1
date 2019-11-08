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

func Test_CalculatePaymentSummary_Input_Member_MemberID_001_Should_Be_TransactionTimesheet(t *testing.T) {
	expected := []model.TransactionTimesheet{
		{
			MemberID:              "001",
			MemberNameTH:          "ประธาน ด่านสกุลเจริญกิจ",
			Year:                  2018,
			Month:                 12,
			Company:               "siam_chamnankit",
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
			MemberID:              "001",
			MemberNameTH:          "ประธาน ด่านสกุลเจริญกิจ",
			Year:                  2018,
			Month:                 12,
			Company:               "shuhari",
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
	member := []model.Member{
		{
			MemberID:              "001",
			Company:               "siam_chamnankit",
			MemberNameTH:          "ประธาน ด่านสกุลเจริญกิจ",
			MemberNameENG:         "Prathan Dansakulcharoenkit",
			Email:                 "prathan@scrum123.com",
			OvertimeRate:          0.00,
			RatePerDay:            15000.00,
			RatePerHour:           1875.00,
			Salary:                80000.00,
			IncomeTax1:            5000.00,
			SocialSecurity:        0.00,
			IncomeTax53Percentage: 10,
			TravelExpense:         0.00,
		},
		{
			MemberID:              "001",
			Company:               "shuhari",
			MemberNameTH:          "ประธาน ด่านสกุลเจริญกิจ",
			MemberNameENG:         "Prathan Dansakulcharoenkit",
			Email:                 "prathan@scrum123.com",
			OvertimeRate:          0.00,
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
			Overtime:                 0,
			TotalHours:               totalHours,
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             155000.00,
			OtherWage:                0.00,
			Company:                  "siam_chamnankit",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
		{
			Day:                      3,
			StartTimeAM:              startTimeAM,
			EndTimeAM:                endTimeAM,
			StartTimePM:              startTimePM,
			EndTimePM:                endTimePM,
			Overtime:                 0,
			TotalHours:               totalHours,
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             20000.00,
			OtherWage:                0.00,
			Company:                  "shuhari",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
	}
	year := 2018
	month := 12

	timesheet := Timesheet{}
	actual := timesheet.CalculatePaymentSummary(member, incomes, year, month)

	assert.Equal(t, expected, actual)
}

func Test_CalculatePaymentSummary_Input_Member_MemberID_001_Should_Be_Append_One_TransactionTimesheet(t *testing.T) {
	expected := []model.TransactionTimesheet{
		{
			MemberID:              "001",
			MemberNameTH:          "ประธาน ด่านสกุลเจริญกิจ",
			Year:                  2018,
			Month:                 12,
			Company:               "siam_chamnankit",
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
	member := []model.Member{
		{
			MemberID:              "001",
			Company:               "siam_chamnankit",
			MemberNameTH:          "ประธาน ด่านสกุลเจริญกิจ",
			MemberNameENG:         "Prathan Dansakulcharoenkit",
			Email:                 "prathan@scrum123.com",
			OvertimeRate:          0.00,
			RatePerDay:            15000.00,
			RatePerHour:           1875.00,
			Salary:                80000.00,
			IncomeTax1:            5000.00,
			SocialSecurity:        0.00,
			IncomeTax53Percentage: 10,
			TravelExpense:         0.00,
		},
		{
			MemberID:              "001",
			Company:               "shuhari",
			MemberNameTH:          "ประธาน ด่านสกุลเจริญกิจ",
			MemberNameENG:         "Prathan Dansakulcharoenkit",
			Email:                 "prathan@scrum123.com",
			OvertimeRate:          0.00,
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
			Overtime:                 0,
			TotalHours:               totalHours,
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             155000.00,
			OtherWage:                0.00,
			Company:                  "siam_chamnankit",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
	}
	year := 2018
	month := 12

	timesheet := Timesheet{}
	actual := timesheet.CalculatePaymentSummary(member, incomes, year, month)

	assert.Equal(t, expected, actual)
}

func Test_CalculatePaymentSummary_Input_Member_MemberID_002_Should_Be_TransactionTimesheet(t *testing.T) {
	expected := []model.TransactionTimesheet{
		{
			MemberID:              "002",
			MemberNameTH:          "นารีนารถ เนรัญชร",
			Year:                  2018,
			Month:                 12,
			Company:               "shuhari",
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
	member := []model.Member{
		{
			MemberID:              "002",
			Company:               "shuhari",
			MemberNameTH:          "นารีนารถ เนรัญชร",
			MemberNameENG:         "Nareenart Narunchon",
			Email:                 "nareenart@scrum123.com",
			OvertimeRate:          0.00,
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
			Overtime:                 0,
			TotalHours:               totalHours,
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             0.00,
			OtherWage:                5000.00,
			Company:                  "shuhari",
			Description:              "work at TN",
		},
	}
	year := 2018
	month := 12

	timesheet := Timesheet{}
	actual := timesheet.CalculatePaymentSummary(member, incomes, year, month)

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
			Overtime:                 0,
			TotalHours:               totalHours,
			CoachingCustomerCharging: 15000.00,
			CoachingPaymentRate:      10000.00,
			TrainingWage:             0.00,
			OtherWage:                0.00,
			Company:                  "siam_chamnankit",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
		{
			Day:                      3,
			StartTimeAM:              startTimeAM,
			EndTimeAM:                endTimeAM,
			StartTimePM:              startTimePM,
			EndTimePM:                endTimePM,
			Overtime:                 0,
			TotalHours:               totalHours,
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             20000.00,
			OtherWage:                0.00,
			Company:                  "shuhari",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
	}
	timesheet := Timesheet{}

	actual := timesheet.CalculatePayment(incomes)

	assert.Equal(t, expected, actual)
}

func Test_GetSummaryByID_Input_MemberID_003_Year_2019_Month_12_Should_Be_SummaryTimesheet(t *testing.T) {
	startTimeAM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 09:00:00")
	endTimeAM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 12:00:00")
	startTimePM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 13:00:00")
	endTimePM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 18:00:00")
	totalHours, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 08:00:00")
	expected := model.SummaryTimesheet{
		MemberNameENG: "Somkiat Puisungnoen",
		Email:         "somkiat@scrum123.com",
		OvertimeRate:  0.00,
		RatePerDay:    15000.00,
		RatePerHour:   1875.00,
		Year:          2019,
		Month:         12,
		Incomes: []model.Incomes{
			{
				ID:                       61,
				MemberID:                 "003",
				Month:                    12,
				Year:                     2019,
				Day:                      1,
				StartTimeAM:              startTimeAM,
				EndTimeAM:                endTimeAM,
				StartTimePM:              startTimePM,
				EndTimePM:                endTimePM,
				Overtime:                 0,
				TotalHours:               totalHours,
				CoachingCustomerCharging: 0.00,
				CoachingPaymentRate:      0.00,
				TrainingWage:             40000.00,
				OtherWage:                0.00,
				Company:                  "shuhari",
				Description:              "Technical Excellence at Khonkean",
			},
			{
				ID:                       62,
				MemberID:                 "003",
				Month:                    12,
				Year:                     2019,
				Day:                      2,
				StartTimeAM:              startTimeAM,
				EndTimeAM:                endTimeAM,
				StartTimePM:              startTimePM,
				EndTimePM:                endTimePM,
				Overtime:                 0,
				TotalHours:               totalHours,
				CoachingCustomerCharging: 0.00,
				CoachingPaymentRate:      0.00,
				TrainingWage:             40000.00,
				OtherWage:                0.00,
				Company:                  "shuhari",
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
	mockRepository := new(mockinternal.MockRepository)
	mockRepository.On("GetMemberListByMemberID", "003").Return([]model.Member{
		{
			ID:                    4,
			MemberID:              "003",
			Company:               "siam_chamnankit",
			MemberNameTH:          "สมเกียรติ ปุ๋ยสูงเนิน",
			MemberNameENG:         "Somkiat Puisungnoen",
			Email:                 "somkiat@scrum123.com",
			OvertimeRate:          0.00,
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
			MemberID:              "003",
			Company:               "shuhari",
			MemberNameTH:          "สมเกียรติ ปุ๋ยสูงเนิน",
			MemberNameENG:         "Somkiat Puisungnoen",
			Email:                 "somkiat@scrum123.com",
			OvertimeRate:          0.00,
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

	mockRepository.On("GetIncomes", "003", 2019, 12).Return([]model.Incomes{
		{
			ID:                       61,
			MemberID:                 "003",
			Month:                    12,
			Year:                     2019,
			Day:                      1,
			StartTimeAM:              startTimeAM,
			EndTimeAM:                endTimeAM,
			StartTimePM:              startTimePM,
			EndTimePM:                endTimePM,
			Overtime:                 0,
			TotalHours:               totalHours,
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             40000.00,
			OtherWage:                0.00,
			Company:                  "shuhari",
			Description:              "Technical Excellence at Khonkean",
		},
		{
			ID:                       62,
			MemberID:                 "003",
			Month:                    12,
			Year:                     2019,
			Day:                      2,
			StartTimeAM:              startTimeAM,
			EndTimeAM:                endTimeAM,
			StartTimePM:              startTimePM,
			EndTimePM:                endTimePM,
			Overtime:                 0,
			TotalHours:               totalHours,
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             40000.00,
			OtherWage:                0.00,
			Company:                  "shuhari",
			Description:              "Technical Excellence at Khonkean",
		},
	}, nil)

	mockRepository.On("GetTimesheet", "003", 2019, 12).Return(model.Timesheet{
		ID:                            "003201912",
		MemberID:                      "003",
		Month:                         12,
		Year:                          2019,
		TotalHours:                    "16:00:00",
		TotalCoachingCustomerCharging: 0.00,
		TotalCoachingPaymentRate:      0.00,
		TotalTrainigWage:              80000.00,
		TotalOtherWage:                0.00,
		PaymentWage:                   80000.00,
	}, nil)

	mockRepository.On("CreateTimesheet", "003", 2019, 12).Return(nil)

	timesheet := Timesheet{
		Repository: mockRepository,
	}
	memberID := "003"
	year := 2019
	month := 12

	actual, err := timesheet.GetSummaryByID(memberID, year, month)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func Test_GetSummaryByID_Input_MemberID_002_Year_2019_Month_12_Should_Be_SummaryTimesheet_No_Incomes_And_Created_Timesheet(t *testing.T) {
	expected := model.SummaryTimesheet{
		MemberNameENG:                 "Nareenart Narunchon",
		Email:                         "nareenart@scrum123.com",
		OvertimeRate:                  0.00,
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

	mockRepository := new(mockinternal.MockRepository)
	mockRepository.On("GetMemberListByMemberID", "002").Return([]model.Member{
		{
			ID:                    3,
			MemberID:              "002",
			Company:               "shuhari",
			MemberNameTH:          "นารีนารถ เนรัญชร",
			MemberNameENG:         "Nareenart Narunchon",
			Email:                 "nareenart@scrum123.com",
			OvertimeRate:          0.00,
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

	mockRepository.On("GetIncomes", "002", 2019, 12).Return([]model.Incomes(nil), nil)
	mockRepository.On("GetTimesheet", "002", 2019, 12).Return(model.Timesheet{}, errors.New("sql: no rows in result set"))
	mockRepository.On("CreateTimesheet", "002", 2019, 12).Return(nil)
	timesheet := Timesheet{
		Repository: mockRepository,
	}
	memberID := "002"
	year := 2019
	month := 12

	actual, err := timesheet.GetSummaryByID(memberID, year, month)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func Test_VerifyAuthentication_Input_AccessToken_MemberID_071_Should_Be_Unauthorized_By_Expired(t *testing.T) {
	expected := "Unauthorized"
	accessToken := "ya29.Il-vB2mB0hkAEN8KdupS3ZEaXBOHk6qhVntGSkeyAMz6KEoJOpwhfHHQF2KT9W2oiwE1op4pZiUuebKcQ1SBRgRlxMRJxB6Qjf0tl86C5Jdsf51thN-yqvZDBUmUx3hnqw"
	memberID := "071"
	os.Setenv("FIX_TIME", "20181201120000")
	expiry, _ := time.Parse("20060102150405", "20181201090000")
	mockRepository := new(mockinternal.MockRepository)
	mockRepository.On("GetVerifyAuthenticationByAccessToken", mock.Anything).Return(model.VerifyAuthentication{
		MemberID: "071",
		Expiry:   expiry,
	}, nil)
	timesheet := Timesheet{
		Repository: mockRepository,
	}

	actual := timesheet.VerifyAuthentication(accessToken, memberID)

	assert.Equal(t, expected, actual)
}

func Test_VerifyAuthentication_Input_AccessToken_MemberID_071_Should_Be_Unauthorized(t *testing.T) {
	expected := "Unauthorized"
	accessToken := "ya29.Il-vB2mB0hkAEN8KdupS3ZEaXBOHk6qhVntGSkeyAMz6KEoJOpwhfHHQF2KT9W2oiwE1op4pZiUuebKcQ1SBRgRlxMRJxB6Qjf0tl86C5Jdsf51thN-yqvZDBUmUx3hnqw"
	memberID := "008"
	os.Setenv("FIX_TIME", "20181201120000")
	expiry, _ := time.Parse("20060102150405", "20181201090000")
	mockRepository := new(mockinternal.MockRepository)
	mockRepository.On("GetVerifyAuthenticationByAccessToken", mock.Anything).Return(model.VerifyAuthentication{
		MemberID: "071",
		Expiry:   expiry,
	}, nil)
	timesheet := Timesheet{
		Repository: mockRepository,
	}

	actual := timesheet.VerifyAuthentication(accessToken, memberID)

	assert.Equal(t, expected, actual)
}

func Test_VerifyAuthentication_Input_AccessToken_MemberID_071_Should_Be_Success(t *testing.T) {
	expected := "Success"
	accessToken := "ya29.Il-vB2mB0hkAEN8KdupS3ZEaXBOHk6qhVntGSkeyAMz6KEoJOpwhfHHQF2KT9W2oiwE1op4pZiUuebKcQ1SBRgRlxMRJxB6Qjf0tl86C5Jdsf51thN-yqvZDBUmUx3hnqw"
	memberID := "071"
	os.Setenv("FIX_TIME", "20181201083000")
	expiry, _ := time.Parse("20060102150405", "20181201090000")
	mockRepository := new(mockinternal.MockRepository)
	mockRepository.On("GetVerifyAuthenticationByAccessToken", mock.Anything).Return(model.VerifyAuthentication{
		MemberID: "071",
		Expiry:   expiry,
	}, nil)
	timesheet := Timesheet{
		Repository: mockRepository,
	}

	actual := timesheet.VerifyAuthentication(accessToken, memberID)

	assert.Equal(t, expected, actual)
}
