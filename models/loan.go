package models

import "time"

type Loan struct {
	ID           int
	MemberID     int
	BookID       int
	LoanDate     time.Time
	DueDate      time.Time
	ReturnedDate time.Time
}
