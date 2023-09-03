package models

import "time"

type User struct {
	ID        uint `gorm:"primaryKey"`
	Username  string`gorm:"not null" json:"username" form:"username" valid:"required~Your username is required"`
	Email     string`gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required,email~invalid email format"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
