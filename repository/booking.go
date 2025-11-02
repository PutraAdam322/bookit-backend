package repository

import (
	"bookit.com/model"

	"gorm.io/gorm"
)

type BookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) *BookingRepository {
	return &BookingRepository{
		db: db,
	}
}

func (r *BookingRepository) GetAll() ([]model.Booking, error) {
	var bookings []model.Booking
	tx := r.db.Find(&bookings)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return bookings, nil
}

func (r *BookingRepository) GetByID(id uint) (*model.Booking, error) {
	var booking model.Booking
	tx := r.db.Preload("BookingSlot.Facility").First(&booking, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &booking, nil
}

func (r *BookingRepository) GetByUserID(uid uint) ([]model.Booking, error) {
	var bookings []model.Booking
	tx := r.db.
		Preload("BookingSlot.Facility").
		Where("user_id = ?", uid).
		Find(&bookings)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return bookings, nil
}

func (r *BookingRepository) Create(booking *model.Booking) (*model.Booking, error) {
	tx := r.db.Create(booking)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return booking, nil
}

func (r *BookingRepository) Update(booking *model.Booking) (*model.Booking, error) {
	tx := r.db.Model(&model.Booking{}).Where("id = ?", booking.ID).Updates(booking)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return booking, nil
}

func (r *BookingRepository) Delete(id uint) error {
	tx := r.db.Delete(&model.Booking{}, id)
	return tx.Error
}
