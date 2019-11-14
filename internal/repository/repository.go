package repository

import (
	"strconv"
	"time"
	"timesheet/internal/model"

	"github.com/jmoiron/sqlx"
)

const (
	initialStatusCheckingTransfer = "รอการตรวจสอบ"
	initialDateTransfer           = ""
	initialComment                = ""

	oneMinute = 60
	oneHour   = 60
)

type TimesheetRepositoryGateways interface {
	GetSummary(year, month int) ([]model.TransactionTimesheet, error)
	CreateIncome(year, month int, employeeID string, income model.Incomes) error
	VerifyTransactionTimesheet(transactionTimesheetList []model.TransactionTimesheet) error
	UpdateTimesheet(timesheet model.Timesheet, employeeID string, year, month int) error
	UpdateStatusTransfer(transactionID, status, date, comment string) error
	DeleteIncome(incomeID int) error
	UpdateEmployeeDetails(employeeDetails model.Employee) error
	UpdatePictureToemployees(picture, email string) error
	GetProfileByEmail(email string) (model.Profile, error)
	VerifyIncomeRequest(employeeID string, companyID int) bool
}

type TimesheetRepositoryGatewaysToTimesheet interface {
	GetEmployeeListByEmployeeID(employeeID string) ([]model.Employee, error)
	GetIncomes(employeeID string, year, month int) ([]model.Incomes, error)
	GetTimesheet(employeeID string, year, month int) (model.Timesheet, error)
	CreateTimesheet(employeeID string, year int, month int) error
	GetTransactionTimesheets(employeeID string, year int) ([]model.TransactionTimesheet, error)
	GetEmployeeIDByEmail(email string) (string, error)
}

type TimesheetRepository struct {
	DatabaseConnection *sqlx.DB
}

func (repository TimesheetRepository) GetSummary(year, month int) ([]model.TransactionTimesheet, error) {
	var transactionTimesheetList []model.TransactionTimesheet
	query := `SELECT * FROM transactions WHERE year = ? AND month = ? ORDER BY employee_id ASC, company_id ASC`
	err := repository.DatabaseConnection.Select(&transactionTimesheetList, query, year, month)
	if err != nil {
		return []model.TransactionTimesheet{}, err
	}
	return transactionTimesheetList, nil
}

func (repository TimesheetRepository) CreateIncome(year, month int, employeeID string, income model.Incomes) error {
	query := `INSERT INTO incomes (employee_id, month, year, day, start_time_am,
		end_time_am, start_time_pm, end_time_pm, total_hours,
		coaching_customer_charging, coaching_payment_rate,
		training_wage, other_wage, company_id, description)
		VALUES ( ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ?, ? , ? )`
	toltalHours := income.EndTimeAM.Sub(income.StartTimeAM)
	toltalHours += income.EndTimePM.Sub(income.StartTimePM)
	toltalHoursInDay := time.Date(year, time.Month(month), income.Day, int(toltalHours.Hours()), int(toltalHours.Minutes())%oneHour, int(toltalHours.Seconds())%oneMinute, 0, time.UTC)
	transaction := repository.DatabaseConnection.MustBegin()
	transaction.MustExec(query, employeeID, month, year, income.Day, income.StartTimeAM,
		income.EndTimeAM, income.StartTimePM, income.EndTimePM,
		toltalHoursInDay, income.CoachingCustomerCharging, income.CoachingPaymentRate,
		income.TrainingWage, income.OtherWage, income.CompanyID, income.Description)
	err := transaction.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (repository TimesheetRepository) GetIncomes(employeeID string, year, month int) ([]model.Incomes, error) {
	var incomeList []model.Incomes
	query := `SELECT * FROM incomes WHERE employee_id = ? AND year = ? AND month = ? ORDER BY day ASC`
	err := repository.DatabaseConnection.Select(&incomeList, query, employeeID, year, month)
	if err != nil {
		return []model.Incomes{}, err
	}
	return incomeList, nil
}

func (repository TimesheetRepository) GetEmployeeListByEmployeeID(employeeID string) ([]model.Employee, error) {
	var employeeList []model.Employee
	query := `SELECT * FROM employees WHERE employee_id = ?`
	err := repository.DatabaseConnection.Select(&employeeList, query, employeeID)
	if err != nil {
		return []model.Employee{}, err
	}
	return employeeList, nil
}

func (repository TimesheetRepository) VerifyTransactionTimesheet(transactionTimesheetList []model.TransactionTimesheet) error {
	for _, transactionTimesheet := range transactionTimesheetList {
		query := `SELECT COUNT(id) FROM transactions WHERE id LIKE ?`
		var count int
		transactionID := transactionTimesheet.EmployeeID + strconv.Itoa(transactionTimesheet.Year) + strconv.Itoa(transactionTimesheet.Month) + strconv.Itoa(transactionTimesheet.CompanyID)
		err := repository.DatabaseConnection.Get(&count, query, transactionID)
		if err != nil {
			return err
		}
		if count == 0 {
			err = repository.CreateTransactionTimesheet(transactionTimesheet, transactionID)
			if err != nil {
				return err
			}
		}
		err = repository.UpdateTransactionTimesheet(transactionTimesheet, transactionID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (repository TimesheetRepository) CreateTransactionTimesheet(transactionTimesheet model.TransactionTimesheet, transactionID string) error {
	query := `INSERT INTO transactions (id, employee_id, month, year, company_id, employee_name_th, employee_name_eng, coaching, 
		training, other, total_incomes, salary, income_tax_1, social_security, net_salary, wage, 
		income_tax_53_percentage, income_tax_53, net_wage, net_transfer, status_checking_transfer, date_transfer, comment) 
		VALUES ( ? , ? ,? , ? ,? , ? ,? , ? ,? , ? ,? , ? ,? , ? ,? , ? ,? , ?, ? , ? , ?, ? , ?)`
	transaction := repository.DatabaseConnection.MustBegin()
	transaction.MustExec(query, transactionID, transactionTimesheet.EmployeeID, transactionTimesheet.Month,
		transactionTimesheet.Year, transactionTimesheet.CompanyID, transactionTimesheet.EmployeeNameTH, transactionTimesheet.EmployeeNameENG,
		transactionTimesheet.Coaching, transactionTimesheet.Training, transactionTimesheet.Other,
		transactionTimesheet.TotalIncomes, transactionTimesheet.Salary, transactionTimesheet.IncomeTax1,
		transactionTimesheet.SocialSecurity, transactionTimesheet.NetSalary, transactionTimesheet.Wage,
		transactionTimesheet.IncomeTax53Percentage, transactionTimesheet.IncomeTax53, transactionTimesheet.NetWage,
		transactionTimesheet.NetTransfer, initialStatusCheckingTransfer, initialDateTransfer, initialComment)
	err := transaction.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (repository TimesheetRepository) UpdateTransactionTimesheet(transactionTimesheet model.TransactionTimesheet, transactionID string) error {
	query := `UPDATE transactions SET coaching = ?, training = ?, other = ?, total_incomes = ?, salary = ?, 
		income_tax_1 = ?, social_security = ?, net_salary = ?, wage = ?, income_tax_53_percentage = ?, 
		income_tax_53 = ?, net_wage = ?, net_transfer = ? WHERE id = ?`
	transaction := repository.DatabaseConnection.MustBegin()
	transaction.MustExec(query, transactionTimesheet.Coaching, transactionTimesheet.Training,
		transactionTimesheet.Other, transactionTimesheet.TotalIncomes, transactionTimesheet.Salary,
		transactionTimesheet.IncomeTax1, transactionTimesheet.SocialSecurity, transactionTimesheet.NetSalary,
		transactionTimesheet.Wage, transactionTimesheet.IncomeTax53Percentage, transactionTimesheet.IncomeTax53,
		transactionTimesheet.NetWage, transactionTimesheet.NetTransfer, transactionID)
	err := transaction.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (repository TimesheetRepository) CreateTimesheet(employeeID string, year int, month int) error {
	query := `INSERT INTO timesheets (id, employee_id, month, year, total_hours, total_coaching_customer_charging,
		total_coaching_payment_rate, total_training_wage, total_other_wage, payment_wage, rate_per_day,rate_per_hour) 
		VALUES ( ? , ? , ? ,? , ? ,? , ? ,? , ? ,? , ?, ? )`
	transaction := repository.DatabaseConnection.MustBegin()
	timesheetID := employeeID + strconv.Itoa(year) + strconv.Itoa(month)
	var timesheet model.Timesheet
	transaction.MustExec(query, timesheetID, employeeID, month, year, timesheet.TotalHours,
		timesheet.TotalCoachingCustomerCharging, timesheet.TotalCoachingPaymentRate,
		timesheet.TotalTrainigWage, timesheet.TotalOtherWage, timesheet.PaymentWage,
		timesheet.RatePerDay, timesheet.RatePerHour)
	err := transaction.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (repository TimesheetRepository) UpdateTimesheet(timesheet model.Timesheet, employeeID string, year, month int) error {
	query := `UPDATE timesheets SET total_hours = ?, total_coaching_customer_charging = ?, 
		total_coaching_payment_rate = ?, total_training_wage = ?, total_other_wage = ?, payment_wage = ? WHERE id = ?`
	transaction := repository.DatabaseConnection.MustBegin()
	timesheetID := employeeID + strconv.Itoa(year) + strconv.Itoa(month)
	transaction.MustExec(query, timesheet.TotalHours, timesheet.TotalCoachingCustomerCharging,
		timesheet.TotalCoachingPaymentRate, timesheet.TotalTrainigWage, timesheet.TotalOtherWage,
		timesheet.PaymentWage, timesheetID)
	err := transaction.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (repository TimesheetRepository) GetTimesheet(employeeID string, year, month int) (model.Timesheet, error) {
	var timesheet model.Timesheet
	query := `SELECT * FROM timesheets WHERE id = ?`
	timesheetID := employeeID + strconv.Itoa(year) + strconv.Itoa(month)
	err := repository.DatabaseConnection.Get(&timesheet, query, timesheetID)
	if err != nil {
		return model.Timesheet{}, err
	}
	return timesheet, nil
}

func (repository TimesheetRepository) UpdateStatusTransfer(transactionID, status, date, comment string) error {
	query := `UPDATE transactions SET status_checking_transfer = ?, date_transfer = ?, comment = ? WHERE id = ?`
	transaction := repository.DatabaseConnection.MustBegin()
	transaction.MustExec(query, status, date, comment, transactionID)
	err := transaction.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (repository TimesheetRepository) DeleteIncome(incomeID int) error {
	query := `DELETE FROM incomes WHERE id = ?`
	transaction := repository.DatabaseConnection.MustBegin()
	transaction.MustExec(query, incomeID)
	err := transaction.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (repository TimesheetRepository) UpdateEmployeeDetails(employeeDetails model.Employee) error {
	query := `UPDATE employees SET employee_name_th = ?, employee_name_eng = ?, email = ?, rate_per_day = ?, rate_per_hour = ?, salary = ?, income_tax_1 = ?, social_security = ?, income_tax_53_percentage = ?, travel_expense = ?, status = ? WHERE id = ?`
	transaction := repository.DatabaseConnection.MustBegin()
	transaction.MustExec(query, employeeDetails.EmployeeNameTH, employeeDetails.EmployeeNameENG, employeeDetails.Email,
		employeeDetails.RatePerDay, employeeDetails.RatePerHour, employeeDetails.Salary,
		employeeDetails.IncomeTax1, employeeDetails.SocialSecurity, employeeDetails.IncomeTax53Percentage,
		employeeDetails.TravelExpense, employeeDetails.Status, employeeDetails.ID)
	err := transaction.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (repository TimesheetRepository) GetEmployeeIDByEmail(email string) (string, error) {
	var employeeID string
	query := `SELECT employee_id FROM employees WHERE email LIKE ?`
	err := repository.DatabaseConnection.Get(&employeeID, query, email)
	if err != nil {
		return employeeID, err
	}
	return employeeID, nil
}

func (repository TimesheetRepository) UpdatePictureToemployees(picture, email string) error {
	query := `UPDATE employees SET picture = ? WHERE email = ?`
	transaction := repository.DatabaseConnection.MustBegin()
	transaction.MustExec(query, picture, email)
	err := transaction.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (repository TimesheetRepository) GetTransactionTimesheets(employeeID string, year int) ([]model.TransactionTimesheet, error) {
	var transactionTimesheetList []model.TransactionTimesheet
	query := `SELECT * FROM transactions WHERE employee_id = ? AND year = ?`
	err := repository.DatabaseConnection.Select(&transactionTimesheetList, query, employeeID, year)
	if err != nil {
		return []model.TransactionTimesheet{}, err
	}
	return transactionTimesheetList, nil
}

func (repository TimesheetRepository) GetProfileByEmail(email string) (model.Profile, error) {
	var profile model.Profile
	query := `SELECT employee_id, picture FROM employees WHERE email LIKE ?`
	err := repository.DatabaseConnection.Get(&profile, query, email)
	profile.Email = email
	if err != nil {
		return profile, err
	}
	return profile, nil
}

func (repository TimesheetRepository) VerifyIncomeRequest(employeeID string, companyID int) bool {
	var count int
	query := `SELECT COUNT(id) FROM employees WHERE employee_id LIKE ? AND company_id = ?`
	err := repository.DatabaseConnection.Get(&count, query, employeeID, companyID)
	if err != nil {
		return false
	}
	if count == 0 {
		return false
	}
	return true
}
