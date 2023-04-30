package routes

import (
	"hotel/config/env"
	httpHelper "hotel/http/helpers"
	"hotel/http/repositories"
	"hotel/http/services"

	echo "github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterBookingRoutes(baseEndpointGroup *echo.Group, db *gorm.DB, config env.Config, httpHelper httpHelper.HTTPHelper) {

	group := baseEndpointGroup.Group("/booking")
	{
		roomHandler := services.BookingHandler{
			Config:                  config,
			Helper:                  httpHelper,
			BookingRepository:       repositories.NewBookingRepository(db),
			BookingDetailRepository: repositories.NewBookingDetailRepository(db),
			RoomRepository:          repositories.NewRoomRepository(db),
			DB:                      db,
		}

		createBooking(group, roomHandler)
	}
}

// @Tags Booking
// @Description Booking Create
// @ID BookingCreate
// @Accept application/json
// @param Authorization header string true "Bearer %"
// @Param body body requests.CreateBookingRequest true "body"
// @Success 200 {object} responses.CommonBaseResponse
// @Router /api/cms/v1/booking/create [post]
func createBooking(baseEndpointGroup *echo.Group, roomHandler services.BookingHandler) {
	baseEndpointGroup.POST("/create", roomHandler.CreateHandler)
}
