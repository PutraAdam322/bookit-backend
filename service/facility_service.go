package service

import (
	"bookit.com/model"
)

type FacilityRepository interface {
	Create(facility *model.Facility) (*model.Facility, error)
	Update(facility *model.Facility) (*model.Facility, error)
	GetAll() ([]model.Facility, error)
	GetByID(id uint) (*model.Facility, error)
}

type FacilityService struct {
	facilityRepository FacilityRepository
}

func NewFacilityService(facilityRepository FacilityRepository) *FacilityService {
	return &FacilityService{
		facilityRepository: facilityRepository,
	}
}

func (f *FacilityService) GetAll() ([]model.Facility, error) {
	return f.facilityRepository.GetAll()
}

func (f *FacilityService) Create(facility *model.Facility) (*model.Facility, error) {
	return f.facilityRepository.Create(facility)
}

func (f *FacilityService) GetByID(id uint) (*model.Facility, error) {
	return f.facilityRepository.GetByID(id)
}

func (f *FacilityService) Update(facility *model.Facility) (*model.Facility, error) {
	_, err := f.facilityRepository.GetByID(facility.ID)
	if err != nil {
		return nil, err
	}

	return f.facilityRepository.Update(facility)
}
