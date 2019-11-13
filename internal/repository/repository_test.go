package repository_test

import (
	"testing"
	"time"
	"timesheet/internal/model"
	. "timesheet/internal/repository"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/stretchr/testify/assert"
)

func Test_GetSummary_Input_Year_2017_Month_12_Should_Be_TransactionTimesheet(t *testing.T) {
	expected := []model.TransactionTimesheet{
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
	}
	databaseConnection, _ := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/timesheet")
	defer databaseConnection.Close()
	repository := TimesheetRepository{
		DatabaseConnection: databaseConnection,
	}
	year := 2017
	month := 12

	actual, err := repository.GetSummary(year, month)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func Test_CreateIncome_Input_Year_2017_Month_12_EmployeeID_001_Income_Should_Be_No_Error(t *testing.T) {
	startTimeAM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 09:00:00")
	endTimeAM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 12:00:00")
	startTimePM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 13:00:00")
	endTimePM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 18:00:00")
	totalHours, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 08:00:00")
	year := 2017
	month := 12
	employeeID := "001"
	income := model.Incomes{
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
	}
	databaseConnection, _ := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/timesheet")
	defer databaseConnection.Close()
	repository := TimesheetRepository{
		DatabaseConnection: databaseConnection,
	}

	err := repository.CreateIncome(year, month, employeeID, income)

	assert.Equal(t, nil, err)
}

func Test_GetEmployeeListByEmployeeID_Input_EmployeeID_001_Should_Be_EmployeeList(t *testing.T) {
	expected := []model.Employee{
		{
			ID:                    1,
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
			Status:                "wage",
			Picture:               "",
		},
		{
			ID:                    2,
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
			Status:                "wage",
			Picture:               "",
		},
	}
	employeeID := "001"
	databaseConnection, _ := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/timesheet")
	defer databaseConnection.Close()
	repository := TimesheetRepository{
		DatabaseConnection: databaseConnection,
	}

	actual, err := repository.GetEmployeeListByEmployeeID(employeeID)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func Test_GetIncomes_Input_EmployeeID_006_Year_2019_Month_12_Should_Be_IncomeList(t *testing.T) {
	startTimeAMDay11, _ := time.Parse("2006-01-02 15:04:05", "2018-12-11 09:00:00")
	endTimeAMDay11, _ := time.Parse("2006-01-02 15:04:05", "2018-12-11 12:00:00")
	startTimePMDay11, _ := time.Parse("2006-01-02 15:04:05", "2018-12-11 13:00:00")
	endTimePMDay11, _ := time.Parse("2006-01-02 15:04:05", "2018-12-11 18:00:00")
	totalHoursDay11, _ := time.Parse("2006-01-02 15:04:05", "2018-12-11 08:00:00")

	startTimeAMDay12, _ := time.Parse("2006-01-02 15:04:05", "2018-12-12 09:00:00")
	endTimeAMDay12, _ := time.Parse("2006-01-02 15:04:05", "2018-12-12 12:00:00")
	startTimePMDay12, _ := time.Parse("2006-01-02 15:04:05", "2018-12-12 13:00:00")
	endTimePMDay12, _ := time.Parse("2006-01-02 15:04:05", "2018-12-12 18:00:00")
	totalHoursDay12, _ := time.Parse("2006-01-02 15:04:05", "2018-12-12 08:00:00")
	expected := []model.Incomes{
		{
			ID:                       58,
			EmployeeID:               "006",
			Month:                    12,
			Year:                     2019,
			Day:                      11,
			StartTimeAM:              startTimeAMDay11,
			EndTimeAM:                endTimeAMDay11,
			StartTimePM:              startTimePMDay11,
			EndTimePM:                endTimePMDay11,
			TotalHours:               totalHoursDay11,
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             0.00,
			OtherWage:                0.00,
			CompanyID:                2,
			Description:              "work at TN",
		}, {
			ID:                       59,
			EmployeeID:               "006",
			Month:                    12,
			Year:                     2019,
			Day:                      12,
			StartTimeAM:              startTimeAMDay12,
			EndTimeAM:                endTimeAMDay12,
			StartTimePM:              startTimePMDay12,
			EndTimePM:                endTimePMDay12,
			TotalHours:               totalHoursDay12,
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             0.00,
			OtherWage:                0.00,
			CompanyID:                2,
			Description:              "work at TN",
		},
	}
	employeeID := "006"
	month := 12
	year := 2019
	databaseConnection, _ := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/timesheet?parseTime=true")
	defer databaseConnection.Close()
	repository := TimesheetRepository{
		DatabaseConnection: databaseConnection,
	}

	actual, err := repository.GetIncomes(employeeID, year, month)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func Test_VerifyTransactionTimesheet_Input_Transaction_EmployeeID_001_Should_Be_Create_TransactionTimesheet_And_Update_TransactionTimesheet(t *testing.T) {
	transactionTimesheetList := []model.TransactionTimesheet{
		{
			EmployeeID:             "001",
			Month:                  12,
			Year:                   2019,
			CompanyID:              2,
			EmployeeNameTH:         "ประธาน ด่านสกุลเจริญกิจ",
			Coaching:               20000.00,
			Training:               0.00,
			Other:                  6500.00,
			TotalIncomes:           6500.00,
			Salary:                 25000.00,
			IncomeTax1:             0.00,
			SocialSecurity:         750.00,
			NetSalary:              24250.00,
			Wage:                   6500.00,
			IncomeTax53Percentage:  5,
			IncomeTax53:            325.00,
			NetWage:                6175.00,
			NetTransfer:            30425.00,
			StatusCheckingTransfer: "รอการตรวจสอบ",
		},
		{
			EmployeeID:             "001",
			Month:                  12,
			Year:                   2019,
			CompanyID:              1,
			EmployeeNameTH:         "ประธาน ด่านสกุลเจริญกิจ",
			Coaching:               30000.00,
			Training:               0.00,
			Other:                  6500.00,
			TotalIncomes:           6500.00,
			Salary:                 25000.00,
			IncomeTax1:             0.00,
			SocialSecurity:         750.00,
			NetSalary:              24250.00,
			Wage:                   6500.00,
			IncomeTax53Percentage:  5,
			IncomeTax53:            325.00,
			NetWage:                6175.00,
			NetTransfer:            30425.00,
			StatusCheckingTransfer: "รอการตรวจสอบ",
		},
	}
	databaseConnection, _ := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/timesheet")
	defer databaseConnection.Close()
	repository := TimesheetRepository{
		DatabaseConnection: databaseConnection,
	}

	err := repository.VerifyTransactionTimesheet(transactionTimesheetList)

	assert.Equal(t, nil, err)
}

func Test_CreateTransactionTimesheet_Input_TransactionID_00620191202_TransactionTimesheet_EmployeeID_006_Should_Be_No_Error(t *testing.T) {
	transactionID := "00620191202"
	transactionTimesheet := model.TransactionTimesheet{
		EmployeeID:             "006",
		EmployeeNameTH:         "ภาณุมาศ แสนโท",
		Month:                  12,
		Year:                   2019,
		CompanyID:              2,
		Coaching:               0.00,
		Training:               0.00,
		Other:                  6500.00,
		TotalIncomes:           6500.00,
		Salary:                 25000.00,
		IncomeTax1:             0.00,
		SocialSecurity:         750.00,
		NetSalary:              24250.00,
		Wage:                   6500.00,
		IncomeTax53Percentage:  5,
		IncomeTax53:            325.00,
		NetWage:                6175.00,
		NetTransfer:            30425.00,
		StatusCheckingTransfer: "รอการตรวจสอบ",
	}
	databaseConnection, _ := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/timesheet")
	defer databaseConnection.Close()
	repository := TimesheetRepository{
		DatabaseConnection: databaseConnection,
	}

	err := repository.CreateTransactionTimesheet(transactionTimesheet, transactionID)

	assert.Equal(t, nil, err)
}

func Test_UpdateTransactionTimesheet_Input_TransactionID_00120191102_TransactionTimesheet_EmployeeID_001_Should_Be_No_Error(t *testing.T) {
	transactionID := "00120191102"
	transactionTimesheet := model.TransactionTimesheet{
		EmployeeID:            "001",
		Month:                 11,
		Year:                  2019,
		CompanyID:             2,
		Coaching:              10000.00,
		Training:              10000.00,
		Other:                 6500.00,
		TotalIncomes:          6500.00,
		Salary:                25000.00,
		IncomeTax1:            1000.00,
		SocialSecurity:        750.00,
		NetSalary:             24250.00,
		Wage:                  6500.00,
		IncomeTax53Percentage: 5,
		IncomeTax53:           325.00,
		NetWage:               6175.00,
		NetTransfer:           30425.00,
	}
	databaseConnection, _ := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/timesheet")
	defer databaseConnection.Close()
	repository := TimesheetRepository{
		DatabaseConnection: databaseConnection,
	}

	err := repository.UpdateTransactionTimesheet(transactionTimesheet, transactionID)

	assert.Equal(t, nil, err)
}

func Test_CreateTimesheet_Input_EmployeeID_006_Month_12_Year_2019_Should_Be_No_Error(t *testing.T) {
	employeeID := "006"
	month := 12
	year := 2019
	databaseConnection, _ := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/timesheet")
	defer databaseConnection.Close()
	repository := TimesheetRepository{
		DatabaseConnection: databaseConnection,
	}

	err := repository.CreateTimesheet(employeeID, year, month)

	assert.Equal(t, nil, err)
}

func Test_UpdateTimesheet_Input_Timesheet_EmployeeID_007_Year_2019_Month_12_Should_Be_No_Error(t *testing.T) {
	timesheet := model.Timesheet{
		TotalHours:                    "120:30:30",
		TotalCoachingCustomerCharging: 90000.00,
		TotalCoachingPaymentRate:      10000.00,
		TotalTrainigWage:              20000.00,
		TotalOtherWage:                30000.00,
		PaymentWage:                   60000.00,
	}
	employeeID := "007"
	month := 12
	year := 2019
	databaseConnection, _ := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/timesheet")
	defer databaseConnection.Close()
	repository := TimesheetRepository{
		DatabaseConnection: databaseConnection,
	}

	err := repository.UpdateTimesheet(timesheet, employeeID, year, month)

	assert.Equal(t, nil, err)
}

func Test_GetTimesheet_Input_EmployeeID_003_Month_12_Year_2017_Should_Be_Timesheet(t *testing.T) {
	expected := model.Timesheet{
		ID:                            "003201712",
		EmployeeID:                    "003",
		Month:                         12,
		Year:                          2017,
		TotalHours:                    "88:00:00",
		TotalCoachingCustomerCharging: 0.00,
		TotalCoachingPaymentRate:      0.00,
		TotalTrainigWage:              120000.00,
		TotalOtherWage:                0.00,
		PaymentWage:                   120000.00,
		RatePerDay:                    15000.00,
		RatePerHour:                   1875.00,
	}
	employeeID := "003"
	month := 12
	year := 2017
	databaseConnection, _ := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/timesheet")
	defer databaseConnection.Close()
	repository := TimesheetRepository{
		DatabaseConnection: databaseConnection,
	}

	actual, err := repository.GetTimesheet(employeeID, year, month)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func Test_UpdateStatusTransfer_Input_TransactionID_0042019121_Status_TransferSuccess_Date_30_12_2019_Comment_FlightTicket_Should_Be_No_Error(t *testing.T) {
	transactionID := "0042019121"
	status := "โอนเงินเรียบร้อยแล้ว"
	date := "30/12/2019"
	comment := "หักค่าตั๋วเครื่องบิน"
	databaseConnection, _ := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/timesheet")
	defer databaseConnection.Close()
	repository := TimesheetRepository{
		DatabaseConnection: databaseConnection,
	}

	err := repository.UpdateStatusTransfer(transactionID, status, date, comment)

	assert.Equal(t, nil, err)
}

func Test_DeleteIncome_Input_IncomeID_47_Should_Be_No_Error(t *testing.T) {
	incomeID := 47
	databaseConnection, _ := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/timesheet")
	defer databaseConnection.Close()
	repository := TimesheetRepository{
		DatabaseConnection: databaseConnection,
	}

	err := repository.DeleteIncome(incomeID)

	assert.Equal(t, nil, err)
}

func Test_UpdateEmployeeDetails_Input_Employee_Should_Be_No_Error(t *testing.T) {
	employeeDetails := model.Employee{
		ID:                    10,
		EmployeeNameTH:        "ภาณุมาศ แสนโท",
		EmployeeNameENG:       "Panumars Seanto",
		Email:                 "panumars@scrum123.com",
		RatePerDay:            15000.00,
		RatePerHour:           1875.00,
		Salary:                25000.00,
		IncomeTax1:            0.00,
		SocialSecurity:        750.00,
		IncomeTax53Percentage: 5,
		Status:                "wage",
		TravelExpense:         0.00,
	}
	databaseConnection, _ := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/timesheet")
	defer databaseConnection.Close()
	repository := TimesheetRepository{
		DatabaseConnection: databaseConnection,
	}

	err := repository.UpdateEmployeeDetails(employeeDetails)

	assert.Equal(t, nil, err)
}

func Test_GetEmployeeIDByEmail_Input_Email_prathan_scrum123_com_Should_Be_001(t *testing.T) {
	expected := "001"
	email := "prathan@scrum123.com"
	databaseConnection, _ := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/timesheet")
	defer databaseConnection.Close()
	repository := TimesheetRepository{
		DatabaseConnection: databaseConnection,
	}

	actual, err := repository.GetEmployeeIDByEmail(email)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func Test_GetEmployeeIDByEmail_Input_Email_somkiat_scrum123_com_Should_Be_003(t *testing.T) {
	expected := "003"
	email := "somkiat@scrum123.com"
	databaseConnection, _ := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/timesheet")
	defer databaseConnection.Close()
	repository := TimesheetRepository{
		DatabaseConnection: databaseConnection,
	}

	actual, err := repository.GetEmployeeIDByEmail(email)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func Test_GetEmployeeIDByEmail_Input_Email_nareenart_scrum123_com_Should_Be_002(t *testing.T) {
	expected := "002"
	email := "nareenart@scrum123.com"
	databaseConnection, _ := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/timesheet")
	defer databaseConnection.Close()
	repository := TimesheetRepository{
		DatabaseConnection: databaseConnection,
	}

	actual, err := repository.GetEmployeeIDByEmail(email)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func Test_UpdatePictureToemployees_Input_Email_prathan_scrum123_com_And_Picture_Should_Be_No_Error(t *testing.T) {
	email := "prathan@scrum123.com"
	picture := "https://lh4.googleusercontent.com/-nA86bkk5Icc/AAAAAAAAAAI/AAAAAAAAAAA/Wixwdu9UCfU/photo.jpg"
	databaseConnection, _ := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/timesheet")
	defer databaseConnection.Close()
	repository := TimesheetRepository{
		DatabaseConnection: databaseConnection,
	}

	err := repository.UpdatePictureToemployees(picture, email)

	assert.Equal(t, nil, err)
}

func Test_GetTransactionTimesheets_Input_EmployeeID_001_Year_2017_Should_Be_TransactionTimesheetList(t *testing.T) {
	expected := []model.TransactionTimesheet{
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
	}
	employeeID := "001"
	year := 2017
	databaseConnection, _ := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/timesheet?parseTime=true")
	defer databaseConnection.Close()
	repository := TimesheetRepository{
		DatabaseConnection: databaseConnection,
	}

	actual, err := repository.GetTransactionTimesheets(employeeID, year)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func Test_GetProfileByEmail_Input_Email_nareenart_scrum123_com_Should_Be_EmployeeID_002_And_Picture(t *testing.T) {
	expected := model.Profile{
		EmployeeID: "002",
		Email:      "nareenart@scrum123.com",
		Picture:    "https://lh4.googleusercontent.com/-nA86bkk5Icc/AAAAAAAAAAI/AAAAAAAAAAA/Wixwdu9UCfU/photo.jpg",
	}
	email := "nareenart@scrum123.com"
	databaseConnection, _ := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/timesheet")
	defer databaseConnection.Close()
	repository := TimesheetRepository{
		DatabaseConnection: databaseConnection,
	}

	actual, err := repository.GetProfileByEmail(email)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}
