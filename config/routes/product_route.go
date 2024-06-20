package routes

import (
	"example/config/env"
	httpHelper "example/http/helpers"
	"example/http/repositories"
	"example/http/services"

	echo "github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterUserRoutes(baseEndpointGroup *echo.Group, db *gorm.DB, config env.Config, httpHelper httpHelper.HTTPHelper) {

	group := baseEndpointGroup.Group("/product")
	{
		productHandler := services.ProductHandler{
			Config:            config,
			Helper:            httpHelper,
			ProductRepository: repositories.NewProductRepository(db),
			DB:                db,
		}

		createProduct(group, productHandler)
	}
}

// @Tags Product
// @Description Product Create
// @ID ProductCreate
// @Accept multipart/form-data
// @Security Bearer
// @Param productCode formData string true "productCode"
// @Param dosageDescription formData string false "dosageDescription"
// @Param usabilityDescription formData string false "usabilityDescription"
// @Param composition formData string false "composition"
// @Param howToUseDescription formData string false "howToUseDescription"
// @Success 200 {object} responses.CommonBaseResponse
// @Router /api/cms/v1/product/create [post]
func createProduct(baseEndpointGroup *echo.Group, productHandler services.ProductHandler) {
	baseEndpointGroup.POST("/create", productHandler.CreateHandler)
}
