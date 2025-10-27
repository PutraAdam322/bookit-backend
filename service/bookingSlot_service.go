package service

import "bookit.com/model"

type BookingSlotRepository interface {
	Create(slot *model.BookingSlot) (*model.BookingSlot, error)
	Update(slot *model.BookingSlot) (*model.BookingSlot, error)
	UpdateByBooking(slot *model.BookingSlot) (*model.BookingSlot, error)
	GetAll() ([]model.BookingSlot, error)
	GetByID(id uint) (*model.BookingSlot, error)
	Delete(id uint) error
	FindAvailableByFacility(facilityID uint) ([]model.BookingSlot, error)
}

type BookingSlotService struct {
	repo BookingSlotRepository
}

func NewBookingSlotService(repo BookingSlotRepository) *BookingSlotService {
	return &BookingSlotService{repo: repo}
}

func (s *BookingSlotService) GetAll() ([]model.BookingSlot, error) {
	return s.repo.GetAll()
}

func (s *BookingSlotService) GetByID(id uint) (*model.BookingSlot, error) {
	return s.repo.GetByID(id)
}

func (s *BookingSlotService) Create(slot *model.BookingSlot) (*model.BookingSlot, error) {
	return s.repo.Create(slot)
}

func (s *BookingSlotService) Update(slot *model.BookingSlot) (*model.BookingSlot, error) {
	_, err := s.repo.GetByID(slot.ID)
	if err != nil {
		return nil, err
	}
	return s.repo.Update(slot)
}

func (s *BookingSlotService) UpdateByBooking(slot *model.BookingSlot) (*model.BookingSlot, error) {
	_, err := s.repo.GetByID(slot.ID)
	if err != nil {
		return nil, err
	}
	return s.repo.UpdateByBooking(slot)
}

func (s *BookingSlotService) Delete(id uint) error {
	_, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *BookingSlotService) FindAvailableByFacility(facilityID uint) ([]model.BookingSlot, error) {
	return s.repo.FindAvailableByFacility(facilityID)
}
