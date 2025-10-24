package model

import "time"

type User struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name" gorm:"not null"`
	Email        string    `json:"email" gorm:"not null; unique"`
	Password     string    `json:"-" gorm:"-"`
	PasswordHash string    `json:"-"`
	IsAdmin      bool      `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
