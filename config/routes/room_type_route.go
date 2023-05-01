package routes

import (
	"hotel/config/env"
	httpHelper "hotel/http/helpers"
	"hotel/http/repositories"
	"hotel/http/services"

	echo "github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterRoomPriceRoutes(baseEndpointGroup *echo.Group, db *gorm.DB, config env.Config, httpHelper httpHelper.HTTPHelper) {

	group := baseEndpointGroup.Group("/room-price")
	{
		roomTypeHandler := services.RoomPriceHandler{
			Config:              config,
			Helper:              httpHelper,
			RoomPriceRepository: repositories.NewRoomPriceRepository(db),
			DB:                  db,
		}

		createRoomPrice(group, roomTypeHandler)
	}
}

// @Tags RoomPrice
// @Description RoomPrice Create
// @ID RoomPriceCreate
// @Accept multipart/form-data
// @param Authorization header string true "Bearer %"
// @Param code formData string true "code"
// @Param type formData string true "type"
// @Param price formData string true "price"
// @Param floorId formData string true "floorId"
// @Success 200 {object} responses.CommonBaseResponse
// @Router /api/cms/v1/room-price/create [post]
func createRoomPrice(baseEndpointGroup *echo.Group, roomTypeHandler services.RoomPriceHandler) {
	baseEndpointGroup.POST("/create", roomTypeHandler.CreateHandler)
}
