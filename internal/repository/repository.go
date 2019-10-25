package repository

import (
	"strconv"
	"timesheet/internal/model"

	"github.com/jmoiron/sqlx"
)

const (
	InitialStatusCheckingTransfer = "รอการตรวจสอบ"
	InitialDateTransfer           = ""
	InitialComment                = ""
)

type TimesheetRepositoryGateways interface {
	GetSummary(year, month int) ([]model.TransactionTimesheet, error)
	GetMemberByID(memberID string) ([]model.Member, error)
	GetIncomes(memberID string, year, month int) ([]model.Incomes, error)
	CreateIncome(year, month int, memberID string, income model.Incomes) error
	VerifyTransactionTimsheet(transactionTimesheet []model.TransactionTimesheet) error
	UpdateTimesheet(timesheet model.Timesheet, memberID string, year, month int) error
	CreateTimesheet(memberID string, year int, month int) error
	GetTimesheet(memberID string, year, month int) (model.Timesheet, error)
	UpdateStatusTransfer(transactionID, status, date, comment string) error
	DeleteIncome(incomeID int) error
	UpdateMemberDetails(memberDetails model.Member) error
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
	transaction := repository.DatabaseConnection.MustBegin()
	transaction.MustExec(query, memberID, month, year, income.Day, income.StartTimeAM,
		income.EndTimeAM, income.StartTimePM, income.EndTimePM, income.Overtime,
		income.TotalHours, income.CoachingCustomerCharging, income.CoachingPaymentRate,
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

func (repository TimesheetRepository) GetMemberByID(memberID string) ([]model.Member, error) {
	var memberList []model.Member
	query := `SELECT * FROM members WHERE member_id = ?`
	err := repository.DatabaseConnection.Select(&memberList, query, memberID)
	if err != nil {
		return []model.Member{}, err
	}
	return memberList, nil
}

func (repository TimesheetRepository) VerifyTransactionTimsheet(transactionTimesheet []model.TransactionTimesheet) error {
	for _, transactionTimesheet := range transactionTimesheet {
		query := `SELECT COUNT(id) FROM transactions WHERE id LIKE ?`
		var count int
		transactionID := transactionTimesheet.MemberID + strconv.Itoa(transactionTimesheet.Year) + strconv.Itoa(transactionTimesheet.Month) + transactionTimesheet.Company
		err := repository.DatabaseConnection.Get(&count, query, transactionID)
		if err != nil {
			return err
		}
		if count == 0 {
			err = repository.CreateTransactionTimsheet(transactionTimesheet, transactionID)
			if err != nil {
				return err
			}
		}
		err = repository.UpdateTransactionTimsheet(transactionTimesheet, transactionID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (repository TimesheetRepository) CreateTransactionTimsheet(transactionTimesheet model.TransactionTimesheet, transactionID string) error {
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
		transactionTimesheet.NetTransfer, InitialStatusCheckingTransfer, InitialDateTransfer, InitialComment)
	err := transaction.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (repository TimesheetRepository) UpdateTransactionTimsheet(transactionTimesheet model.TransactionTimesheet, transactionID string) error {
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
	var payment model.Timesheet
	query := `SELECT * FROM timesheets WHERE id = ?`
	timesheetID := memberID + strconv.Itoa(year) + strconv.Itoa(month)
	err := repository.DatabaseConnection.Get(&payment, query, timesheetID)
	if err != nil {
		return model.Timesheet{}, err
	}
	return payment, nil
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
