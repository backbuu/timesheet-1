package model

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
	DateTransfer           *string `db:"date_transfer" json:"date_transfer"`
	Comment                *string `db:"comment" json:"comment"`
}

type Incomes struct {
	ID                       int    `json:"id"`
	MemberID                 string `json:"member_id"`
	Month                    int    `json:"month"`
	Year                     int    `json:"year"`
	Day                      int    `json:"day"`
	StartTimeAMHours         int    `json:"start_time_am_hours"`
	StartTimeAMMinutes       int    `json:"start_time_am_minutes"`
	StartTimeAMSeconds       int    `json:"start_time_am_seconds"`
	EndTimeAMHours           int    `json:"end_time_am_hours"`
	EndTimeAMMinutes         int    `json:"end_time_am_minutes"`
	EndTimeAMSeconds         int    `json:"end_time_am_seconds"`
	StartTimePMHours         int    `json:"start_time_pm_hours"`
	StartTimePMMinutes       int    `json:"start_time_pm_minutes"`
	StartTimePMSeconds       int    `json:"start_time_pm_seconds"`
	EndTimePMHours           int    `json:"end_time_pm_hours"`
	EndTimePMMinutes         int    `json:"end_time_pm_minutes"`
	EndTimePMSeconds         int    `json:"end_time_pm_seconds"`
	Overtime                 int    `json:"overtime"`
	TotalHoursHours          int    `json:"total_hours_hours"`
	TotalHoursMinutes        int    `json:"total_hours_minutes"`
	TotalHoursSeconds        int    `json:"total_hours_seconds"`
	CoachingCustomerCharging int    `json:"coaching_customer_charging"`
	CoachingPaymentRate      int    `json:"coaching_payment_rate"`
	TrainingWage             int    `json:"training_wage"`
	OtherWage                int    `json:"other_wage"`
	Company                  string `json:"company"`
	Description              string `json:"description"`
}

type Payment struct {
	MemberID                      string  `json:"member_id"`
	Month                         int     `json:"month"`
	Year                          int     `json:"year"`
	TotalHoursHours               int     `json:"total_hours_hours"`
	TotalHoursMinutes             int     `json:"total_hours_sinutes"`
	TotalHoursSeconds             int     `json:"total_hours_seconds"`
	TotalCoachingCustomerCharging float64 `json:"total_coaching_customer_charging"`
	TotalCoachingPaymentRate      float64 `json:"total_coaching_payment_rate"`
	TotalTrainigWage              float64 `json:"total_trainig_wage"`
	TotalOtherWage                float64 `json:"total_other_wage"`
	PaymentWage                   float64 `json:"payment_wage"`
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
	Status                *string `db:"status" json:"status"`
	TravelExpense         float64 `db:"travel_expense" json:"travel_expense"`
}

type Time struct {
	Hours   int
	Minutes int
	Seconds int
}
