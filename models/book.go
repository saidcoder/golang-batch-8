package models

import "time"

type Book struct {
	ID              int
	Title           string    `gorm:"type:varchar(100)" json:"title"`
	Author          string    `gorm:"type:varchar(60)" json:"author"`
	PublicationYear int       `json:"publication_year"`
	ISBN            string    `gorm:"type:varchar(20)" json:"isbn"`
	AvailableCopies int       `json:"available_copies"`
	TotalCopies     int       `json:"total_copies"`
	CreatedAt       time.Time `json:"created_at" time_format:"2006-01-02 00:00:00" time_utf:"8"`
	UpdatedAt       time.Time `json:"updated_at" time_format:"2006-01-02 00:00:00" time_utf:"8"`
}
