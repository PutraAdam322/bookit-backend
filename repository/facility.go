package repository

import (
	"fmt"

	"bookit.com/model"

	"gorm.io/gorm"
)

type FacilityRepository struct {
	db *gorm.DB
}

func NewFacilityRepository(db *gorm.DB) *FacilityRepository {
	return &FacilityRepository{
		db: db,
	}
}

func (f *FacilityRepository) GetAll() ([]model.Facility, error) {
	var facilities []model.Facility
	tx := f.db.Find(&facilities)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return facilities, nil
}

func (f *FacilityRepository) GetByID(id uint) (*model.Facility, error) {
	var facility model.Facility
	tx := f.db.Where("id = ?", id).First(&facility)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &facility, nil
}

func (f *FacilityRepository) Create(facility *model.Facility) (*model.Facility, error) {
	fmt.Println(&facility)
	tx := f.db.Create(&facility)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return facility, nil
}

func (f *FacilityRepository) Update(facility *model.Facility) (*model.Facility, error) {
	tx := f.db.Where("id = ?", facility.ID).Updates(facility)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return facility, nil
}
