package handler

import (
	"net/http"
	"timesheet/internal/model"
	"timesheet/internal/repository"
	"timesheet/internal/timesheet"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Date struct {
	Year  int `json:"year"`
	Month int `json:"month"`
}

type CalculatePaymentRequest struct {
	MemberID string `json:"member_id"`
	Year     int    `json:"year"`
	Month    int    `json:"month"`
}

type IncomeRequest struct {
	Year     int           `json:"year"`
	Month    int           `json:"month"`
	MemberID string        `json:"member_id"`
	Incomes  model.Incomes `json:"incomes"`
}

type TimesheetRequest struct {
	MemberID string `json:"member_id"`
	Year     int    `json:"year"`
	Month    int    `json:"month"`
}

type UpdateStatusRequest struct {
	MemberID      string `json:"member_id"`
	TransactionID string `json:"transaction_id"`
	Status        string `json:"status"`
	Date          string `json:"date"`
	Comment       string `json:"comment"`
}

type DeleteIncomeRequest struct {
	MemberID string `json:"member_id"`
	IncomeID int    `json:"id"`
}

type MemberRequest struct {
	MemberID string `json:"member_id"`
}

type HolidayRequest struct {
	Month int `json:"month"`
}

type SummaryInYearRequest struct {
	MemberID string `json:"member_id"`
	Year     int    `json:"year"`
}

type TimesheetAPI struct {
	Timesheet             timesheet.TimesheetGateways
	Repository            repository.TimesheetRepositoryGateways
	RepositoryToTimesheet repository.TimesheetRepositoryGatewaysToTimesheet
}

func (api TimesheetAPI) GetSummaryByIDHandler(context *gin.Context) {
	var request TimesheetRequest
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	summaryTimesheet, err := api.Timesheet.GetSummaryByID(request.MemberID, request.Year, request.Month)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, summaryTimesheet)
}

func (api TimesheetAPI) GetSummaryHandler(context *gin.Context) {
	var request Date
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	transactionTimesheetList, err := api.Repository.GetSummary(request.Year, request.Month)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, transactionTimesheetList)
}

func (api TimesheetAPI) CreateIncomeHandler(context *gin.Context) {
	var request IncomeRequest
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	requestHeader := context.GetHeader("Authorization")
	if requestHeader == "" {
		context.Status(http.StatusUnauthorized)
		return
	}
	token, _ := jwt.Parse(requestHeader, func(token *jwt.Token) (interface{}, error) {
		return []byte(""), nil
	})
	claims := token.Claims.(jwt.MapClaims)
	email := claims["email"].(string)
	expiry := claims["exp"].(float64)
	status := api.Timesheet.VerifyAuthentication(email, expiry, request.MemberID)
	if status != "Success" {
		context.Status(http.StatusUnauthorized)
		return
	}
	err = api.Repository.CreateIncome(request.Year, request.Month, request.MemberID, request.Incomes)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.Status(http.StatusCreated)
}

func (api TimesheetAPI) CalculatePaymentHandler(context *gin.Context) {
	var request CalculatePaymentRequest
	err := context.BindJSON(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	requestHeader := context.GetHeader("Authorization")
	if requestHeader == "" {
		context.Status(http.StatusUnauthorized)
		return
	}
	token, _ := jwt.Parse(requestHeader, func(token *jwt.Token) (interface{}, error) {
		return []byte(""), nil
	})
	claims := token.Claims.(jwt.MapClaims)
	email := claims["email"].(string)
	expiry := claims["exp"].(float64)
	status := api.Timesheet.VerifyAuthentication(email, expiry, request.MemberID)
	if status != "Success" {
		context.Status(http.StatusUnauthorized)
		return
	}
	incomeList, err := api.RepositoryToTimesheet.GetIncomes(request.MemberID, request.Year, request.Month)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	timesheet := api.Timesheet.CalculatePayment(incomeList)
	err = api.Repository.UpdateTimesheet(timesheet, request.MemberID, request.Year, request.Month)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	memberList, err := api.RepositoryToTimesheet.GetMemberListByMemberID(request.MemberID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	transactionTimesheetList := api.Timesheet.CalculatePaymentSummary(memberList, incomeList, request.Year, request.Month)
	err = api.Repository.VerifyTransactionTimesheet(transactionTimesheetList)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.Status(http.StatusOK)
}

func (api TimesheetAPI) UpdateStatusCheckingTransferHandler(context *gin.Context) {
	var request UpdateStatusRequest
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	requestHeader := context.GetHeader("Authorization")
	if requestHeader == "" {
		context.Status(http.StatusUnauthorized)
		return
	}
	token, _ := jwt.Parse(requestHeader, func(token *jwt.Token) (interface{}, error) {
		return []byte(""), nil
	})
	claims := token.Claims.(jwt.MapClaims)
	email := claims["email"].(string)
	expiry := claims["exp"].(float64)
	status := api.Timesheet.VerifyAuthentication(email, expiry, request.MemberID)
	if status != "Success" {
		context.Status(http.StatusUnauthorized)
		return
	}
	err = api.Repository.UpdateStatusTransfer(request.TransactionID, request.Status, request.Date, request.Comment)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.Status(http.StatusOK)
}

func (api TimesheetAPI) DeleteIncomeHandler(context *gin.Context) {
	var request DeleteIncomeRequest
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	requestHeader := context.GetHeader("Authorization")
	if requestHeader == "" {
		context.Status(http.StatusUnauthorized)
		return
	}
	token, _ := jwt.Parse(requestHeader, func(token *jwt.Token) (interface{}, error) {
		return []byte(""), nil
	})
	claims := token.Claims.(jwt.MapClaims)
	email := claims["email"].(string)
	expiry := claims["exp"].(float64)
	status := api.Timesheet.VerifyAuthentication(email, expiry, request.MemberID)
	if status != "Success" {
		context.Status(http.StatusUnauthorized)
		return
	}
	err = api.Repository.DeleteIncome(request.IncomeID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.Status(http.StatusOK)
}

func (api TimesheetAPI) ShowMemberDetailsByIDHandler(context *gin.Context) {
	var request MemberRequest
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	memberList, err := api.RepositoryToTimesheet.GetMemberListByMemberID(request.MemberID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, memberList)
}

func (api TimesheetAPI) UpdateMemberDetailsHandler(context *gin.Context) {
	var request model.Member
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	requestHeader := context.GetHeader("Authorization")
	if requestHeader == "" {
		context.Status(http.StatusUnauthorized)
		return
	}
	token, _ := jwt.Parse(requestHeader, func(token *jwt.Token) (interface{}, error) {
		return []byte(""), nil
	})
	claims := token.Claims.(jwt.MapClaims)
	email := claims["email"].(string)
	expiry := claims["exp"].(float64)
	status := api.Timesheet.VerifyAuthentication(email, expiry, request.MemberID)
	if status != "Success" {
		context.Status(http.StatusUnauthorized)
		return
	}
	err = api.Repository.UpdateMemberDetails(request)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.Status(http.StatusOK)
}

func (api TimesheetAPI) GetProfileHandler(context *gin.Context) {
	requestHeader := context.GetHeader("Authorization")
	if requestHeader == "" {
		context.Status(http.StatusUnauthorized)
		return
	}
	token, _ := jwt.Parse(requestHeader, func(token *jwt.Token) (interface{}, error) {
		return []byte(""), nil
	})
	claims := token.Claims.(jwt.MapClaims)
	email := claims["email"].(string)
	profile, err := api.Repository.GetProfileByEmail(email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, profile)
}

func (api TimesheetAPI) ShowSummaryInYearHandler(context *gin.Context) {
	var request SummaryInYearRequest
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	summaryTransactionTimesheet, err := api.Timesheet.GetSummaryInYearByID(request.MemberID, request.Year)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, summaryTransactionTimesheet)
}
