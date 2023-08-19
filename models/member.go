package models

import "time"

type Member struct {
	ID        int
	FirstName string `gorm:"type:varchar(100)" json:"firt_name"`
	LastName  string `gorm:"type:varchar(100)" json:"last_name"`
	Email     string `gorm:"type:varchar(100)" json:"email"`
	Phone     string `gorm:"type:varchar(100)" json:"phone"`
	Address   string `gorm:"type:varchar(250)" json:"address"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
