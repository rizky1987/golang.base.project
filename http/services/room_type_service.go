package services

import (
	"hotel/commonHelpers"
	"hotel/config/env"
	"hotel/http/helpers"
	"hotel/http/interfaces"
	"hotel/http/requests"
	"runtime"

	"github.com/labstack/echo/v4"
	_ "github.com/swaggo/echo-swagger"
	"gorm.io/gorm"

	SQLEntity "hotel/databases/entities/sql"
)

type RoomTypeHandler struct {
	DB                 *gorm.DB
	Helper             helpers.HTTPHelper
	Config             env.Config
	RoomTypeRepository interfaces.RoomTypeInterface
}

func (_h *RoomTypeHandler) CreateHandler(ctx echo.Context) error {
	var (
		err   error
		input requests.CreateRoomTypeRequest
	)

	// sessionData, err := _h.Helper.ValidateCMSJWTData(ctx)
	// if err != nil {
	// 	_, fileLocation, fileLine, _ := runtime.Caller(0)
	// 	return _h.Helper.SendUnauthorizedError(ctx, err.Error(), fileLocation, fileLine)
	// }

	err = ctx.Bind(&input)
	if err != nil {
		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendValidationError(ctx, err.Error(), fileLocation, fileLine)
	}

	if err = _h.Helper.Validate.Struct(input); err != nil {
		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendValidationError(ctx, err.Error(), fileLocation, fileLine)
	}

	dataCurrentRoomType, err := _h.RoomTypeRepository.GetRoomTypeByCode(input.Code)
	if err != nil {

		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendDatabaseError(ctx, err.Error(), fileLocation, fileLine)
	}

	if dataCurrentRoomType != nil {

		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendDuplicateError(ctx, "RoomType", input.Code, fileLocation, fileLine)
	}

	roomType := SQLEntity.RoomType{
		Code:        input.Code,
		Name:        input.Name,
		CreatedBy:   "test",
		CreatedDate: commonHelpers.GetCurrentTimeAsiaJakarta(),
	}

	_, err = _h.RoomTypeRepository.Create(nil, roomType)
	if err != nil {

		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendDatabaseError(ctx, err.Error(), fileLocation, fileLine)
	}

	return _h.Helper.SendSuccess(ctx, "Create", "RoomType", "Code", input.Code)
}
