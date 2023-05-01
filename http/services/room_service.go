package services

import (
	"hotel/commonHelpers"
	"hotel/config/env"
	"hotel/http/helpers"
	"hotel/http/interfaces"
	"hotel/http/requests"
	"hotel/http/responses"
	"runtime"

	"github.com/labstack/echo/v4"
	_ "github.com/swaggo/echo-swagger"
	"gorm.io/gorm"

	SQLEntity "hotel/databases/entities/sql"
)

type RoomHandler struct {
	DB             *gorm.DB
	Helper         helpers.HTTPHelper
	Config         env.Config
	RoomRepository interfaces.RoomInterface
}

func (_h *RoomHandler) CreateHandler(ctx echo.Context) error {
	var (
		err   error
		input requests.CreateRoomRequest
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

	dataCurrentRoom, err := _h.RoomRepository.GetRoomByCode(input.Code)
	if err != nil {

		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendDatabaseError(ctx, err.Error(), fileLocation, fileLine)
	}

	if dataCurrentRoom != nil {

		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendDuplicateError(ctx, "Room", input.Code, fileLocation, fileLine)
	}

	floorId, err := commonHelpers.StringToNewUUID(input.FloorId)
	if err != nil {
		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendValidationError(ctx, err.Error(), fileLocation, fileLine)
	}

	room := SQLEntity.Room{
		FloorId:     floorId,
		Code:        input.Code,
		Number:      input.Number,
		CreatedBy:   sessionData.Username,
		CreatedDate: commonHelpers.GetCurrentTimeAsiaJakarta(),
	}

	_, err = _h.RoomRepository.Create(nil, room)
	if err != nil {

		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendDatabaseError(ctx, err.Error(), fileLocation, fileLine)
	}

	return _h.Helper.SendSuccess(ctx, "Create", "Room", "Code", input.Code)
}

func (_h *RoomHandler) GetAvailibilityRoom(ctx echo.Context) error {
	var (
		err error
	)

	_, err = _h.Helper.ValidateCMSJWTData(ctx)
	if err != nil {
		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendUnauthorizedError(ctx, err.Error(), fileLocation, fileLine)
	}

	startDate := ctx.QueryParam("startDate")

	_, err = commonHelpers.ConvertStringToDateFormatYYYYMMDD(ctx.QueryParam("startDate"))
	if err != nil {
		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendValidationError(ctx, err.Error(), fileLocation, fileLine)
	}

	endDate := ctx.QueryParam("endDate")

	_, err = commonHelpers.ConvertStringToDateFormatYYYYMMDD(ctx.QueryParam("endDate"))
	if err != nil {
		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendValidationError(ctx, err.Error(), fileLocation, fileLine)
	}

	floorNumber, err := commonHelpers.ConvertStringToInteger(ctx.QueryParam("floorNumber"))
	if err != nil {
		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendValidationError(ctx, err.Error(), fileLocation, fileLine)
	}

	roomNumber, err := commonHelpers.ConvertStringToInteger(ctx.QueryParam("roomNumber"))
	if err != nil {
		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendValidationError(ctx, err.Error(), fileLocation, fileLine)
	}

	roomTypeName := ctx.QueryParam("roomTypeName")
	startFloorPrice, err := commonHelpers.ConvertStringToInteger(ctx.QueryParam("startFloorPrice"))
	if err != nil {
		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendValidationError(ctx, err.Error(), fileLocation, fileLine)
	}

	endfloorPrice, err := commonHelpers.ConvertStringToInteger(ctx.QueryParam("endfloorPrice"))
	if err != nil {
		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendValidationError(ctx, err.Error(), fileLocation, fileLine)
	}

	tempAvaibilityRooms, err := _h.RoomRepository.GetAvailibilityRooms(
		startDate, endDate, floorNumber, roomNumber, roomTypeName, startFloorPrice, endfloorPrice)
	if err != nil {

		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendDatabaseError(ctx, err.Error(), fileLocation, fileLine)
	}
	results := responses.ConvertAvailibilityRoomEntityToResponse(tempAvaibilityRooms)
	return _h.Helper.SendAllDataSuccess(ctx, "Success", results)
}
