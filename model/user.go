package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           uint      `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name" gorm:"not null"`
	Email        string    `json:"email" gorm:"not null; unique"`
	Password     string    `json:"-" gorm:"-"`
	PasswordHash string    `json:"-"`
	IsAdmin      bool      `json:"-"`
	Bookings     []Booking `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
