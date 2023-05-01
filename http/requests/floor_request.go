package requests

type (
	CreateFloorRequest struct {
		Number int `form:"number" validate:"required"`
	}
)
