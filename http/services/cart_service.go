package services

import (
	"example/config/env"
	"example/http/helpers"
	"example/http/interfaces"
	"runtime"

	"example/http/responses"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CartHandler struct {
	DB             *gorm.DB
	Helper         helpers.HTTPHelper
	Config         env.Config
	CartRepository interfaces.CartInterface
}

func (_h *CartHandler) GetAllCart(ctx echo.Context) error {
	var (
		err error
	)

	cartEntities, err := _h.CartRepository.GetAllCart()
	if err != nil {
		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendDatabaseError(ctx, err.Error(), fileLocation, fileLine)
	}

	results := responses.ConvertListCartMongoEntityToCartResponseResponse(cartEntities)
	return _h.Helper.SendAllDataSuccess(ctx, "Cart", results)
}
