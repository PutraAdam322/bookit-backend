package model

import "time"

type Facility struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name" gorm:"not null"`
	Price    float32 `json:"price" gorm:"not null"`
	Capacity uint    `json:"capacity" gorm:"not null"`
	//Available bool      `json:"available" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
