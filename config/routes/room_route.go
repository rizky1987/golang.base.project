package routes

import (
	"hotel/config/env"
	httpHelper "hotel/http/helpers"
	"hotel/http/repositories"
	"hotel/http/services"

	echo "github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterRoomRoutes(baseEndpointGroup *echo.Group, db *gorm.DB, config env.Config, httpHelper httpHelper.HTTPHelper) {

	group := baseEndpointGroup.Group("/room")
	{
		roomHandler := services.RoomHandler{
			Config:         config,
			Helper:         httpHelper,
			RoomRepository: repositories.NewRoomRepository(db),
			DB:             db,
		}

		createRoom(group, roomHandler)
		getAvailibilityRoom(group, roomHandler)
	}
}

// @Tags Room
// @Description Room Create
// @ID RoomCreate
// @Accept multipart/form-data
// @param Authorization header string true "Bearer %"
// @Param roomPriceId formData string true "roomPriceId"
// @Param code formData string true "code"
// @Param number formData int true "number"
// @Success 200 {object} responses.CommonBaseResponse
// @Router /api/cms/v1/room/create [post]
func createRoom(baseEndpointGroup *echo.Group, roomHandler services.RoomHandler) {
	baseEndpointGroup.POST("/create", roomHandler.CreateHandler)
}

// @Tags Room
// @Description Room Get Availibility Room
// @ID RoomGetAvailibilityRoom
// @Accept multipart/form-data
// @param Authorization header string true "Bearer %"
// @Param startDate query string false "startDate"
// @Param endDate query string false "endDate"
// @Param floorNumber query int false "floorNumber"
// @Param roomNumber query int false "roomNumber"
// @Param roomTypeName query string false "roomTypeName"
// @Param startPrice query int false "startFloorPrice"
// @Param endPrice query int false "endfloorPrice"
// @Success 200 {object} responses.CommonBaseResponse
// @Router /api/cms/v1/room/availibility-room [get]
func getAvailibilityRoom(baseEndpointGroup *echo.Group, roomHandler services.RoomHandler) {
	baseEndpointGroup.GET("/availibility-room", roomHandler.GetAvailibilityRoom)
}
