package handler

import (
	"net/http"
	"timesheet/internal/model"
	"timesheet/internal/repository"
	"timesheet/internal/timesheet"

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
	TransactionID string `json:"transaction_id"`
	Status        string `json:"status"`
	Date          string `json:"date"`
	Comment       string `json:"comment"`
}

type DeleteIncomeRequest struct {
	IncomeID int `json:"id"`
}

type MemberRequest struct {
	MemberID string `json:"member_id"`
}

type TimesheetAPI struct {
	Timesheet           timesheet.TimesheetGateways
	TimesheetRepository repository.TimesheetRepositoryGateways
}

func (api TimesheetAPI) GetSummaryByIDHandler(context *gin.Context) {
	var request TimesheetRequest
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	summaryTimesheet, err := api.Timesheet.GetSummaryByID(request.MemberID, request.Year, request.Month)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	context.JSON(http.StatusOK, summaryTimesheet)
}

func (api TimesheetAPI) GetSummaryHandler(context *gin.Context) {
	accessToken := getAccessToken()
	if accessToken != "" {
		bearer := "Bearer " + accessToken
		context.Writer.Header().Set("Authorization", bearer)
	}
	var request Date
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	transactionTimesheetList, err := api.TimesheetRepository.GetSummary(request.Year, request.Month)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	context.JSON(http.StatusOK, transactionTimesheetList)
}

func (api TimesheetAPI) CreateIncomeHandler(context *gin.Context) {
	var request IncomeRequest
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	err = api.TimesheetRepository.CreateIncome(request.Year, request.Month, request.MemberID, request.Incomes)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	context.Status(http.StatusCreated)
}

func (api TimesheetAPI) CalculatePaymentHandler(context *gin.Context) {
	var request CalculatePaymentRequest
	err := context.BindJSON(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	incomeList, err := api.TimesheetRepository.GetIncomes(request.MemberID, request.Year, request.Month)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	timesheet := api.Timesheet.CalculatePayment(incomeList)
	err = api.TimesheetRepository.UpdateTimesheet(timesheet, request.MemberID, request.Year, request.Month)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	memberList, err := api.TimesheetRepository.GetMemberByID(request.MemberID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	transactionTimesheetList := api.Timesheet.CalculatePaymentSummary(memberList, incomeList, request.Year, request.Month)
	err = api.TimesheetRepository.VerifyTransactionTimsheet(transactionTimesheetList)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	context.Status(http.StatusOK)
}

func (api TimesheetAPI) UpdateStatusCheckingTransferHandler(context *gin.Context) {
	// reqToken := context.GetHeader("Authorization")
	// splitToken := strings.Split(reqToken, "Bearer")
	// reqToken = splitToken[1]
	// log.Printf("%+v", reqToken)

	var request UpdateStatusRequest
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	err = api.TimesheetRepository.UpdateStatusTransfer(request.TransactionID, request.Status, request.Date, request.Comment)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	context.Status(http.StatusOK)
}

func (api TimesheetAPI) DeleteIncomeHandler(context *gin.Context) {
	var request DeleteIncomeRequest
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	err = api.TimesheetRepository.DeleteIncome(request.IncomeID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	context.Status(http.StatusOK)
}

func (api TimesheetAPI) ShowMemberDetailsByIDHandler(context *gin.Context) {
	var request MemberRequest
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	memberList, err := api.TimesheetRepository.GetMemberByID(request.MemberID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	context.JSON(http.StatusOK, memberList)
}

func (api TimesheetAPI) UpdateMemberDetailsHandler(context *gin.Context) {
	var request model.Member
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	err = api.TimesheetRepository.UpdateMemberDetails(request)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	context.Status(http.StatusOK)
}
