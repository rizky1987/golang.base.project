package requests

type (
	CreateUserRequest struct {
		Username string `form:"username" validate:"required"`
		Password string `form:"password" validate:"required"`
	}
)
