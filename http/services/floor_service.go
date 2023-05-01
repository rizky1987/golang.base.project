package services

import (
	"fmt"
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

type FloorHandler struct {
	DB              *gorm.DB
	Helper          helpers.HTTPHelper
	Config          env.Config
	FloorRepository interfaces.FloorInterface
}

func (_h *FloorHandler) CreateHandler(ctx echo.Context) error {
	var (
		err   error
		input requests.CreateFloorRequest
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

	dataCurrentFloor, err := _h.FloorRepository.GetFloorByNumber(input.Number)
	if err != nil {

		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendDatabaseError(ctx, err.Error(), fileLocation, fileLine)
	}

	if dataCurrentFloor != nil {

		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendDuplicateError(ctx, "Floor", fmt.Sprintf("%d", dataCurrentFloor.Number), fileLocation, fileLine)
	}

	roomTypeId, err := commonHelpers.StringToNewUUID(input.RoomTypeId)
	if err != nil {
		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendValidationError(ctx, err.Error(), fileLocation, fileLine)
	}

	floor := SQLEntity.Floor{
		RoomTypeId:  roomTypeId,
		Number:      input.Number,
		Price:       input.Price,
		CreatedBy:   sessionData.Username,
		CreatedDate: commonHelpers.GetCurrentTimeAsiaJakarta(),
	}

	_, err = _h.FloorRepository.Create(nil, floor)
	if err != nil {

		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendDatabaseError(ctx, err.Error(), fileLocation, fileLine)
	}

	return _h.Helper.SendSuccess(ctx, "Create", "Floor", "Number", fmt.Sprintf("%d", input.Number))
}
