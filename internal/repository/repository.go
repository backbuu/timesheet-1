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

	workingHours = 8
	oneMinute    = 60
	oneHour      = 60
)

type TimesheetRepositoryGateways interface {
	GetSummary(year, month int) ([]model.TransactionTimesheet, error)
	CreateIncome(year, month int, memberID string, income model.Incomes) error
	VerifyTransactionTimesheet(transactionTimesheetList []model.TransactionTimesheet) error
	UpdateTimesheet(timesheet model.Timesheet, memberID string, year, month int) error
	UpdateStatusTransfer(transactionID, status, date, comment string) error
	DeleteIncome(incomeID int) error
	UpdateMemberDetails(memberDetails model.Member) error
	UpdatePictureToMembers(picture, email string) error
	GetProfileByEmail(email string) (model.Profile, error)
}

type TimesheetRepositoryGatewaysToTimesheet interface {
	GetMemberListByMemberID(memberID string) ([]model.Member, error)
	GetIncomes(memberID string, year, month int) ([]model.Incomes, error)
	GetTimesheet(memberID string, year, month int) (model.Timesheet, error)
	CreateTimesheet(memberID string, year int, month int) error
	GetTransactionTimesheets(memberID string, year int) ([]model.TransactionTimesheet, error)
	GetMemberIDByEmail(email string) (string, error)
}

type TimesheetRepository struct {
	DatabaseConnection *sqlx.DB
}

func (repository TimesheetRepository) GetSummary(year, month int) ([]model.TransactionTimesheet, error) {
	var transactionTimesheetList []model.TransactionTimesheet
	query := `SELECT * FROM transactions WHERE year = ? AND month = ? ORDER BY member_id ASC, company DESC`
	err := repository.DatabaseConnection.Select(&transactionTimesheetList, query, year, month)
	if err != nil {
		return []model.TransactionTimesheet{}, err
	}
	return transactionTimesheetList, nil
}

func (repository TimesheetRepository) CreateIncome(year, month int, memberID string, income model.Incomes) error {
	query := `INSERT INTO incomes (member_id, month, year, day, start_time_am,
		end_time_am, start_time_pm, end_time_pm, overtime, total_hours,
		coaching_customer_charging, coaching_payment_rate,
		training_wage, other_wage, company, description)
		VALUES ( ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ?, ? , ? , ? )`
	toltalHours := income.EndTimeAM.Sub(income.StartTimeAM)
	toltalHours += income.EndTimePM.Sub(income.StartTimePM)
	var overtime float64
	if toltalHours > workingHours {
		overtime = toltalHours.Hours() - workingHours
	}
	toltalHoursInDay := time.Date(2006, 1, 2, int(toltalHours.Hours()), int(toltalHours.Minutes())%oneHour, int(toltalHours.Seconds())%oneMinute, 0, time.UTC)
	transaction := repository.DatabaseConnection.MustBegin()
	transaction.MustExec(query, memberID, month, year, income.Day, income.StartTimeAM,
		income.EndTimeAM, income.StartTimePM, income.EndTimePM, overtime,
		toltalHoursInDay, income.CoachingCustomerCharging, income.CoachingPaymentRate,
		income.TrainingWage, income.OtherWage, income.Company, income.Description)
	err := transaction.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (repository TimesheetRepository) GetIncomes(memberID string, year, month int) ([]model.Incomes, error) {
	var incomeList []model.Incomes
	query := `SELECT * FROM incomes WHERE member_id = ? AND year = ? AND month = ? ORDER BY day ASC`
	err := repository.DatabaseConnection.Select(&incomeList, query, memberID, year, month)
	if err != nil {
		return []model.Incomes{}, err
	}
	return incomeList, nil
}

func (repository TimesheetRepository) GetMemberListByMemberID(memberID string) ([]model.Member, error) {
	var memberList []model.Member
	query := `SELECT * FROM members WHERE member_id = ?`
	err := repository.DatabaseConnection.Select(&memberList, query, memberID)
	if err != nil {
		return []model.Member{}, err
	}
	return memberList, nil
}

func (repository TimesheetRepository) VerifyTransactionTimesheet(transactionTimesheetList []model.TransactionTimesheet) error {
	for _, transactionTimesheet := range transactionTimesheetList {
		query := `SELECT COUNT(id) FROM transactions WHERE id LIKE ?`
		var count int
		transactionID := transactionTimesheet.MemberID + strconv.Itoa(transactionTimesheet.Year) + strconv.Itoa(transactionTimesheet.Month) + transactionTimesheet.Company
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
	query := `INSERT INTO transactions (id, member_id, month, year, company, member_name_th, coaching, 
		training, other, total_incomes, salary, income_tax_1, social_security, net_salary, wage, 
		income_tax_53_percentage, income_tax_53, net_wage, net_transfer, status_checking_transfer, date_transfer, comment) 
		VALUES ( ? , ? ,? , ? ,? , ? ,? , ? ,? , ? ,? , ? ,? , ? ,? , ? ,? , ? , ? , ?, ? , ?)`
	transaction := repository.DatabaseConnection.MustBegin()
	transaction.MustExec(query, transactionID, transactionTimesheet.MemberID, transactionTimesheet.Month,
		transactionTimesheet.Year, transactionTimesheet.Company, transactionTimesheet.MemberNameTH,
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

func (repository TimesheetRepository) CreateTimesheet(memberID string, year int, month int) error {
	query := `INSERT INTO timesheets (id, member_id, month, year, total_hours, total_coaching_customer_charging,
		total_coaching_payment_rate, total_training_wage, total_other_wage, payment_wage) 
		VALUES ( ? , ? , ? ,? , ? ,? , ? ,? , ? ,? )`
	transaction := repository.DatabaseConnection.MustBegin()
	timesheetID := memberID + strconv.Itoa(year) + strconv.Itoa(month)
	var payment model.Timesheet
	transaction.MustExec(query, timesheetID, memberID, month, year, payment.TotalHours,
		payment.TotalCoachingCustomerCharging, payment.TotalCoachingPaymentRate,
		payment.TotalTrainigWage, payment.TotalOtherWage, payment.PaymentWage)
	err := transaction.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (repository TimesheetRepository) UpdateTimesheet(timesheet model.Timesheet, memberID string, year, month int) error {
	query := `UPDATE timesheets SET total_hours = ?, total_coaching_customer_charging = ?, 
		total_coaching_payment_rate = ?, total_training_wage = ?, total_other_wage = ?, payment_wage = ? WHERE id = ?`
	transaction := repository.DatabaseConnection.MustBegin()
	timesheetID := memberID + strconv.Itoa(year) + strconv.Itoa(month)
	transaction.MustExec(query, timesheet.TotalHours, timesheet.TotalCoachingCustomerCharging,
		timesheet.TotalCoachingPaymentRate, timesheet.TotalTrainigWage, timesheet.TotalOtherWage,
		timesheet.PaymentWage, timesheetID)
	err := transaction.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (repository TimesheetRepository) GetTimesheet(memberID string, year, month int) (model.Timesheet, error) {
	var timesheet model.Timesheet
	query := `SELECT * FROM timesheets WHERE id = ?`
	timesheetID := memberID + strconv.Itoa(year) + strconv.Itoa(month)
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

func (repository TimesheetRepository) UpdateMemberDetails(memberDetails model.Member) error {
	query := `UPDATE members SET member_name_th = ?, member_name_eng = ?, email = ?, overtime_rate = ?, rate_per_day = ?, rate_per_hour = ?, salary = ?, income_tax_1 = ?, social_security = ?, income_tax_53_percentage = ?, travel_expense = ?, status = ? WHERE id = ?`
	transaction := repository.DatabaseConnection.MustBegin()
	transaction.MustExec(query, memberDetails.MemberNameTH, memberDetails.MemberNameENG, memberDetails.Email,
		memberDetails.OvertimeRate, memberDetails.RatePerDay, memberDetails.RatePerHour, memberDetails.Salary,
		memberDetails.IncomeTax1, memberDetails.SocialSecurity, memberDetails.IncomeTax53Percentage,
		memberDetails.TravelExpense, memberDetails.Status, memberDetails.ID)
	err := transaction.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (repository TimesheetRepository) GetMemberIDByEmail(email string) (string, error) {
	var memberID string
	query := `SELECT member_id FROM members WHERE email LIKE ?`
	err := repository.DatabaseConnection.Get(&memberID, query, email)
	if err != nil {
		return memberID, err
	}
	return memberID, nil
}

func (repository TimesheetRepository) UpdatePictureToMembers(picture, email string) error {
	query := `UPDATE members SET picture = ? WHERE email = ?`
	transaction := repository.DatabaseConnection.MustBegin()
	transaction.MustExec(query, picture, email)
	err := transaction.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (repository TimesheetRepository) GetTransactionTimesheets(memberID string, year int) ([]model.TransactionTimesheet, error) {
	var transactionTimesheetList []model.TransactionTimesheet
	query := `SELECT * FROM transactions WHERE member_id = ? AND year = ?`
	err := repository.DatabaseConnection.Select(&transactionTimesheetList, query, memberID, year)
	if err != nil {
		return []model.TransactionTimesheet{}, err
	}
	return transactionTimesheetList, nil
}

func (repository TimesheetRepository) GetProfileByEmail(email string) (model.Profile, error) {
	var profile model.Profile
	query := `SELECT member_id, picture FROM members WHERE email LIKE ?`
	err := repository.DatabaseConnection.Get(&profile, query, email)
	profile.Email = email
	if err != nil {
		return profile, err
	}
	return profile, nil
}
