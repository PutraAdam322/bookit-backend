package model

import (
	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	ID            uint        `json:"id" gorm:"primaryKey"`
	TotalPrice    float32     `json:"total_price" gorm:"not null"`
	Status        string      `json:"status"`
	UserID        uint        `json:"user_id,omitempty"`
	BookingSlotID uint        `json:"booking_slot_id,omitempty"`
	BookingSlot   BookingSlot `gorm:"foreignKey:BookingSlotID;references:ID;" json:"booking_slot"`
	User          User        `json:"-" gorm:"-"`
}
