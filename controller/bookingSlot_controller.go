package controller

import (
	"net/http"
	"strconv"

	"bookit.com/dto"
	"bookit.com/model"

	apix "bookit.com/utils/api"
	validatorx "bookit.com/utils/validator"

	"github.com/gin-gonic/gin"
)

type BookingSlotService interface {
	Create(slot *model.BookingSlot) (*model.BookingSlot, error)
	Update(slot *model.BookingSlot) (*model.BookingSlot, error)
	UpdateByBooking(slot *model.BookingSlot) (*model.BookingSlot, error)
	UpdateByCancel(slot *model.BookingSlot) (*model.BookingSlot, error)
	GetAll() ([]model.BookingSlot, error)
	GetByID(id uint) (*model.BookingSlot, error)
	Delete(id uint) error
	FindAvailableByFacility(facilityID uint) ([]model.BookingSlot, error)
}

type BookingSlotController struct {
	service BookingSlotService
}

func NewBookingSlotController(s BookingSlotService) *BookingSlotController {
	return &BookingSlotController{service: s}
}

func (c *BookingSlotController) GetAll(ctx *gin.Context) {
	slots, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, apix.HTTPResponse{
			Message: "failed to get booking slots",
			Data:    err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, apix.HTTPResponse{
		Message: "get all booking slots successfully",
		Data:    slots,
	})
}

func (c *BookingSlotController) GetByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id64, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "invalid id param",
			Data:    err.Error(),
		})
		return
	}
	id := uint(id64)

	slot, err := c.service.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, apix.HTTPResponse{
			Message: "failed to get booking slot",
			Data:    err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, apix.HTTPResponse{
		Message: "get booking slot successfully",
		Data:    slot,
	})
}

func (c *BookingSlotController) Create(ctx *gin.Context) {
	var input dto.CreateBookingSlotDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ve, _ := validatorx.ParseValidatorErrors(err)
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "invalid input data",
			Data:    ve,
		})
		return
	}

	isAvailable := true
	if input.IsAvailable != nil {
		isAvailable = *input.IsAvailable
	}

	slot := model.BookingSlot{
		StartTime:   input.StartTime,
		EndTime:     input.EndTime,
		IsAvailable: isAvailable,
	}

	created, err := c.service.Create(&slot)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, apix.HTTPResponse{
			Message: "failed to create booking slot",
			Data:    err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, apix.HTTPResponse{
		Message: "booking slot created",
		Data:    created,
	})
}

func (c *BookingSlotController) Update(ctx *gin.Context) {
	var input dto.UpdateBookingSlotDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ve, _ := validatorx.ParseValidatorErrors(err)
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "invalid input data",
			Data:    ve,
		})
		return
	}

	slot := model.BookingSlot{
		ID:         input.ID,
		FacilityID: input.FacilityID,
		StartTime:  input.StartTime,
		EndTime:    input.EndTime,
	}

	if input.IsAvailable != nil {
		slot.IsAvailable = *input.IsAvailable
	}

	updated, err := c.service.Update(&slot)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, apix.HTTPResponse{
			Message: "failed to update booking slot",
			Data:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, apix.HTTPResponse{
		Message: "booking slot is updated",
		Data:    updated,
	})
}

func (c *BookingSlotController) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id64, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "invalid id param",
			Data:    err.Error(),
		})
		return
	}
	id := uint(id64)

	if err := c.service.Delete(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, apix.HTTPResponse{
			Message: "failed to delete booking slot",
			Data:    err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, apix.HTTPResponse{
		Message: "booking slot deleted",
	})
}

// Extra endpoint: GET /bookingslots/available?facility_id=123
func (c *BookingSlotController) GetAvailableByFacility(ctx *gin.Context) {
	facIDParam := ctx.Query("facility_id")
	if facIDParam == "" {
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "facility_id query param is required",
		})
		return
	}
	fid64, err := strconv.ParseUint(facIDParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "invalid facility_id param",
			Data:    err.Error(),
		})
		return
	}
	fid := uint(fid64)

	slots, err := c.service.FindAvailableByFacility(fid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, apix.HTTPResponse{
			Message: "failed to get available slots",
			Data:    err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, apix.HTTPResponse{
		Message: "available slots fetched",
		Data:    slots,
	})
}
