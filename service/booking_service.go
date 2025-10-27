package service

import (
	"bookit.com/model"
)

type BookingRepository interface {
	Create(booking *model.Booking) (*model.Booking, error)
	Update(booking *model.Booking) (*model.Booking, error)
	GetAll() ([]model.Booking, error)
	GetByID(id uint) (*model.Booking, error)
	//Delete(id uint) error
}

type BookingService struct {
	bookingRepository BookingRepository
}

func NewBookingService(bookingRepository BookingRepository) *BookingService {
	return &BookingService{
		bookingRepository: bookingRepository,
	}
}

// GetAll retrieves all bookings
func (s *BookingService) GetAll() ([]model.Booking, error) {
	return s.bookingRepository.GetAll()
}

// CreateBooking creates a new booking
func (s *BookingService) CreateBooking(booking *model.Booking) (*model.Booking, error) {
	return s.bookingRepository.Create(booking)
}

// GetByID retrieves a booking by ID
func (s *BookingService) GetByID(id uint) (*model.Booking, error) {
	return s.bookingRepository.GetByID(id)
}

// Update updates a booking after checking existence
func (s *BookingService) Update(booking *model.Booking) (*model.Booking, error) {
	_, err := s.bookingRepository.GetByID(booking.ID)
	if err != nil {
		return nil, err
	}
	return s.bookingRepository.Update(booking)
}

// Delete removes a booking by ID
/*func (s *BookingService) Delete(id uint) error {
	_, err := s.bookingRepository.GetByID(id)
	if err != nil {
		return err
	}
	return s.bookingRepository.Delete(id)
}*/
