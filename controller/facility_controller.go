package controller

import (
	"net/http"

	"bookit.com/dto"
	"bookit.com/model"

	apix "bookit.com/utils/api"
	validatorx "bookit.com/utils/validator"

	"github.com/gin-gonic/gin"
)

type FacilityService interface {
	GetAll() ([]model.Facility, error)
	GetByID(id uint) (*model.Facility, error)
	Create(facility *model.Facility) (*model.Facility, error)
	Update(facility *model.Facility) (*model.Facility, error)
}

type FacilityController struct {
	facilityService FacilityService
}

func NewFacilityController(facilityService FacilityService) *FacilityController {
	return &FacilityController{
		facilityService: facilityService,
	}
}

func (c *FacilityController) GetAll(ctx *gin.Context) {
	var facilities []model.Facility
	var err error

	facilities, err = c.facilityService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, apix.HTTPResponse{
			Message: "failed to get facility",
			Data:    err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, apix.HTTPResponse{
		Message: "get all facility successfully",
		Data:    facilities,
	})
}

func (c *FacilityController) Create(ctx *gin.Context) {
	var input dto.CreateFacilityDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ve, _ := validatorx.ParseValidatorErrors(err)
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "invalid input data",
			Data:    ve,
		})
		return
	}

	facility := model.Facility{
		Name:     input.Name,
		Price:    input.Price,
		Capacity: input.Capacity,
		//Available: input.Available,
	}

	created, err := c.facilityService.Create(&facility)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, apix.HTTPResponse{
			Message: "failed to create facility",
			Data:    err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, created)

}

func (c *FacilityController) Update(ctx *gin.Context) {
	var input dto.UpdateFacilityDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ve, _ := validatorx.ParseValidatorErrors(err)
		ctx.JSON(http.StatusBadRequest, apix.HTTPResponse{
			Message: "invalid input data",
			Data:    ve,
		})
		return
	}

	facility := model.Facility{
		Name:     input.Name,
		Price:    input.Price,
		Capacity: input.Capacity,
		//Available: input.Available,
	}

	updated, err := c.facilityService.Update(&facility)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, apix.HTTPResponse{
			Message: "failed to update facility",
			Data:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, apix.HTTPResponse{
		Message: "facility is updated",
		Data:    updated,
	})

}
