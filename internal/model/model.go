package model

import (
	"time"
)

type TransactionTimesheet struct {
	ID                     string  `json:"id"`
	EmployeeID             string  `db:"employee_id" json:"employee_id"`
	EmployeeNameTH         string  `db:"employee_name_th" json:"employee_name_th"`
	EmployeeNameENG        string  `db:"employee_name_eng" json:"employee_name_eng"`
	Month                  int     `db:"month" json:"month"`
	Year                   int     `db:"year" json:"year"`
	CompanyID              int     `db:"company_id" json:"company_id"`
	Coaching               float64 `db:"coaching" json:"coaching"`
	Training               float64 `db:"training" json:"training"`
	Other                  float64 `db:"other" json:"other"`
	TotalIncomes           float64 `db:"total_incomes" json:"total_incomes"`
	Salary                 float64 `db:"salary" json:"salary"`
	IncomeTax1             float64 `db:"income_tax_1" json:"income_tax_1"`
	SocialSecurity         float64 `db:"social_security" json:"social_security"`
	NetSalary              float64 `db:"net_salary" json:"net_salary"`
	Wage                   float64 `db:"wage" json:"wage"`
	IncomeTax53Percentage  int     `db:"income_tax_53_percentage" json:"income_tax_53_percentage"`
	IncomeTax53            float64 `db:"income_tax_53" json:"income_tax_53"`
	NetWage                float64 `db:"net_wage" json:"net_wage"`
	NetTransfer            float64 `db:"net_transfer" json:"net_transfer"`
	StatusCheckingTransfer string  `db:"status_checking_transfer" json:"status_checking_transfer"`
	DateTransfer           string  `db:"date_transfer" json:"date_transfer"`
	Comment                string  `db:"comment" json:"comment"`
}

type SummaryTimesheet struct {
	EmployeeNameENG               string    `json:"employee_name_eng"`
	Email                         string    `json:"email"`
	RatePerDay                    float64   `json:"rate_per_day"`
	RatePerHour                   float64   `json:"rate_per_hour"`
	Year                          int       `json:"year"`
	Month                         int       `json:"month"`
	Incomes                       []Incomes `json:"incomes"`
	TimesheetID                   string    `json:"timesheet_id"`
	TotalHours                    string    `json:"total_hours"`
	TotalCoachingCustomerCharging float64   `json:"total_coaching_customer_charging"`
	TotalCoachingPaymentRate      float64   `json:"total_coaching_payment_rate"`
	TotalTrainigWage              float64   `json:"total_training_wage"`
	TotalOtherWage                float64   `json:"total_other_wage"`
	PaymentWage                   float64   `json:"payment_wage"`
}

type Incomes struct {
	ID                       int       `json:"id"`
	EmployeeID               string    `db:"employee_id" json:"employee_id"`
	Month                    int       `db:"month" json:"month"`
	Year                     int       `db:"year" json:"year"`
	Day                      int       `db:"day" json:"day"`
	StartTimeAM              time.Time `db:"start_time_am" json:"start_time_am"`
	EndTimeAM                time.Time `db:"end_time_am" json:"end_time_am"`
	StartTimePM              time.Time `db:"start_time_pm" json:"start_time_pm"`
	EndTimePM                time.Time `db:"end_time_pm" json:"end_time_pm"`
	TotalHours               time.Time `db:"total_hours" json:"total_hours"`
	CoachingCustomerCharging float64   `db:"coaching_customer_charging" json:"coaching_customer_charging"`
	CoachingPaymentRate      float64   `db:"coaching_payment_rate" json:"coaching_payment_rate"`
	TrainingWage             float64   `db:"training_wage" json:"training_wage"`
	OtherWage                float64   `db:"other_wage" json:"other_wage"`
	CompanyID                int       `db:"company_id" json:"company_id"`
	Description              string    `db:"description" json:"description"`
}

type Timesheet struct {
	ID                            string  `json:"id"`
	EmployeeID                    string  `db:"employee_id" json:"employee_id"`
	Month                         int     `db:"month" json:"month"`
	Year                          int     `db:"year" json:"year"`
	TotalHours                    string  `db:"total_hours" json:"total_hours"`
	TotalCoachingCustomerCharging float64 `db:"total_coaching_customer_charging" json:"total_coaching_customer_charging"`
	TotalCoachingPaymentRate      float64 `db:"total_coaching_payment_rate" json:"total_coaching_payment_rate"`
	TotalTrainigWage              float64 `db:"total_training_wage" json:"total_training_wage"`
	TotalOtherWage                float64 `db:"total_other_wage" json:"total_other_wage"`
	PaymentWage                   float64 `db:"payment_wage" json:"payment_wage"`
	RatePerDay                    float64 `db:"rate_per_day" json:"rate_per_day"`
	RatePerHour                   float64 `db:"rate_per_hour" json:"rate_per_hour"`
}

type Employee struct {
	ID                    int     `json:"id"`
	EmployeeID            string  `db:"employee_id" json:"employee_id"`
	CompanyID             int     `db:"company_id" json:"company_id"`
	EmployeeNameTH        string  `db:"employee_name_th" json:"employee_name_th"`
	EmployeeNameENG       string  `db:"employee_name_eng" json:"employee_name_eng"`
	Email                 string  `db:"email" json:"email"`
	RatePerDay            float64 `db:"rate_per_day" json:"rate_per_day"`
	RatePerHour           float64 `db:"rate_per_hour" json:"rate_per_hour"`
	Salary                float64 `db:"salary" json:"salary"`
	IncomeTax1            float64 `db:"income_tax_1" json:"income_tax_1"`
	SocialSecurity        float64 `db:"social_security" json:"social_security"`
	IncomeTax53Percentage int     `db:"income_tax_53_percentage" json:"income_tax_53_percentage"`
	Status                string  `db:"status" json:"status"`
	TravelExpense         float64 `db:"travel_expense" json:"travel_expense"`
	Picture               string  `db:"picture" json:"picture"`
}

type Token struct {
	AccessToken           string    `json:"access_token"`
	TokenType             string    `json:"token_type"`
	RefreshToken          string    `json:"refresh_token"`
	IDTokenExpirationTime time.Time `json:"id_token_expiration_time"`
}

type UserInfo struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture"`
	HD            string `json:"hd"`
}

type Profile struct {
	EmployeeID string `db:"employee_id" json:"employee_id"`
	Email      string `db:"email" json:"email"`
	Picture    string `db:"picture" json:"picture"`
}

type SummaryTransactionTimesheet struct {
	EmployeeID             string                 `json:"employee_id"`
	Year                   int                    `json:"year"`
	TransactionTimesheets  []TransactionTimesheet `json:"transaction_timesheets"`
	TotalCoachingInYear    float64                `json:"total_coaching_in_year"`
	TotalTrainingInYear    float64                `json:"total_training_in_year"`
	TotalOtherInYear       float64                `json:"total_other_in_year"`
	TotalIncomesInYear     float64                `json:"total_incomes_in_year"`
	TotalSalaryInYear      float64                `json:"total_salary_in_year"`
	TotalNetSalaryInYear   float64                `json:"total_net_salary_in_year"`
	TotalWageInYear        float64                `json:"total_wage_in_year"`
	TotalNetWageInYear     float64                `json:"total_net_wage_in_year"`
	TotalNetTransferInYear float64                `json:"total_net_transfer_in_year"`
}
