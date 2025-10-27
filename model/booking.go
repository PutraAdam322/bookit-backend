package model

import (
	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	ID            uint        `json:"id" gorm:"primaryKey"`
	TotalPrice    float32     `json:"total_price" gorm:"not null"`
	UserID        uint        `json:"user_id"`
	BookingSlotID uint        `json:"booking_slot_id"`
	BookingSlot   BookingSlot `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	User          User        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
