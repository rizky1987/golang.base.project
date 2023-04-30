package requests

type (
	CreateFloorRequest struct {
		RoomTypeId string `form:"roomTypeId" validate:"required"`
		Number     int    `form:"number" validate:"required"`
		Price      int    `form:"price" validate:"required"`
	}
)
