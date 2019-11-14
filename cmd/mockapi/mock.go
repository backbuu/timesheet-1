package mockapi

import (
	"timesheet/internal/model"

	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) GetSummary(year, month int) ([]model.TransactionTimesheet, error) {
	argument := mock.Called(year, month)
	return argument.Get(0).([]model.TransactionTimesheet), argument.Error(1)
}

func (mock *MockRepository) CreateIncome(year, month int, employeeID string, income model.Incomes) error {
	argument := mock.Called(year, month, employeeID, income)
	return argument.Error(0)
}

func (mock *MockRepository) VerifyTransactionTimesheet(transactionTimesheetList []model.TransactionTimesheet) error {
	argument := mock.Called(transactionTimesheetList)
	return argument.Error(0)
}

func (mock *MockRepository) UpdateTimesheet(timesheet model.Timesheet, employeeID string, year, month int) error {
	argument := mock.Called(timesheet, employeeID, year, month)
	return argument.Error(0)
}

func (mock *MockRepository) UpdateStatusTransfer(transactionID, status, date, comment string) error {
	argument := mock.Called(transactionID, status, date, comment)
	return argument.Error(0)
}

func (mock *MockRepository) DeleteIncome(incomeID int) error {
	argument := mock.Called(incomeID)
	return argument.Error(0)
}

func (mock *MockRepository) UpdateEmployeeDetails(employeeDetails model.Employee) error {
	argument := mock.Called(employeeDetails)
	return argument.Error(0)
}

func (mock *MockRepository) CreateAuthentication(userInfo model.UserInfo, token model.Token) error {
	argument := mock.Called(userInfo, token)
	return argument.Error(0)
}

func (mock *MockRepository) GetProfileByAccessToken(accessToken string) (model.Profile, error) {
	argument := mock.Called(accessToken)
	return argument.Get(0).(model.Profile), argument.Error(1)
}

func (mock *MockRepository) DeleteAuthentication(accessToken string) error {
	argument := mock.Called(accessToken)
	return argument.Error(0)
}

func (mock *MockRepository) UpdatePictureToemployees(picture, email string) error {
	argument := mock.Called(picture, email)
	return argument.Error(0)
}

func (mock *MockRepository) GetProfileByEmail(email string) (model.Profile, error) {
	argument := mock.Called(email)
	return argument.Get(0).(model.Profile), argument.Error(1)
}

func (mock *MockRepository) VerifyIncomeRequest(employeeID string, companyID int) bool {
	argument := mock.Called(employeeID, companyID)
	return argument.Bool(0)
}

type MockRepositoryToTimesheet struct {
	mock.Mock
}

func (mock *MockRepositoryToTimesheet) GetIncomes(employeeID string, year, month int) ([]model.Incomes, error) {
	argument := mock.Called(employeeID, year, month)
	return argument.Get(0).([]model.Incomes), argument.Error(1)
}

func (mock *MockRepositoryToTimesheet) GetEmployeeListByEmployeeID(employeeID string) ([]model.Employee, error) {
	argument := mock.Called(employeeID)
	return argument.Get(0).([]model.Employee), argument.Error(1)
}

func (mock *MockRepositoryToTimesheet) CreateTimesheet(employeeID string, year int, month int) error {
	argument := mock.Called(employeeID, year, month)
	return argument.Error(0)
}

func (mock *MockRepositoryToTimesheet) GetTimesheet(employeeID string, year, month int) (model.Timesheet, error) {
	argument := mock.Called(employeeID, year, month)
	return argument.Get(0).(model.Timesheet), argument.Error(1)
}

func (mock *MockRepositoryToTimesheet) GetTransactionTimesheets(employeeID string, year int) ([]model.TransactionTimesheet, error) {
	argument := mock.Called(employeeID, year)
	return argument.Get(0).([]model.TransactionTimesheet), argument.Error(1)
}

func (mock *MockRepositoryToTimesheet) GetEmployeeIDByEmail(email string) (string, error) {
	argument := mock.Called(email)
	return argument.String(0), argument.Error(1)
}

type MockTimesheet struct {
	mock.Mock
}

func (mock *MockTimesheet) CalculatePayment(incomes []model.Incomes) model.Timesheet {
	argument := mock.Called(incomes)
	return argument.Get(0).(model.Timesheet)
}

func (mock *MockTimesheet) CalculatePaymentSummary(employee []model.Employee, incomes []model.Incomes, year, month int) []model.TransactionTimesheet {
	argument := mock.Called(employee, incomes, year, month)
	return argument.Get(0).([]model.TransactionTimesheet)
}

func (mock *MockTimesheet) GetSummaryByID(employeeID string, year, month int) (model.SummaryTimesheet, error) {
	argument := mock.Called(employeeID, year, month)
	return argument.Get(0).(model.SummaryTimesheet), argument.Error(1)
}

func (mock *MockTimesheet) VerifyAuthentication(email string, idTokenExpirationTime float64) bool {
	argument := mock.Called(email, idTokenExpirationTime)
	return argument.Bool(0)
}

func (mock *MockTimesheet) GetSummaryInYearByEmployeeID(employeeID string, year int) (model.SummaryTransactionTimesheet, error) {
	argument := mock.Called(employeeID, year)
	return argument.Get(0).(model.SummaryTransactionTimesheet), argument.Error(1)
}
