package models

import "time"

type Loan struct {
	ID           int
	MemberID     int
	BookID       int
	LoanDate     time.Time `json:"loan_date" time_format:"2006-01-02 00:00:00" time_utf:"8"`
	DueDate      time.Time `json:"due_date" time_format:"2006-01-02 00:00:00" time_utf:"8"`
	ReturnedDate time.Time `json:"returned_date" time_format:"2006-01-02 00:00:00" time_utf:"8"`
}
