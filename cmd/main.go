package main

import (
	"flag"
	"log"
	"net/http"
	"timesheet/cmd/handler"
	"timesheet/config"
	"timesheet/internal/repository"
	"timesheet/internal/timesheet"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	var databaseConfigPath string
	flag.StringVar(&databaseConfigPath, "config", "./config/database.yml", "Database config path")
	flag.Parse()

	databaseConfig, err := config.SetupDatabaseConfig(databaseConfigPath)
	if err != nil {
		log.Fatal("Cannot read config", err.Error())
	}

	databaseConnection, err := sqlx.Connect("mysql", databaseConfig.GetURI())
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
		TimesheetRepository: timesheetRepository,
	}

	router := gin.Default()

	router.POST("/showSummaryTimesheet", api.GetSummaryHandler)
	router.POST("/addIncomeItem", api.CreateIncomeHandler)
	router.POST("/calculatePayment", api.CalculatePaymentHandler)
	router.POST("/showTimesheetByID", api.GetSummaryByIDHandler)
	router.POST("/updateStatusCheckingTransfer", api.UpdateStatusCheckingTransferHandler)
	router.POST("/deleteIncomeItem", api.DeleteIncomeHandler)
	router.StaticFS("/", http.Dir("ui"))
	log.Fatal(router.Run())
}
