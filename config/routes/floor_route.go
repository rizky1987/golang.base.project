package routes

import (
	"hotel/config/env"
	httpHelper "hotel/http/helpers"
	"hotel/http/repositories"
	"hotel/http/services"

	echo "github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterFloorRoutes(baseEndpointGroup *echo.Group, db *gorm.DB, config env.Config, httpHelper httpHelper.HTTPHelper) {

	group := baseEndpointGroup.Group("/floor")
	{
		floorHandler := services.FloorHandler{
			Config:          config,
			Helper:          httpHelper,
			FloorRepository: repositories.NewFloorRepository(db),
			DB:              db,
		}

		createFloor(group, floorHandler)
	}
}

// @Tags Floor
// @Description Floor Create
// @ID FloorCreate
// @Accept multipart/form-data
// @param Authorization header string true "Bearer %"
// @Param number formData int true "number"
// @Success 200 {object} responses.CommonBaseResponse
// @Router /api/cms/v1/floor/create [post]
func createFloor(baseEndpointGroup *echo.Group, floorHandler services.FloorHandler) {
	baseEndpointGroup.POST("/create", floorHandler.CreateHandler)
}
