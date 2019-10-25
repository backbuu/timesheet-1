package mockinternal

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

func (mock *MockRepository) CreateIncome(year, month int, memberID string, income model.Incomes) error {
	argument := mock.Called(year, month, memberID, income)
	return argument.Error(0)
}

func (mock *MockRepository) GetIncomes(memberID string, year, month int) ([]model.Incomes, error) {
	argument := mock.Called(memberID, year, month)
	return argument.Get(0).([]model.Incomes), argument.Error(1)
}

func (mock *MockRepository) GetMemberByID(memberID string) ([]model.Member, error) {
	argument := mock.Called(memberID)
	return argument.Get(0).([]model.Member), argument.Error(1)
}

func (mock *MockRepository) VerifyTransactionTimsheet(transactionTimesheet []model.TransactionTimesheet) error {
	argument := mock.Called(transactionTimesheet)
	return argument.Error(0)
}

func (mock *MockRepository) UpdateTimesheet(timesheet model.Timesheet, memberID string, year, month int) error {
	argument := mock.Called(timesheet, memberID, year, month)
	return argument.Error(0)
}

func (mock *MockRepository) CreateTimesheet(memberID string, year int, month int) error {
	argument := mock.Called(memberID, year, month)
	return argument.Error(0)
}
func (mock *MockRepository) GetTimesheet(memberID string, year, month int) (model.Timesheet, error) {
	argument := mock.Called(memberID, year, month)
	return argument.Get(0).(model.Timesheet), argument.Error(1)
}

func (mock *MockRepository) UpdateStatusTransfer(transactionID, status, date, comment string) error {
	argument := mock.Called(transactionID, status, date, comment)
	return argument.Error(0)
}

func (mock *MockRepository) DeleteIncome(incomeID int) error {
	argument := mock.Called(incomeID)
	return argument.Error(0)
}

func (mock *MockRepository) UpdateMemberDetails(memberDetails model.Member) error {
	argument := mock.Called(memberDetails)
	return argument.Error(0)
}