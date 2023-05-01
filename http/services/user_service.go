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
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	SQLEntity "hotel/databases/entities/sql"
)

type UserHandler struct {
	DB             *gorm.DB
	Helper         helpers.HTTPHelper
	Config         env.Config
	UserRepository interfaces.UserInterface
}

func (_h *UserHandler) CreateHandler(ctx echo.Context) error {
	var (
		err   error
		input requests.CreateUserRequest
	)

	err = ctx.Bind(&input)
	if err != nil {
		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendValidationError(ctx, err.Error(), fileLocation, fileLine)
	}

	if err = _h.Helper.Validate.Struct(input); err != nil {
		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendValidationError(ctx, err.Error(), fileLocation, fileLine)
	}

	dataCurrentUser, err := _h.UserRepository.GetUserByUsername(input.Username)
	if err != nil {

		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendDatabaseError(ctx, err.Error(), fileLocation, fileLine)
	}

	if dataCurrentUser != nil {

		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendDuplicateError(ctx, "User", input.Username, fileLocation, fileLine)
	}

	password, err := commonHelpers.ConvertStringToBCrypt(input.Password)
	if err != nil {
		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendValidationError(ctx, err.Error(), fileLocation, fileLine)
	}

	user := SQLEntity.User{
		Username: input.Username,
		Password: password,
	}

	_, err = _h.UserRepository.Create(nil, user)
	if err != nil {

		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendDatabaseError(ctx, err.Error(), fileLocation, fileLine)
	}

	return _h.Helper.SendSuccess(ctx, "Create", "User", "Username", input.Username)
}

func (_h *UserHandler) LoginHandler(ctx echo.Context) error {
	var (
		err   error
		input requests.CreateUserRequest
	)

	err = ctx.Bind(&input)
	if err != nil {
		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendValidationError(ctx, err.Error(), fileLocation, fileLine)
	}

	if err = _h.Helper.Validate.Struct(input); err != nil {
		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendValidationError(ctx, err.Error(), fileLocation, fileLine)
	}

	dataCurrentUser, err := _h.UserRepository.GetUserByUsername(input.Username)
	if err != nil {

		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendDatabaseError(ctx, err.Error(), fileLocation, fileLine)
	}

	if err = bcrypt.CompareHashAndPassword(
		[]byte(commonHelpers.TrimWhiteSpace(dataCurrentUser.Password)),
		[]byte(commonHelpers.TrimWhiteSpace(input.Password))); err != nil {

		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendDatabaseError(ctx, "Username or Password does not match.", fileLocation, fileLine)
	}

	if dataCurrentUser == nil {

		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendValidationError(ctx, "User not found", fileLocation, fileLine)
	}

	token, err := helpers.GenetarateToken(dataCurrentUser)
	if err != nil {
		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return _h.Helper.SendValidationError(ctx, err.Error(), fileLocation, fileLine)
	}

	return _h.Helper.SendAllDataSuccess(ctx, "User", token)
}
