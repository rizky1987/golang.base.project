package services

import (
	"errors"
	"fmt"
	"hotel/commonHelpers"
	"hotel/config/env"
	"hotel/http/helpers"
	"hotel/http/interfaces"
	"hotel/http/requests"
	"runtime"
	"strings"
	"time"

	mssql "github.com/denisenkom/go-mssqldb"
	"github.com/labstack/echo/v4"
	_ "github.com/swaggo/echo-swagger"
	"gorm.io/gorm"

	SQLEntity "hotel/databases/entities/sql"
)

type BookingHandler struct {
	DB                      *gorm.DB
	Helper                  helpers.HTTPHelper
	Config                  env.Config
	BookingRepository       interfaces.BookingInterface
	BookingDetailRepository interfaces.BookingDetailInterface
	RoomRepository          interfaces.RoomInterface
}

func (_h *BookingHandler) CreateHandler(ctx echo.Context) error {
	var (
		err   error
		input requests.CreateBookingRequest
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

	startDate, err := commonHelpers.ConvertStringToDateTimeNullAble(input.StartDate)
	if err != nil {
		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendValidationError(ctx, err.Error(), fileLocation, fileLine)
	}

	endDate, err := commonHelpers.ConvertStringToDateTimeNullAble(input.EndDate)
	if err != nil {
		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendValidationError(ctx, err.Error(), fileLocation, fileLine)
	}

	bookingCode := "bookingCode1"
	dataCurrentBooking, err := _h.BookingRepository.GetBookingByCode(bookingCode)
	if err != nil {

		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendDatabaseError(ctx, err.Error(), fileLocation, fileLine)
	}

	if dataCurrentBooking != nil {

		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendDuplicateError(ctx, "Booking", bookingCode, fileLocation, fileLine)
	}

	if len(input.BookingDetails) < 1 {

		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendValidationError(ctx, "Please input at least one booked room", fileLocation, fileLine)
	}

	bookingEntity := &SQLEntity.Booking{
		ID:          commonHelpers.GenerateNewUUID(),
		Code:        bookingCode,
		BookedName:  input.BookedName,
		DownPayment: input.DownPayment,
		StartDate:   *startDate,
		EndDate:     *endDate,
		CreatedBy:   sessionData.Username,
		CreatedDate: commonHelpers.GetCurrentTimeAsiaJakarta(),
	}

	newStartDate, err, fileLocation, fileLine := _h.validateBookingDateAndTime(input.IsTimeRulesAgree,
		bookingEntity.StartDate, bookingEntity.EndDate)
	if err != nil {
		return _h.Helper.SendValidationError(ctx, err.Error(), fileLocation, fileLine)
	}

	bookingEntity.StartDate = *newStartDate

	bookingDetailEntities, totalPrice, err, fileLocation, fileLine := _h.validateRoomIds(bookingEntity.ID,
		input.StartDate, input.EndDate, input.BookingDetails)
	if err != nil {

		return _h.Helper.SendValidationError(ctx, err.Error(), fileLocation, fileLine)
	}

	if len(bookingDetailEntities) < 1 {

		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendValidationError(ctx, "Please input at least one valid booking detail data", fileLocation, fileLine)
	}

	if totalPrice < bookingEntity.DownPayment {
		_, fileLocation, fileLine, _ := runtime.Caller(0)

		errMessage := fmt.Sprintf("Downpayment is bigger than total price. DP %d, TotalPrice %d",
			bookingEntity.DownPayment, totalPrice)
		return _h.Helper.SendValidationError(ctx, errMessage, fileLocation, fileLine)
	}

	if totalPrice == bookingEntity.DownPayment {
		bookingEntity.IsPaidOff = true
	}

	transaction := _h.DB.Begin()
	err = _h.BookingRepository.Create(transaction, bookingEntity)
	if err != nil {

		transaction.Rollback()
		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendDatabaseError(ctx, err.Error(), fileLocation, fileLine)
	}

	err = _h.BookingDetailRepository.CreateBulk(transaction, bookingDetailEntities)
	if err != nil {

		transaction.Rollback()
		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendDatabaseError(ctx, err.Error(), fileLocation, fileLine)
	}

	transaction.Commit()

	return _h.Helper.SendSuccess(ctx, "Create", "Booking", "Code", bookingCode)
}

func (_h *BookingHandler) validateBookingDateAndTime(isTimeRulesAgree bool, startDate time.Time, endDate time.Time) (*time.Time, error, string, int) {

	currentDateTime := commonHelpers.GetCurrentTimeAsiaJakarta()

	currentHour := currentDateTime.Hour()
	realDate := startDate

	if startDate.Before(currentDateTime) {

		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return nil, errors.New("Can't booking with previous date"), fileLocation, fileLine
	}

	if endDate.Before(startDate) {

		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return nil, errors.New("startDate can't bigger than endDate"), fileLocation, fileLine
	}

	if currentHour < 6 {

		realDate = startDate.AddDate(0, 0, -1)
	} else if currentHour >= 6 && currentHour < 14 {

		if !isTimeRulesAgree {

			_, fileLocation, fileLine, _ := runtime.Caller(0)
			return nil, errors.New("You are not aggre with time rules"), fileLocation, fileLine
		} else {
			realDate = startDate.AddDate(0, 0, -1)
		}
	} else {
		realDate = startDate
	}

	return &realDate, nil, "", 0
}

func (_h *BookingHandler) validateRoomIds(bookingId mssql.UniqueIdentifier,
	bookingStartDate, bookingEndDate string,
	bookingDetails []*requests.CreateBookingDetailRequest) ([]*SQLEntity.BookingDetail, int, error, string, int) {

	var roomIds []string

	for i := 0; i < len(bookingDetails); i++ {

		roomId, err := commonHelpers.StringToNewUUID(bookingDetails[0].RoomId)
		if err != nil {

			_, fileLocation, fileLine, _ := runtime.Caller(0)
			return nil, 0, err, fileLocation, fileLine
		}

		roomIds = append(roomIds, roomId.String())
	}

	if len(roomIds) < 1 {
		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return nil, 0, errors.New("no valid room id found"), fileLocation, fileLine
	}

	tempBookingRoomAvaibilities, err := _h.BookingRepository.GetBookingRoomAvaibility(roomIds, bookingStartDate, bookingEndDate)
	if err != nil {
		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return nil, 0, err, fileLocation, fileLine
	}

	var bookingRoomAvaibilityErrorMessages []string
	if len(tempBookingRoomAvaibilities) > 0 {

		for l := 0; l < len(tempBookingRoomAvaibilities); l++ {
			bookingRoomAvaibilityErrorMessages = append(bookingRoomAvaibilityErrorMessages,
				fmt.Sprintf("room %s already booked with booking code : %s",
					tempBookingRoomAvaibilities[l].RoomCode,
					tempBookingRoomAvaibilities[l].BookingCode))
		}

		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return nil, 0, errors.New(strings.Join(bookingRoomAvaibilityErrorMessages, "<br/>")), fileLocation, fileLine
	}

	tempRoomDetails, err := _h.RoomRepository.GetRoomDetailByRoomIds(roomIds)
	if err != nil {
		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return nil, 0, err, fileLocation, fileLine
	}

	var bookingDetailEntities []*SQLEntity.BookingDetail
	var notFoundMessages []string
	totalPrice := 0
	for j := 0; j < len(roomIds); j++ {

		isFound := false
		var bookingDetail *SQLEntity.BookingDetail

		for k := 0; k < len(tempRoomDetails); k++ {

			if tempRoomDetails[k].RoomId.String() == roomIds[j] {

				isFound = true
				bookingDetail = &SQLEntity.BookingDetail{
					BookingId: bookingId,
					RoomId:    tempRoomDetails[k].RoomId,
					Price:     tempRoomDetails[k].FloorPrice,
				}

				totalPrice += bookingDetail.Price
				break
			}
		}

		if isFound && bookingDetail != nil {
			bookingDetailEntities = append(bookingDetailEntities, bookingDetail)
		} else {
			notFoundMessages = append(notFoundMessages, fmt.Sprintf("Room Id %s not found", roomIds[j]))
		}
	}

	if len(notFoundMessages) > 0 {
		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return nil, 0, errors.New(strings.Join(notFoundMessages, "\n")), fileLocation, fileLine
	}

	return bookingDetailEntities, totalPrice, nil, "", 0
}
