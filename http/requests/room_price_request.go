package requests

type (
	CreateRoomPriceRequest struct {
		Code    string `form:"code" validate:"required"`
		Type    string `form:"type" validate:"required"`
		Price   string `form:"price" validate:"required"`
		FloorId string `form:"floorId" validate:"required"`
	}
)
