package routes

import (
	"example/config/env"
	httpHelper "example/http/helpers"
	"example/http/repositories"
	"example/http/services"

	"gopkg.in/mgo.v2"

	echo "github.com/labstack/echo/v4"
)

func RegisterCartRoutes(baseEndpointGroup *echo.Group,
	mongoDBSession *mgo.Session, databaseName string, config env.Config, httpHelper httpHelper.HTTPHelper) {

	group := baseEndpointGroup.Group("/cart")
	{
		cartHandler := services.CartHandler{
			Config:         config,
			Helper:         httpHelper,
			CartRepository: repositories.NewCartRepository(mongoDBSession, databaseName),
		}

		getAllCart(group, cartHandler)
	}
}

// @Tags Cart
// @Description Get all Cart
// @ID GetAllCart
// @Success 200 {object} responses.CartSuccessResponse
// @Failure 400 {object} responses.CartFailedResponse
// @Router /api/cms/v1/cart/get-data [post]
func getAllCart(baseEndpointGroup *echo.Group, cartHandler services.CartHandler) {
	baseEndpointGroup.POST("/get-data", cartHandler.GetAllCart)
}
