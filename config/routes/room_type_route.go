package routes

import (
	"hotel/config/env"
	httpHelper "hotel/http/helpers"
	"hotel/http/repositories"
	"hotel/http/services"

	echo "github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterRoomTypeRoutes(baseEndpointGroup *echo.Group, db *gorm.DB, config env.Config, httpHelper httpHelper.HTTPHelper) {

	group := baseEndpointGroup.Group("/room-type")
	{
		roomTypeHandler := services.RoomTypeHandler{
			Config:             config,
			Helper:             httpHelper,
			RoomTypeRepository: repositories.NewRoomTypeRepository(db),
			DB:                 db,
		}

		createRoomType(group, roomTypeHandler)
	}
}

// @Tags RoomType
// @Description RoomType Create
// @ID RoomTypeCreate
// @Accept multipart/form-data
// @param Authorization header string true "Bearer %"
// @Param code formData string true "code"
// @Param name formData string true "name"
// @Success 200 {object} responses.CommonBaseResponse
// @Router /api/cms/v1/room-type/create [post]
func createRoomType(baseEndpointGroup *echo.Group, roomTypeHandler services.RoomTypeHandler) {
	baseEndpointGroup.POST("/create", roomTypeHandler.CreateHandler)
}
