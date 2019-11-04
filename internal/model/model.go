package model

import (
	"time"
)

type TransactionTimesheet struct {
	ID                     string  `json:"id"`
	MemberID               string  `db:"member_id" json:"member_id"`
	MemberNameTH           string  `db:"member_name_th" json:"member_name_th"`
	Month                  int     `db:"month" json:"month"`
	Year                   int     `db:"year" json:"year"`
	Company                string  `db:"company" json:"company"`
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
	MemberNameENG                 string    `json:"member_name_eng"`
	Email                         string    `json:"email"`
	OvertimeRate                  float64   `json:"overtime_rate"`
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
	MemberID                 string    `db:"member_id" json:"member_id"`
	Month                    int       `db:"month" json:"month"`
	Year                     int       `db:"year" json:"year"`
	Day                      int       `db:"day" json:"day"`
	StartTimeAM              time.Time `db:"start_time_am" json:"start_time_am"`
	EndTimeAM                time.Time `db:"end_time_am" json:"end_time_am"`
	StartTimePM              time.Time `db:"start_time_pm" json:"start_time_pm"`
	EndTimePM                time.Time `db:"end_time_pm" json:"end_time_pm"`
	Overtime                 int       `db:"overtime" json:"overtime"`
	TotalHours               time.Time `db:"total_hours" json:"total_hours"`
	CoachingCustomerCharging float64   `db:"coaching_customer_charging" json:"coaching_customer_charging"`
	CoachingPaymentRate      float64   `db:"coaching_payment_rate" json:"coaching_payment_rate"`
	TrainingWage             float64   `db:"training_wage" json:"training_wage"`
	OtherWage                float64   `db:"other_wage" json:"other_wage"`
	Company                  string    `db:"company" json:"company"`
	Description              string    `db:"description" json:"description"`
}

type Timesheet struct {
	ID                            string  `json:"id"`
	MemberID                      string  `db:"member_id" json:"member_id"`
	Month                         int     `db:"month" json:"month"`
	Year                          int     `db:"year" json:"year"`
	TotalHours                    string  `db:"total_hours" json:"total_hours"`
	TotalCoachingCustomerCharging float64 `db:"total_coaching_customer_charging" json:"total_coaching_customer_charging"`
	TotalCoachingPaymentRate      float64 `db:"total_coaching_payment_rate" json:"total_coaching_payment_rate"`
	TotalTrainigWage              float64 `db:"total_training_wage" json:"total_training_wage"`
	TotalOtherWage                float64 `db:"total_other_wage" json:"total_other_wage"`
	PaymentWage                   float64 `db:"payment_wage" json:"payment_wage"`
}

type Member struct {
	ID                    int     `json:"id"`
	MemberID              string  `db:"member_id" json:"member_id"`
	Company               string  `db:"company" json:"company"`
	MemberNameTH          string  `db:"member_name_th" json:"member_name_th"`
	MemberNameENG         string  `db:"member_name_eng" json:"member_name_eng"`
	Email                 string  `db:"email" json:"email"`
	OvertimeRate          float64 `db:"overtime_rate" json:"overtime_rate"`
	RatePerDay            float64 `db:"rate_per_day" json:"rate_per_day"`
	RatePerHour           float64 `db:"rate_per_hour" json:"rate_per_hour"`
	Salary                float64 `db:"salary" json:"salary"`
	IncomeTax1            float64 `db:"income_tax_1" json:"income_tax_1"`
	SocialSecurity        float64 `db:"social_security" json:"social_security"`
	IncomeTax53Percentage int     `db:"income_tax_53_percentage" json:"income_tax_53_percentage"`
	Status                string  `db:"status" json:"status"`
	TravelExpense         float64 `db:"travel_expense" json:"travel_expense"`
}

type Authentication struct {
	MemberID     string    `db:"member_id" json:"member_id"`
	Email        string    `db:"email" json:"email"`
	Picture      string    `db:"picture" json:"picture"`
	AccessToken  string    `db:"access_token" json:"access_token"`
	TokenType    string    `db:"token_type" json:"token_type"`
	RefreshToken string    `db:"refresh_token" json:"refresh_token"`
	Expiry       time.Time `db:"expiry" json:"expiry"`
}
type Token struct {
	AccessToken  string    `json:"access_token"`
	TokenType    string    `json:"token_type"`
	RefreshToken string    `json:"refresh_token"`
	Expiry       time.Time `json:"expiry"`
}
type UserInfo struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture"`
	HD            string `json:"hd"`
}

type Holiday struct {
	ID    int    `json:"id"`
	Day   int    `db:"day" json:"day"`
	Month int    `db:"month" json:"month"`
	Name  string `db:"name" json:"name"`
}

type Profile struct {
	MemberID string `db:"member_id" json:"member_id"`
	Email    string `db:"email" json:"email"`
	Picture  string `db:"picture" json:"picture"`
}
