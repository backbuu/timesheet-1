package mockinternal

import (
	"timesheet/internal/model"

	"github.com/stretchr/testify/mock"
)

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
