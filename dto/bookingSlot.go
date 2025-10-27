package dto

import "time"

type CreateBookingSlotDTO struct {
	FacilityID  uint      `json:"facility_id,omitempty" binding:"required"`
	StartTime   time.Time `json:"start_time,omitempty" binding:"required"`
	EndTime     time.Time `json:"end_time,omitempty" binding:"required"`
	IsAvailable *bool     `json:"is_available,omitempty"` // optional, pointer so omission keeps default
}

type UpdateBookingSlotDTO struct {
	ID          uint      `json:"id,omitempty" binding:"required"`
	FacilityID  uint      `json:"facility_id,omitempty" binding:"required"`
	StartTime   time.Time `json:"start_time,omitempty" binding:"required"`
	EndTime     time.Time `json:"end_time,omitempty" binding:"required"`
	IsAvailable *bool     `json:"is_available,omitempty"`
}
