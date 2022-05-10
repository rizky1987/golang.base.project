package helpers

import (
	"fmt"
	"net/http"
	"reflect"

	response "example/http/responses"
	"example/utils"

	ut "github.com/go-playground/universal-translator"
	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

type ServerResponse struct {
	Code int
	Type string
}

var (
	SuccessServerResponse                      ServerResponse = ServerResponse{200, "success"}
	BadRequestErrorServerResponse              ServerResponse = ServerResponse{400, "bad_request"}
	UnauthorizedErrorServerResponse            ServerResponse = ServerResponse{401, "unauthorized"}
	DatabaseErrorServerResponse                ServerResponse = ServerResponse{402, "database_error"}
	ForbiddenErrorServerResponse               ServerResponse = ServerResponse{403, "forbidden"}
	NotFoundServerResponse                     ServerResponse = ServerResponse{404, "not_found"}
	RequestTimeOutServerResponse               ServerResponse = ServerResponse{408, "request_time_out"}
	InternalServerErrorServerResponse          ServerResponse = ServerResponse{500, "internal_server_error"}
	NotImplementedServerResponse               ServerResponse = ServerResponse{501, "not_implemented"}
	ServiceTemporarilyOverloadedServerResponse ServerResponse = ServerResponse{502, "service_temporarily_overloaded"}
	ServiceUnavailableServerResponse           ServerResponse = ServerResponse{503, "service_unavailable"}
	DuplicateErrorServerResponse               ServerResponse = ServerResponse{400, "duplicate_data"}
)

// HTTPHelper ...
type HTTPHelper struct {
	Validate   *validator.Validate
	Translator ut.Translator
}

func (u *HTTPHelper) getTypeData(i interface{}) string {
	v := reflect.ValueOf(i)
	v = reflect.Indirect(v)

	return v.Type().String()
}

// GetStatusCode ...
func (u *HTTPHelper) GetStatusCode(err error) int {
	statusCode := http.StatusOK
	if err != nil {
		switch u.getTypeData(err) {
		case "models.ErrorUnauthorized":
			statusCode = http.StatusUnauthorized
		case "models.ErrorNotFound":
			statusCode = http.StatusNotFound
		case "models.ErrorConflict":
			statusCode = http.StatusConflict
		case "models.ErrorInternalServer":
			statusCode = http.StatusInternalServerError
		default:
			statusCode = http.StatusInternalServerError
		}
	}

	return statusCode
}

// SetResponse ...
// Set response data.
func (u *HTTPHelper) SetCommonResponse(c echo.Context, serverResponse ServerResponse, innerMessage, fileLocation string, fileLine int) response.CommonBaseResponse {

	result := response.CommonBaseResponse{

		Alert: response.AlertResponse{
			Code:         serverResponse.Code,
			Message:      serverResponse.Type,
			InnerMessage: innerMessage,
		},
	}

	if serverResponse.Code != 200 {
		errorCode := utils.GenerateErrorCode()
		result.Alert.InnerMessage = fmt.Sprintf("%s || %s", errorCode, innerMessage)

		utils.SaveErrorToApplicationInsight(errorCode, serverResponse.Type, innerMessage, fileLocation, fileLine)
	}

	return result
}

// SendError ...
// Send error response to consumers.
// func (u *HTTPHelper) SendError(c echo.Context, errMessages []string) error {
// 	res := u.SetResponse(c, errMessages, nil, ServiceUnavailableServerResponse)

// 	return u.SendResponse(res)
// }

// SendBadRequest ...
// Send bad request response to consumers.
// func (u *HTTPHelper) SendBadRequest(c echo.Context, errMessages []string) error {

// 	res := u.SetResponse(c, errMessages, nil, BadRequestErrorServerResponse)

// 	return u.SendResponse(res)
// }

// SendNotFoundRequest ...
// Send bad request response to consumers.
// func (u *HTTPHelper) SendNotFoundRequest(c echo.Context, errMessages []string) error {

// 	res := u.SetResponse(c, errMessages, nil, NotFoundServerResponse)

// 	return u.SendResponse(res)
// }

// SendUnauthorizedError ...
// Send unauthorized response to consumers.
func (u *HTTPHelper) SendUnauthorizedError(c echo.Context, errorMessage string, fileLocation string, fileLine int) error {

	res := u.SetCommonResponse(c, UnauthorizedErrorServerResponse, errorMessage, fileLocation, fileLine)

	return c.JSON(res.Alert.Code, res)
}

// SendSuccess ...
func (u *HTTPHelper) SendSuccess(c echo.Context, proccessType, entityName, entityField, entityData string) error {

	innerMessage := fmt.Sprintf(SuccessCrudMessage, proccessType, entityName, entityField, entityData)
	res := u.SetCommonResponse(c, SuccessServerResponse, innerMessage, "", 0)
	return c.JSON(res.Alert.Code, res)
}

// SendValidationError ...
// Send validation error response to consumers.
func (u *HTTPHelper) SendValidationError(c echo.Context, errorMessage, fileLocation string, fileLine int) error {
	res := u.SetCommonResponse(c, BadRequestErrorServerResponse, errorMessage, fileLocation, fileLine)

	return c.JSON(res.Alert.Code, res)
}

// SendDatabaseError ...
// Send database error response to consumers.
func (u *HTTPHelper) SendDatabaseError(c echo.Context, errorMessage, fileLocation string, fileLine int) error {
	res := u.SetCommonResponse(c, DatabaseErrorServerResponse, errorMessage, fileLocation, fileLine)

	return c.JSON(res.Alert.Code, res)
}

// SendDatabaseError ...
// Send database error response to consumers.
func (u *HTTPHelper) SendDuplicateError(c echo.Context, entityName, entityData, fileLocation string, fileLine int) error {
	errMessage := fmt.Sprintf(ErrorDuplicateMessage, entityName, entityData)

	res := u.SetCommonResponse(c, DuplicateErrorServerResponse, errMessage, fileLocation, fileLine)

	return c.JSON(res.Alert.Code, res)
}

// Error Middleware
//Error Middleware
func SendErrorMiddleware(c echo.Context, message string, serverResponse ServerResponse) error {
	return c.JSON(serverResponse.Code, map[string]interface{}{
		"code":      serverResponse.Code,
		"code_type": serverResponse.Type,
		"message":   []string{message},
	})
}

func (u *HTTPHelper) EmptyJsonMap() map[string]interface{} {
	return nil //make(map[string]interface{})
}
