package model

import (
	"time"

	"gorm.io/gorm"
)

type BookingSlot struct {
	gorm.Model
	ID          uint      `json:"id" gorm:"primaryKey"`
	FacilityID  uint      `json:"facility_id" gorm:"index"`
	StartTime   time.Time `json:"start_time" gorm:"not null;index"`
	EndTime     time.Time `json:"end_time" gorm:"not null;index"`
	IsAvailable bool      `json:"is_available" gorm:"not null;default:true"`
}
