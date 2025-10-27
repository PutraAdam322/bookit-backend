package dto

type CreateBookingDTO struct {
	TotalPrice    float32 `json:"total_price,omitempty" binding:"required"`
	BookingSlotID uint    `json:"booking_slot_id,omitempty" binding:"required"`
}

type UpdateBookingDTO struct {
	ID            uint    `json:"id,omitempty" binding:"required"`
	TotalPrice    float32 `json:"total_price,omitempty" binding:"required"`
	BookingSlotID uint    `json:"booking_slot_id,omitempty" binding:"required"`
}
