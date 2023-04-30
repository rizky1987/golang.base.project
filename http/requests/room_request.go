package requests

type (
	CreateRoomRequest struct {
		FloorId string `form:"floorId" validate:"required"`
		Code    string `form:"code" validate:"required"`
		Number  int    `form:"number" validate:"required"`
	}
)
