package routes

import (
	"hotel/config/env"
	httpHelper "hotel/http/helpers"
	"hotel/http/repositories"
	"hotel/http/services"

	echo "github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterUserRoutes(baseEndpointGroup *echo.Group, db *gorm.DB, config env.Config, httpHelper httpHelper.HTTPHelper) {

	group := baseEndpointGroup.Group("/user")
	{
		userTypeHandler := services.UserHandler{
			Config:         config,
			Helper:         httpHelper,
			UserRepository: repositories.NewUserRepository(db),
			DB:             db,
		}

		createUser(group, userTypeHandler)
		loginUser(group, userTypeHandler)
	}
}

// @Tags User
// @Description User Create
// @ID UserCreate
// @Accept multipart/form-data
// @Param username formData string true "username"
// @Param password formData string true "password"
// @Success 200 {object} responses.CommonBaseResponse
// @Router /api/cms/v1/user/create [post]
func createUser(baseEndpointGroup *echo.Group, userTypeHandler services.UserHandler) {
	baseEndpointGroup.POST("/create", userTypeHandler.CreateHandler)
}

// @Tags User
// @Description User Login
// @ID UserLogin
// @Accept multipart/form-data
// @Param username formData string true "username"
// @Param password formData string true "password"
// @Success 200 {object} responses.CommonBaseResponse
// @Router /api/cms/v1/user/login [post]
func loginUser(baseEndpointGroup *echo.Group, userTypeHandler services.UserHandler) {
	baseEndpointGroup.POST("/login", userTypeHandler.LoginHandler)
}
