package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"timesheet/cmd/handler"
	"timesheet/internal/repository"
	"timesheet/internal/timesheet"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DatabaseConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
}

var databaseConfig = &DatabaseConfig{
	Username: os.Getenv("DATABASE_USERNAME"),
	Password: os.Getenv("DATABASE_PASSWORD"),
	Host:     os.Getenv("DATABASE_HOST"),
	Port:     os.Getenv("DATABASE_PORT"),
	Database: os.Getenv("DATABASE"),
}

func (databaseConfig DatabaseConfig) getURI() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", databaseConfig.Username, databaseConfig.Password, databaseConfig.Host, databaseConfig.Port, databaseConfig.Database)
}

func main() {
	databaseConnection, err := sqlx.Connect("mysql", databaseConfig.getURI())
	if err != nil {
		log.Fatal("Cannot connect database", err.Error())
	}
	defer databaseConnection.Close()
	timesheetRepository := repository.TimesheetRepository{
		DatabaseConnection: databaseConnection,
	}
	api := handler.TimesheetAPI{
		Timesheet: timesheet.Timesheet{
			Repository: timesheetRepository,
		},
		Repository:            timesheetRepository,
		RepositoryToTimesheet: timesheetRepository,
	}

	router := gin.Default()
	router.GET("/login", handler.OauthGoogleLogin)
	router.GET("/callback", api.OauthGoogleCallback)
	router.POST("/logout", api.OauthGoogleLogout)
	router.GET("/deleteOauthState", handler.DeleteOauthStateCookie)
	router.GET("/showProfile", api.GetProfileHandler)
	router.POST("/showSummaryTimesheet", api.GetSummaryHandler)
	router.POST("/createIncome", api.CreateIncomeHandler)
	router.POST("/calculatePayment", api.CalculatePaymentHandler)
	router.POST("/showTimesheetByEmployeeID", api.GetSummaryByEmployeeIDHandler)
	router.POST("/updateStatusCheckingTransfer", api.UpdateStatusCheckingTransferHandler)
	router.POST("/deleteIncomeItem", api.DeleteIncomeHandler)
	router.POST("/showEmployeeDetailsByEmployeeID", api.ShowEmployeeDetailsByEmployeeIDHandler)
	router.POST("/updateEmployeeDetails", api.UpdateEmployeeDetailsHandler)
	router.POST("/showSummaryInYear", api.ShowSummaryInYearHandler)
	router.StaticFS("/home", http.Dir("ui"))
	log.Fatal(router.Run())
}
