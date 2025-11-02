package repository

import (
	"errors"
	"fmt"

	"bookit.com/model"

	"gorm.io/gorm"
)

type BookingSlotRepository struct {
	db *gorm.DB
}

func NewBookingSlotRepository(db *gorm.DB) *BookingSlotRepository {
	return &BookingSlotRepository{db: db}
}

func (r *BookingSlotRepository) GetAll() ([]model.BookingSlot, error) {
	var slots []model.BookingSlot
	tx := r.db.Find(&slots)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return slots, nil
}

func (r *BookingSlotRepository) GetByID(id uint) (*model.BookingSlot, error) {
	var slot model.BookingSlot
	tx := r.db.Preload("Facility").First(&slot, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &slot, nil
}

func (r *BookingSlotRepository) Create(slot *model.BookingSlot) (*model.BookingSlot, error) {
	tx := r.db.Create(slot)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return slot, nil
}

func (r *BookingSlotRepository) Update(slot *model.BookingSlot) (*model.BookingSlot, error) {
	tx := r.db.Model(&model.BookingSlot{}).Where("id = ?", slot.ID).Updates(slot)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return slot, nil
}

func (r *BookingSlotRepository) UpdateByBooking(slot *model.BookingSlot) (*model.BookingSlot, error) {
	slot.IsAvailable = false
	fmt.Printf("id slot : %d %t \n", slot.ID, slot.IsAvailable)
	tx := r.db.Model(&model.BookingSlot{}).Where("id = ? AND updated_at = ?", slot.ID, slot.CreatedAt).Update("is_available", false)
	if tx.RowsAffected == 0 {
		return nil, errors.New("slot is booked already")
	}
	if tx.Error != nil {
		return nil, tx.Error
	}
	return slot, nil
}

func (r *BookingSlotRepository) UpdateByCancel(slot *model.BookingSlot) (*model.BookingSlot, error) {
	slot.IsAvailable = true
	tx := r.db.Model(&model.BookingSlot{}).Where("id = ?", slot.ID).Update("is_available", true)
	if tx.RowsAffected == 0 {
		return nil, errors.New("error")
	}
	if tx.Error != nil {
		return nil, tx.Error
	}
	return slot, nil
}

func (r *BookingSlotRepository) Delete(id uint) error {
	tx := r.db.Delete(&model.BookingSlot{}, id)
	return tx.Error
}

func (r *BookingSlotRepository) FindAvailableByFacility(facilityID uint) ([]model.BookingSlot, error) {
	var slots []model.BookingSlot
	tx := r.db.Where("facility_id = ? AND is_available = ?", facilityID, true).Find(&slots)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return slots, nil
}
