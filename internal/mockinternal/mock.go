package mockinternal

import (
	"timesheet/internal/model"

	"github.com/stretchr/testify/mock"
)

type MockRepositoryToTimesheet struct {
	mock.Mock
}

func (mock *MockRepositoryToTimesheet) GetIncomes(memberID string, year, month int) ([]model.Incomes, error) {
	argument := mock.Called(memberID, year, month)
	return argument.Get(0).([]model.Incomes), argument.Error(1)
}

func (mock *MockRepositoryToTimesheet) GetMemberListByMemberID(memberID string) ([]model.Member, error) {
	argument := mock.Called(memberID)
	return argument.Get(0).([]model.Member), argument.Error(1)
}

func (mock *MockRepositoryToTimesheet) GetVerifyAuthenticationByAccessToken(accessToken string) (model.VerifyAuthentication, error) {
	argument := mock.Called(accessToken)
	return argument.Get(0).(model.VerifyAuthentication), argument.Error(1)
}

func (mock *MockRepositoryToTimesheet) CreateTimesheet(memberID string, year int, month int) error {
	argument := mock.Called(memberID, year, month)
	return argument.Error(0)
}

func (mock *MockRepositoryToTimesheet) GetTimesheet(memberID string, year, month int) (model.Timesheet, error) {
	argument := mock.Called(memberID, year, month)
	return argument.Get(0).(model.Timesheet), argument.Error(1)
}

func (mock *MockRepositoryToTimesheet) GetTransactionTimesheets(memberID string, year int) ([]model.TransactionTimesheet, error) {
	argument := mock.Called(memberID, year)
	return argument.Get(0).([]model.TransactionTimesheet), argument.Error(1)
}
