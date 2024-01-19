package model

import "time"

type Users struct {
	Id                 int       `gorm:"type:int;primary_key"`
	Username           string    `gorm:"type:varchar(255);not null"`
	Email              string    `gorm:"uniqueIndex;not null"`
	Password           string    `gorm:"not null"`
	PasswordResetToken int       `gorm:"unique;default:null"`
	PasswordResetAt    time.Time `gorm:"default:null"`
}
