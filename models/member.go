package models

import "time"

type Member struct {
	ID           int
	FirstName    string
	LastName     string
	Email        string
	PasswordHash string
	Phone        string
	Address      string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
