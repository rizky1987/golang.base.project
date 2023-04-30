package middleware

import (
	"fmt"
	"runtime"

	"hotel/utils"

	"github.com/labstack/echo/v4"
)

func MiddlewareLogging(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {

			if err := recover(); err != nil {
				buf := make([]byte, 2048)
				n := runtime.Stack(buf, false)
				buf = buf[:n]

				errorCode := utils.GenerateErrorCode()
				errMessage := fmt.Sprintf("recovering from err %v\n %s", err, buf)

				utils.SaveErrorToApplicationInsight(errorCode, "unexpected_error", errMessage, "", 0)

				c.JSON(403, map[string]interface{}{
					"code":      403,
					"code_type": "unexpected_error",
					"message":   errorCode,
				})
			}
		}()
		return next(c)
	}
}
