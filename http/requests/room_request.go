package requests

type (
	CreateRoomRequest struct {
		RoomPriceId string `form:"roomPriceId" validate:"required"`
		Code        string `form:"code" validate:"required"`
		Number      int    `form:"number" validate:"required"`
	}
)
