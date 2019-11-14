package handler

import (
	"net/http"
	"timesheet/internal/model"
	"timesheet/internal/repository"
	"timesheet/internal/timesheet"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type DateRequest struct {
	Year  int `json:"year"`
	Month int `json:"month"`
}

type CalculatePaymentRequest struct {
	EmployeeID string `json:"employee_id"`
	Year       int    `json:"year"`
	Month      int    `json:"month"`
}

type CreateIncomeRequest struct {
	Year       int           `json:"year"`
	Month      int           `json:"month"`
	EmployeeID string        `json:"employee_id"`
	Incomes    model.Incomes `json:"incomes"`
}

type TimesheetRequest struct {
	EmployeeID string `json:"employee_id"`
	Year       int    `json:"year"`
	Month      int    `json:"month"`
}

type UpdateStatusRequest struct {
	EmployeeID    string `json:"employee_id"`
	TransactionID string `json:"transaction_id"`
	Status        string `json:"status"`
	Date          string `json:"date"`
	Comment       string `json:"comment"`
}

type DeleteIncomeRequest struct {
	EmployeeID string `json:"employee_id"`
	IncomeID   int    `json:"id"`
}

type EmployeeRequest struct {
	EmployeeID string `json:"employee_id"`
}

type SummaryInYearRequest struct {
	EmployeeID string `json:"employee_id"`
	Year       int    `json:"year"`
}

type TimesheetAPI struct {
	Timesheet             timesheet.TimesheetGateways
	Repository            repository.TimesheetRepositoryGateways
	RepositoryToTimesheet repository.TimesheetRepositoryGatewaysToTimesheet
}

func (api TimesheetAPI) GetSummaryByEmployeeIDHandler(context *gin.Context) {
	var request TimesheetRequest
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	summaryTimesheet, err := api.Timesheet.GetSummaryByID(request.EmployeeID, request.Year, request.Month)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, summaryTimesheet)
}

func (api TimesheetAPI) GetSummaryHandler(context *gin.Context) {
	var request DateRequest
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
	var request CreateIncomeRequest
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	email, idTokenExpirationTime := getEmailAndIDTokenExpirationTimeFromHeaderAuthorization(context)
	if !api.Timesheet.VerifyAuthentication(email, idTokenExpirationTime) {
		context.Status(http.StatusUnauthorized)
		return
	}
	if !api.Repository.VerifyIncomeRequest(request.EmployeeID,request.Incomes.CompanyID) {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = api.Repository.CreateIncome(request.Year, request.Month, request.EmployeeID, request.Incomes)
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
	email, idTokenExpirationTime := getEmailAndIDTokenExpirationTimeFromHeaderAuthorization(context)
	if !api.Timesheet.VerifyAuthentication(email, idTokenExpirationTime) {
		context.Status(http.StatusUnauthorized)
		return
	}
	incomeList, err := api.RepositoryToTimesheet.GetIncomes(request.EmployeeID, request.Year, request.Month)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	timesheet := api.Timesheet.CalculatePayment(incomeList)
	err = api.Repository.UpdateTimesheet(timesheet, request.EmployeeID, request.Year, request.Month)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	employList, err := api.RepositoryToTimesheet.GetEmployeeListByEmployeeID(request.EmployeeID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	transactionTimesheetList := api.Timesheet.CalculatePaymentSummary(employList, incomeList, request.Year, request.Month)
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
	email, idTokenExpirationTime := getEmailAndIDTokenExpirationTimeFromHeaderAuthorization(context)
	if !api.Timesheet.VerifyAuthentication(email, idTokenExpirationTime) {
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
	email, idTokenExpirationTime := getEmailAndIDTokenExpirationTimeFromHeaderAuthorization(context)
	if !api.Timesheet.VerifyAuthentication(email, idTokenExpirationTime) {
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

func (api TimesheetAPI) ShowEmployeeDetailsByEmployeeIDHandler(context *gin.Context) {
	var request EmployeeRequest
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	employList, err := api.RepositoryToTimesheet.GetEmployeeListByEmployeeID(request.EmployeeID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, employList)
}

func (api TimesheetAPI) UpdateEmployeeDetailsHandler(context *gin.Context) {
	var request model.Employee
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	email, idTokenExpirationTime := getEmailAndIDTokenExpirationTimeFromHeaderAuthorization(context)
	if !api.Timesheet.VerifyAuthentication(email, idTokenExpirationTime) {
		context.Status(http.StatusUnauthorized)
		return
	}
	err = api.Repository.UpdateEmployeeDetails(request)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.Status(http.StatusOK)
}

func (api TimesheetAPI) GetProfileHandler(context *gin.Context) {
	email, _ := getEmailAndIDTokenExpirationTimeFromHeaderAuthorization(context)
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
	summaryTransactionTimesheet, err := api.Timesheet.GetSummaryInYearByID(request.EmployeeID, request.Year)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, summaryTransactionTimesheet)
}

func getEmailAndIDTokenExpirationTimeFromHeaderAuthorization(context *gin.Context) (string, float64) {
	var email string
	var idTokenExpirationTime float64
	requestHeader := context.GetHeader("Authorization")
	if requestHeader == "" {
		context.Status(http.StatusUnauthorized)
		return email, idTokenExpirationTime
	}
	token, _ := jwt.Parse(requestHeader, func(token *jwt.Token) (interface{}, error) {
		return []byte(""), nil
	})
	claims := token.Claims.(jwt.MapClaims)
	email = claims["email"].(string)
	idTokenExpirationTime = claims["exp"].(float64)
	return email, idTokenExpirationTime
}
