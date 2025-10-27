package controller

import (
	"net/http"

	"bookit.com/dto"
	"bookit.com/model"

	apix "bookit.com/utils/api"
	validatorx "bookit.com/utils/validator"

	"github.com/gin-gonic/gin"
)

type BookingService interface {
	CreateBooking(booking *model.Booking) (*model.Booking, error)
	Update(booking *model.Booking) (*model.Booking, error)
	GetAll() ([]model.Booking, error)
	GetByID(id uint) (*model.Booking, error)
}

type BookingController struct {
	bookingService     BookingService
	bookingSlotService BookingSlotService
}

func NewBookingController(bookingService BookingService, bookingSlotService BookingSlotService) *BookingController {
	return &BookingController{
		bookingService:     bookingService,
		bookingSlotService: bookingSlotService,
	}
}

func (c *BookingController) GetAll(ctx *gin.Context) {
	bookings, err := c.bookingService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, apix.HTTPResponse{
			Message: "failed to get bookings",
			Data:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, apix.HTTPResponse{
		Message: "get all bookings successfully",
		Data:    bookings,
	})
}

func (c *BookingController) Create(ctx *gin.Context) {
	IntUserID := ctx.GetInt("user_id")
	UserID := uint(IntUserID)
	var input dto.CreateBookingDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ve, _ := validatorx.ParseValidatorErrors(err)
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "invalid input data",
			Data:    ve,
		})
		return
	}

	booking := model.Booking{
		TotalPrice:    input.TotalPrice,
		BookingSlotID: input.BookingSlotID,
		UserID:        UserID,
	}

	bookingSlot, errG := c.bookingSlotService.GetByID(input.BookingSlotID)
	_, errU := c.bookingSlotService.UpdateByBooking(bookingSlot)
	created, err := c.bookingService.CreateBooking(&booking)
	if errG != nil || errU != nil || err != nil {
		ctx.JSON(http.StatusInternalServerError, apix.HTTPResponse{
			Message: "failed to create booking",
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, created)
}

func (c *BookingController) Update(ctx *gin.Context) {
	var input dto.UpdateBookingDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ve, _ := validatorx.ParseValidatorErrors(err)
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "invalid input data",
			Data:    ve,
		})
		return
	}

	booking := model.Booking{
		ID:         input.ID,
		TotalPrice: input.TotalPrice,
	}

	updated, err := c.bookingService.Update(&booking)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, apix.HTTPResponse{
			Message: "failed to update booking",
			Data:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, apix.HTTPResponse{
		Message: "booking is updated",
		Data:    updated,
	})
}
