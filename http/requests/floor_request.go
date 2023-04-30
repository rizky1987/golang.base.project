package requests

type (
	CreateFloorRequest struct {
		RoomTypeId string `form:"roomTypeId" validate:"required"`
		Number     int    `form:"number"`
		Price      int    `form:"usabilityDescription"`
	}
)
