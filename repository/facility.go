package repository

import (
	"fmt"
	"time"

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
	var fac model.Facility

	loc, err := time.LoadLocation("Asia/Kuala_Lumpur")
	if err != nil {
		loc = time.Local
	}
	now := time.Now().In(loc)
	todayStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
	weekday := now.Weekday()
	daysToSat := (int(time.Saturday) - int(weekday) + 7) % 7
	satStart := todayStart.AddDate(0, 0, daysToSat)
	sunEnd := satStart.AddDate(0, 0, 1).Add(time.Hour*23 + time.Minute*59 + time.Second*59).Add(time.Nanosecond * 999999999)

	fmt.Println(satStart, sunEnd)

	tx := f.db.
		Preload("BookingSlots", "start_time BETWEEN (?) AND (?)", satStart, sunEnd).
		First(&fac, id)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return &fac, nil
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
