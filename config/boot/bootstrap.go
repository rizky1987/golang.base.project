package boot

import (
	"example/config/env"
	"example/docs"
	httpHelper "example/http/helpers"
	"fmt"

	customMiddleware "example/config/middleware"

	"example/databases/connection/mongo"

	ut "github.com/go-playground/universal-translator"
	"github.com/labstack/echo/v4"
	v4middleware "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	validator "gopkg.in/go-playground/validator.v9"
)

type HTTPHandler struct {
	E               *echo.Echo
	Config          env.Config
	Helper          httpHelper.HTTPHelper
	ValidatorDriver *validator.Validate
	Translator      ut.Translator
}

func (h *HTTPHandler) RegisterApiHandler() *HTTPHandler {

	h.Helper = httpHelper.HTTPHelper{
		Validate:   h.ValidatorDriver,
		Translator: h.Translator,
	}

	databaseHost := h.Config.GetString("database.mongo_db.host")
	databaseName := h.Config.GetString("database.mongo_db.database")
	databaseUser := h.Config.GetString("database.mongo_db.user")
	databasePassword := h.Config.GetString("database.mongo_db.password")
	databasePort := h.Config.GetString("database.mongo_db.port")

	// mongoDBSession, err := mongo.Connect(databaseHost, databaseName, databaseUser, databasePassword, databasePort)
	_, err := mongo.Connect(databaseHost, databaseName, databaseUser, databasePassword, databasePort)
	if err != nil {
		panic(err)
	}
	// End DB Connection

	host := fmt.Sprintf("%s:%s", h.Config.GetString(`app.host`), h.Config.GetString(`app.port`))
	//Begin Global Swagger Configuration
	h.E.GET("/swagger/*", echoSwagger.WrapHandler)
	docs.SwaggerInfo.Title = "API for Radya Labs"
	docs.SwaggerInfo.Description = "This is API from Radya Labs"
	docs.SwaggerInfo.Version = h.Config.GetString(`app.version`)
	docs.SwaggerInfo.Host = host
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	//End Global Swagger Configuration

	h.E.Use(customMiddleware.MiddlewareLogging)
	// Begin Register All End Point

	// End Register All End Point

	return h
}

func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Example/1.0")

		return next(c)
	}
}

// RegisterMiddleware ...
func (h *HTTPHandler) RegisterMiddleware() {

	h.E.Use(v4middleware.CORSWithConfig(v4middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	h.E.Use(v4middleware.GzipWithConfig(v4middleware.GzipConfig{
		Level: 5,
	}))

	if h.Config.GetBool("app.debug") == true {
		h.E.Use(v4middleware.Logger())
		h.E.HideBanner = true
		h.E.Debug = true
	} else {
		h.E.HideBanner = true
		h.E.Debug = false
		h.E.Use(v4middleware.Recover())
	}
}
