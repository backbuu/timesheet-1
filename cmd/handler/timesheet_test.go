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

func Test_GetSummaryByIDHandler_Input_Year_2018_Month_12_MemberID_003_Should_Be_Timesheet(t *testing.T) {
	startTimeAM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 09:00:00")
	endTimeAM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 12:00:00")
	startTimePM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 13:00:00")
	endTimePM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 18:00:00")
	totalHours, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 08:00:00")
	expected := `{"member_name_eng":"Somkiat Puisungnoen","email":"somkiat@scrum123.com","overtime_rate":0,"rate_per_day":15000,"rate_per_hour":1875,"year":2019,"month":12,"incomes":[{"id":61,"member_id":"003","month":12,"year":2019,"day":1,"start_time_am":"2018-12-01T09:00:00Z","end_time_am":"2018-12-01T12:00:00Z","start_time_pm":"2018-12-01T13:00:00Z","end_time_pm":"2018-12-01T18:00:00Z","overtime":0,"total_hours":"2018-12-01T08:00:00Z","coaching_customer_charging":0,"coaching_payment_rate":0,"training_wage":40000,"other_wage":0,"company":"shuhari","description":"Technical Excellence at Khonkean"},{"id":62,"member_id":"003","month":12,"year":2019,"day":2,"start_time_am":"2018-12-01T09:00:00Z","end_time_am":"2018-12-01T12:00:00Z","start_time_pm":"2018-12-01T13:00:00Z","end_time_pm":"2018-12-01T18:00:00Z","overtime":0,"total_hours":"2018-12-01T08:00:00Z","coaching_customer_charging":0,"coaching_payment_rate":0,"training_wage":40000,"other_wage":0,"company":"shuhari","description":"Technical Excellence at Khonkean"}],"timesheet_id":"003201912","total_hours":"16:00:00","total_coaching_customer_charging":0,"total_coaching_payment_rate":0,"total_training_wage":80000,"total_other_wage":0,"payment_wage":80000}`
	timesheetRequest := TimesheetRequest{
		Year:     2017,
		Month:    12,
		MemberID: "003",
	}
	jsonRequest, _ := json.Marshal(timesheetRequest)
	request := httptest.NewRequest("POST", "/showTimesheetByID", bytes.NewBuffer(jsonRequest))
	writer := httptest.NewRecorder()

	mockTimesheet := new(mockapi.MockTimesheet)
	mockTimesheet.On("GetSummaryByID", "003", 2017, 12).Return(model.SummaryTimesheet{
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
	}, nil)
	api := TimesheetAPI{
		Timesheet: mockTimesheet,
	}
	testRoute := gin.Default()
	testRoute.POST("/showTimesheetByID", api.GetSummaryByIDHandler)
	testRoute.ServeHTTP(writer, request)

	response := writer.Result()
	actual, err := ioutil.ReadAll(response.Body)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, string(actual))
}

func Test_GetSummaryHandler_Input_Year_2018_Month_12_Should_Be_Timesheet(t *testing.T) {
	expected := `[{"id":"001201812siam_chamnankit","member_id":"001","member_name_th":"ประธาน ด่านสกุลเจริญกิจ","month":12,"year":2018,"company":"siam_chamnankit","coaching":85000,"training":30000,"other":40000,"total_incomes":155000,"salary":80000,"income_tax_1":5000,"social_security":0,"net_salary":75000,"wage":75000,"income_tax_53_percentage":10,"income_tax_53":7500,"net_wage":67500,"net_transfer":142500,"status_checking_transfer":"รอการตรวจสอบ","date_transfer":"","comment":""},{"id":"001201812shuhari","member_id":"001","member_name_th":"ประธาน ด่านสกุลเจริญกิจ","month":12,"year":2018,"company":"shuhari","coaching":0,"training":40000,"other":0,"total_incomes":40000,"salary":0,"income_tax_1":0,"social_security":0,"net_salary":0,"wage":40000,"income_tax_53_percentage":10,"income_tax_53":4000,"net_wage":36000,"net_transfer":36000,"status_checking_transfer":"รอการตรวจสอบ","date_transfer":"","comment":""}]`
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
	startTimeAM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 09:00:00")
	endTimeAM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 12:00:00")
	startTimePM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 13:00:00")
	endTimePM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 18:00:00")
	totalHours, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 08:00:00")
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
	calculatePaymentRequest := CalculatePaymentRequest{
		MemberID: "001",
		Year:     2018,
		Month:    12,
	}
	jsonRequest, _ := json.Marshal(calculatePaymentRequest)
	request := httptest.NewRequest("POST", "/calculatePayment", bytes.NewBuffer(jsonRequest))
	writer := httptest.NewRecorder()

	startTimeAM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 09:00:00")
	endTimeAM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 12:00:00")
	startTimePM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 13:00:00")
	endTimePM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 18:00:00")
	totalHours, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 08:00:00")
	mockRepository := new(mockapi.MockRepository)
	mockRepository.On("GetIncomes", "001", 2018, 12).Return([]model.Incomes{
		{
			Day:                      19,
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
			Description:              "[IMC]GSB: Agile Project Mgmt",
		},
		{
			Day:                      28,
			StartTimeAM:              startTimeAM,
			EndTimeAM:                endTimeAM,
			StartTimePM:              startTimePM,
			EndTimePM:                endTimePM,
			Overtime:                 0,
			TotalHours:               totalHours,
			CoachingCustomerCharging: 130000.00,
			CoachingPaymentRate:      85000.00,
			TrainingWage:             30000.00,
			OtherWage:                40000.00,
			Company:                  "siam_chamnankit",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
	}, nil)

	mockTimesheet := new(mockapi.MockTimesheet)
	mockTimesheet.On("CalculatePayment", mock.Anything).Return(model.Timesheet{
		TotalHours:                    "16:00:00",
		TotalCoachingCustomerCharging: 130000.00,
		TotalCoachingPaymentRate:      85000.00,
		TotalTrainigWage:              70000.00,
		TotalOtherWage:                40000.00,
		PaymentWage:                   195000.00,
	})

	mockRepository.On("UpdateTimesheet", mock.Anything, "001", 2018, 12).Return(nil)

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

	assert.Equal(t, http.StatusOK, actual.StatusCode)
}

func Test_UpdateStatusCheckingTransferHandler_Input_TransactionID_004201912siam_chamnankit_Status_TransferSuccess_Date_30_12_2019_Comment_FlightTicket_Should_Be_Status_200(t *testing.T) {
	requestUpdate := UpdateStatusRequest{
		TransactionID: "004201912siam_chamnankit",
		Status:        "โอนเงินเรียบร้อยแล้ว",
		Date:          "30/12/2019",
		Comment:       "หักค่าตั๋วเครื่องบิน",
	}

	jsonRequest, _ := json.Marshal(requestUpdate)
	request := httptest.NewRequest("POST", "/updateStatusCheckingTransfer", bytes.NewBuffer(jsonRequest))
	writer := httptest.NewRecorder()

	mockRepository := new(mockapi.MockRepository)
	mockRepository.On("UpdateStatusTransfer", "004201912siam_chamnankit", "โอนเงินเรียบร้อยแล้ว", "30/12/2019", "หักค่าตั๋วเครื่องบิน").Return(nil)

	api := TimesheetAPI{
		TimesheetRepository: mockRepository,
	}

	testRoute := gin.Default()
	testRoute.POST("/updateStatusCheckingTransfer", api.UpdateStatusCheckingTransferHandler)
	testRoute.ServeHTTP(writer, request)

	response := writer.Result()

	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func Test_DeleteIncomeHandler_Input_ID_47_Should_Be_Status_200(t *testing.T) {
	requestDelete := DeleteIncomeRequest{
		ID: 47,
	}

	jsonRequest, _ := json.Marshal(requestDelete)
	request := httptest.NewRequest("POST", "/deleteIncomeItem", bytes.NewBuffer(jsonRequest))
	writer := httptest.NewRecorder()

	mockRepository := new(mockapi.MockRepository)
	mockRepository.On("DeleteIncome", 47).Return(nil)

	api := TimesheetAPI{
		TimesheetRepository: mockRepository,
	}

	testRoute := gin.Default()
	testRoute.POST("/deleteIncomeItem", api.DeleteIncomeHandler)
	testRoute.ServeHTTP(writer, request)

	response := writer.Result()

	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func Test_ShowMemberDetailsByIDHandler_Input_MemberID_001_Should_Be_MemberDetails(t *testing.T) {
	expected := `[{"id":1,"member_id":"001","company":"siam_chamnankit","member_name_th":"ประธาน ด่านสกุลเจริญกิจ","member_name_eng":"Prathan Dansakulcharoenkit","email":"prathan@scrum123.com","overtime_rate":0,"rate_per_day":15000,"rate_per_hour":1875,"salary":80000,"income_tax_1":5000,"social_security":0,"income_tax_53_percentage":10,"status":"wage","travel_expense":0},{"id":2,"member_id":"001","company":"shuhari","member_name_th":"ประธาน ด่านสกุลเจริญกิจ","member_name_eng":"Prathan Dansakulcharoenkit","email":"prathan@scrum123.com","overtime_rate":0,"rate_per_day":15000,"rate_per_hour":1875,"salary":0,"income_tax_1":0,"social_security":0,"income_tax_53_percentage":10,"status":"wage","travel_expense":0}]`
	memberRequest := MemberRequest{
		MemberID: "001",
	}

	jsonRequest, _ := json.Marshal(memberRequest)
	request := httptest.NewRequest("POST", "/showMemberDetailsByID", bytes.NewBuffer(jsonRequest))
	writer := httptest.NewRecorder()

	mockRepository := new(mockapi.MockRepository)
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
			Status:                "wage",
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
			Status:                "wage",
			TravelExpense:         0.00,
		},
	}, nil)

	api := TimesheetAPI{
		TimesheetRepository: mockRepository,
	}

	testRoute := gin.Default()
	testRoute.POST("/showMemberDetailsByID", api.ShowMemberDetailsByIDHandler)
	testRoute.ServeHTTP(writer, request)

	response := writer.Result()
	actual, err := ioutil.ReadAll(response.Body)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, string(actual))
}

func Test_UpdateMemberDetailsHandler_Input_Member_Should_Be_Status_200(t *testing.T) {
	requestUpdateMember := model.Member{
		ID:                    1,
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
		Status:                "wage",
		TravelExpense:         0.00,
	}

	jsonRequest, _ := json.Marshal(requestUpdateMember)
	request := httptest.NewRequest("POST", "/updateMemberDetails", bytes.NewBuffer(jsonRequest))
	writer := httptest.NewRecorder()

	mockRepository := new(mockapi.MockRepository)
	mockRepository.On("UpdateMemberDetails", model.Member{
		ID:                    1,
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
		Status:                "wage",
		TravelExpense:         0.00,
	}).Return(nil)

	api := TimesheetAPI{
		TimesheetRepository: mockRepository,
	}

	testRoute := gin.Default()
	testRoute.POST("/updateMemberDetails", api.UpdateMemberDetailsHandler)
	testRoute.ServeHTTP(writer, request)

	response := writer.Result()

	assert.Equal(t, http.StatusOK, response.StatusCode)
}
