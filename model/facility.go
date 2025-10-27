package model

import "gorm.io/gorm"

type Facility struct {
	gorm.Model
	ID           uint    `json:"id" gorm:"primaryKey"`
	Name         string  `json:"name" gorm:"not null"`
	Price        float32 `json:"price" gorm:"not null"`
	Capacity     uint    `json:"capacity" gorm:"not null"`
	Available    bool    `json:"available" gorm:"not null"`
	BookingSlots []BookingSlot
}
