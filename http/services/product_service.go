package services

import (
	"example/config/env"
	"example/http/helpers"
	"example/http/interfaces"
	"example/http/requests"
	"runtime"

	"github.com/labstack/echo/v4"
	_ "github.com/swaggo/echo-swagger"
	"gorm.io/gorm"

	SQLEntity "example/databases/entities/sql"
)

type ProductHandler struct {
	DB                *gorm.DB
	Helper            helpers.HTTPHelper
	Config            env.Config
	ProductRepository interfaces.ProductInterface
}

func (_h *ProductHandler) CreateHandler(ctx echo.Context) error {
	var (
		err   error
		input requests.CreateProductRequest
	)

	_, err = _h.Helper.ValidateCMSJWTData(ctx)
	if err != nil {
		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendUnauthorizedError(ctx, err.Error(), fileLocation, fileLine)
	}

	err = ctx.Bind(&input)
	if err != nil {
		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendValidationError(ctx, err.Error(), fileLocation, fileLine)
	}

	if err = _h.Helper.Validate.Struct(input); err != nil {
		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendValidationError(ctx, err.Error(), fileLocation, fileLine)
	}

	dataCurrentProduct, err := _h.ProductRepository.GetProductByCode(input.ProductCode)
	if err != nil {

		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendDatabaseError(ctx, err.Error(), fileLocation, fileLine)
	}

	if dataCurrentProduct != nil {

		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendDuplicateError(ctx, "Product", dataCurrentProduct.ProductCode, fileLocation, fileLine)
	}

	product := SQLEntity.Product{
		ProductCode:            input.ProductCode,
		DosageDescription:      &input.DosageDescription,
		UsabilityDescription:   &input.UsabilityDescription,
		CompositionDescription: &input.CompositionDescription,
		HowToUseDescription:    &input.HowToUseDescription,
	}

	_, err = _h.ProductRepository.Create(nil, product)
	if err != nil {

		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendDatabaseError(ctx, err.Error(), fileLocation, fileLine)
	}

	return _h.Helper.SendSuccess(ctx, "Create", "Product", "Code", " product.ProductCode")
}
