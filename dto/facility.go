package dto

type CreateFacilityDTO struct {
	Name      string  `json:"name,ommitempty" binding:"required"`
	Price     float32 `json:"price,omitempty" binding:"required"`
	Capacity  uint    `json:"capacity,omitempty" binding:"required"`
	Available bool    `json:"available,omitempty" binding:"required"`
}

type UpdateFacilityDTO struct {
	Name      string  `json:"name,ommitempty" binding:"required"`
	Price     float32 `json:"price,omitempty" binding:"required"`
	Capacity  uint    `json:"capacity,omitempty" binding:"required"`
	Available bool    `json:"available,omitempty" binding:"required"`
}
