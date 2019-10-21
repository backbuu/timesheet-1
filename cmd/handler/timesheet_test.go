package handler_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	. "timesheet/cmd/handler"
	"timesheet/cmd/mockapi"
	"timesheet/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_GetSummaryByIDHandler_Input_Year_2019_Month_12_MemberID_003_Should_Be_Timesheet(t *testing.T) {
	startTimeAM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 09:00:00")
	endTimeAM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 12:00:00")
	startTimePM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 13:00:00")
	endTimePM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 18:00:00")
	totalHours, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 08:00:00")
	expected := `{"member_name_eng":"Somkiat Puisungnoen", "email":"somkiat@scrum123.com", "overtime_rate":0, "rate_per_day":15000.00, "rate_per_hour":1875.00, "month":12,"year":2018, "incomes":[{"day":18, "start_time_am":"2018-12-01 09:00:00", "end_time_am":"2018-12-01 12:00:00", "start_time_pm":"2018-12-01 13:00:00", "end_time_pm":"2018-12-01 18:00:00", "overtime":0, "total_hours":8, "coaching_customer_charging":0.00, "coaching_payment_rate":0.00, "training_wage":40000.00, "other_wage":0.00, "company":"siam_chamnankit", "description":"TDD with JAVA"},{"day":19, "start_time_am":"2018-12-01 09:00:00", "end_time_am":"2018-12-01 12:00:00", "start_time_pm":"2018-12-01 13:00:00", "end_time_pm":"2018-12-01 18:00:00", "overtime":0, "total_hours":8, "coaching_customer_charging":0.00, "coaching_payment_rate":0.00, "training_wage":40000.00, "other_wage":0.00, "company":"siam_chamnankit", "description":"TDD with JAVA"}]}`
	timesheetRequest := TimesheetRequest{
		Year:     2018,
		Month:    12,
		MemberID: "003",
	}
	jsonRequest, _ := json.Marshal(timesheetRequest)
	request := httptest.NewRequest("POST", "/showTimesheetByID", bytes.NewBuffer(jsonRequest))
	writer := httptest.NewRecorder()

	mockRepository := new(mockapi.MockRepository)
	mockRepository.On("GetTimesheet", 2018, 12).Return(model.Timesheet{
		MemberNameENG:                 "Somkiat Puisungnoen",
		Email:                         "somkiat@scrum123.com",
		OvertimeRate:                  0.00,
		RatePerDay:                    15000.00,
		RatePerHour:                   1875.00,
		Month:                         12,
		Year:                          2019,
		TotalHoursPerMonth:            16,
		TotalCoachingCustomerCharging: 0.00,
		TotalCoachingPaymentRate:      0.00,
		TotalTrainigWage:              80000.00,
		TotalOtherWage:                0.00,
		PaymentWage:                   80000.00,
	}, nil)

	mockRepository := new(mockapi.MockRepository)
	mockRepository.On("GetIncomes", "003", 2019, 12).Return([]model.Incomes{
		{
			Day:                      18,
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
			Company:                  "siam_chamnankit",
			Description:              "TDD with JAVA",
		},
		{
			Day:                      19,
			StartTimeAM:              startTimeAM,
			EndTimeAM:                endTimeAM,
			StartTimePM:              startTimePM,
			EndTimePM:                endTimePM,
			Overtime:                 0,
			TotalHours:               totalHours,
			TotalHoursHours:          8,
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             40000.00,
			OtherWage:                0.00,
			Company:                  "siam_chamnankit",
			Description:              "TDD with JAVA",
		},
	}, nil)

	api := TimesheetAPI{
		TimesheetRepository: mockRepository,
	}
	testRoute := gin.Default()
	testRoute.POST("/showSummaryTimesheetByID", api.GetSummaryByIDHandler)
	testRoute.ServeHTTP(writer, request)

	response := writer.Result()
	actual, err := ioutil.ReadAll(response.Body)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, string(actual))
}

func Test_GetSummaryHandler_Input_Year_2018_Month_12_Should_Be_Timesheet(t *testing.T) {
	expected := `[{"id":"001201812siam_chamnankit","member_id":"001","member_name_th":"ประธาน ด่านสกุลเจริญกิจ","month":12,"year":2018,"company":"siam_chamnankit","coaching":85000,"training":30000,"other":40000,"total_incomes":155000,"salary":80000,"income_tax_1":5000,"social_security":0,"net_salary":75000,"wage":75000,"income_tax_53_percentage":10,"income_tax_53":7500,"net_wage":67500,"net_transfer":142500,"status_checking_transfer":"รอการตรวจสอบ","date_transfer":null,"comment":null},{"id":"001201812shuhari","member_id":"001","member_name_th":"ประธาน ด่านสกุลเจริญกิจ","month":12,"year":2018,"company":"shuhari","coaching":0,"training":40000,"other":0,"total_incomes":40000,"salary":0,"income_tax_1":0,"social_security":0,"net_salary":0,"wage":40000,"income_tax_53_percentage":10,"income_tax_53":4000,"net_wage":36000,"net_transfer":36000,"status_checking_transfer":"รอการตรวจสอบ","date_transfer":null,"comment":null}]`
	date := Date{
		Year:  2018,
		Month: 12,
	}
	jsonRequest, _ := json.Marshal(date)
	request := httptest.NewRequest("POST", "/showSummaryTimesheet", bytes.NewBuffer(jsonRequest))
	writer := httptest.NewRecorder()

	mockRepository := new(mockapi.MockRepository)
	mockRepository.On("GetSummary", 2018, 12).Return([]model.TransactionTimesheet{
		{
			ID:                     "001201812siam_chamnankit",
			MemberID:               "001",
			MemberNameTH:           "ประธาน ด่านสกุลเจริญกิจ",
			Month:                  12,
			Year:                   2018,
			Company:                "siam_chamnankit",
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
		}, {
			ID:                     "001201812shuhari",
			MemberID:               "001",
			MemberNameTH:           "ประธาน ด่านสกุลเจริญกิจ",
			Month:                  12,
			Year:                   2018,
			Company:                "shuhari",
			Coaching:               0.00,
			Training:               40000.00,
			Other:                  0.00,
			TotalIncomes:           40000.00,
			Salary:                 0.00,
			IncomeTax1:             0.00,
			SocialSecurity:         0.00,
			NetSalary:              0.00,
			Wage:                   40000.00,
			IncomeTax53Percentage:  10,
			IncomeTax53:            4000.00,
			NetWage:                36000.00,
			NetTransfer:            36000.00,
			StatusCheckingTransfer: "รอการตรวจสอบ",
		},
	}, nil)

	api := TimesheetAPI{
		TimesheetRepository: mockRepository,
	}
	testRoute := gin.Default()
	testRoute.POST("/showSummaryTimesheet", api.GetSummaryHandler)
	testRoute.ServeHTTP(writer, request)

	response := writer.Result()
	actual, err := ioutil.ReadAll(response.Body)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, string(actual))
}

func Test_CreateIncomeHandler_Input_Year_2018_Month_12_MemberID_001_Income_Should_Be_Status_200(t *testing.T) {
	requestIncome := IncomeRequest{
		Year:     2018,
		Month:    12,
		MemberID: "001",
		Incomes: model.Incomes{
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
	}
	jsonRequest, _ := json.Marshal(requestIncome)
	request := httptest.NewRequest("POST", "/addIncomeItem", bytes.NewBuffer(jsonRequest))
	writer := httptest.NewRecorder()

	mockRepository := new(mockapi.MockRepository)
	mockRepository.On("CreateIncome", 2018, 12, "001", model.Incomes{
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
		CoachingCustomerCharging: 15000.00,
		CoachingPaymentRate:      10000.00,
		TrainingWage:             0.00,
		OtherWage:                0.00,
		Company:                  "siam_chamnankit",
		Description:              "[KBTG] 2 Days Agile Project Management",
	}).Return(nil)

	api := TimesheetAPI{
		TimesheetRepository: mockRepository,
	}

	testRoute := gin.Default()
	testRoute.POST("/addIncomeItem", api.CreateIncomeHandler)
	testRoute.ServeHTTP(writer, request)

	response := writer.Result()

	assert.Equal(t, http.StatusCreated, response.StatusCode)
}

func Test_CalculatePaymentHandler_Input_MemberID_001_Year_2018_Month_12_Should_Be_200(t *testing.T) {
	expectedStatus := http.StatusOK
	calculatePaymentRequest := CalculatePaymentRequest{
		MemberID: "001",
		Year:     2018,
		Month:    12,
	}
	jsonRequest, _ := json.Marshal(calculatePaymentRequest)
	request := httptest.NewRequest("POST", "/calculatePayment", bytes.NewBuffer(jsonRequest))
	writer := httptest.NewRecorder()

	mockRepository := new(mockapi.MockRepository)
	mockRepository.On("GetIncomes", "001", 2018, 12).Return([]model.Incomes{
		{
			Day:                      19,
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
			TrainingWage:             40000.00,
			OtherWage:                0.00,
			Company:                  "shuhari",
			Description:              "[IMC]GSB: Agile Project Mgmt",
		},
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
			CoachingCustomerCharging: 130000.00,
			CoachingPaymentRate:      85000.00,
			TrainingWage:             30000.00,
			OtherWage:                40000.00,
			Company:                  "siam_chamnankit",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
	}, nil)

	mockTimesheet := new(mockapi.MockTimesheet)
	mockTimesheet.On("CalculatePayment", mock.Anything).Return(model.Payment{
		TotalHoursHours:               16,
		TotalHoursMinutes:             0,
		TotalHoursSeconds:             0,
		TotalCoachingCustomerCharging: 130000.00,
		TotalCoachingPaymentRate:      85000.00,
		TotalTrainigWage:              70000.00,
		TotalOtherWage:                40000.00,
		PaymentWage:                   195000.00,
	})

	mockRepository.On("VerifyTimesheet", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	mockRepository.On("GetMemberByID", "001").Return([]model.Member{
		{
			ID:                    1,
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
			ID:                    2,
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
	}, nil)

	mockTimesheet.On("CalculatePaymentSummary", mock.Anything, mock.Anything, 2018, 12).Return([]model.TransactionTimesheet{
		{
			MemberID:              "001",
			MemberNameTH:          "ประธาน ด่านสกุลเจริญกิจ",
			Month:                 12,
			Year:                  2018,
			Company:               "siam_chamnankit",
			Coaching:              85000.00,
			Training:              30000.00,
			Other:                 40000.00,
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
		}, {
			MemberID:              "001",
			MemberNameTH:          "ประธาน ด่านสกุลเจริญกิจ",
			Month:                 12,
			Year:                  2018,
			Company:               "shuhari",
			Coaching:              0.00,
			Training:              40000.00,
			Other:                 0.00,
			TotalIncomes:          40000.00,
			Salary:                0.00,
			IncomeTax1:            0.00,
			SocialSecurity:        0.00,
			NetSalary:             0.00,
			Wage:                  40000.00,
			IncomeTax53Percentage: 10,
			IncomeTax53:           4000.00,
			NetWage:               36000.00,
			NetTransfer:           36000.00,
		},
	})

	mockRepository.On("VerifyTransactionTimsheet", mock.Anything).Return(nil)

	api := TimesheetAPI{
		Timesheet:           mockTimesheet,
		TimesheetRepository: mockRepository,
	}

	testRoute := gin.Default()
	testRoute.POST("/calculatePayment", api.CalculatePaymentHandler)
	testRoute.ServeHTTP(writer, request)

	actual := writer.Result()

	assert.Equal(t, expectedStatus, actual.StatusCode)
}
