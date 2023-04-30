package requests

type (
	CreateRoomTypeRequest struct {
		Code string `form:"code" validate:"required"`
		Name string `form:"name" validate:"required"`
	}
)
