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

type RoomPriceHandler struct {
	DB                  *gorm.DB
	Helper              helpers.HTTPHelper
	Config              env.Config
	RoomPriceRepository interfaces.RoomPriceInterface
}

func (_h *RoomPriceHandler) CreateHandler(ctx echo.Context) error {
	var (
		err   error
		input requests.CreateRoomPriceRequest
	)

	sessionData, err := _h.Helper.ValidateCMSJWTData(ctx)
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

	floorId, err := commonHelpers.StringToNewUUID(input.FloorId)
	if err != nil {
		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendValidationError(ctx, err.Error(), fileLocation, fileLine)
	}

	dataCurrentRoomPrice, err := _h.RoomPriceRepository.GetRoomPriceByCode(input.Code)
	if err != nil {

		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendDatabaseError(ctx, err.Error(), fileLocation, fileLine)
	}

	if dataCurrentRoomPrice != nil {

		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendDuplicateError(ctx, "RoomPrice", input.Code, fileLocation, fileLine)
	}

	roomPrice := SQLEntity.RoomPrice{
		Code:        input.Code,
		Type:        input.Type,
		FloorId:     floorId,
		Price:       input.Price,
		CreatedBy:   sessionData.Username,
		CreatedDate: commonHelpers.GetCurrentTimeAsiaJakarta(),
	}

	_, err = _h.RoomPriceRepository.Create(nil, roomPrice)
	if err != nil {

		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendDatabaseError(ctx, err.Error(), fileLocation, fileLine)
	}

	return _h.Helper.SendSuccess(ctx, "Create", "RoomPrice", "Code", input.Code)
}
