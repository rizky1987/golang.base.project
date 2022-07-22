package main

import (
	"example/commonHelpers"
	"example/config/boot"
	cfg "example/config/env"
	"fmt"

	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
	enTranslations "gopkg.in/go-playground/validator.v9/translations/en"
)

type App struct {
	config cfg.Config
}

var app App

var (
	validatorDriver *validator.Validate
	uni             *ut.UniversalTranslator
	translator      ut.Translator
)

func init() {
	config := cfg.NewViperConfig()
	app = App{config: config}

}

func main() {

	defer catch()

	registerValidator()
	e := echo.New()
	apiHandler := boot.HTTPHandler{
		E:               e,
		Config:          app.config,
		ValidatorDriver: validatorDriver,
		Translator:      translator,
	}

	apiHandler.RegisterApiHandler()

	host := fmt.Sprintf("%s:%s", commonHelpers.GetConfigurationStringValue(`app.host`), commonHelpers.GetConfigurationStringValue(`app.port`))
	e.Start(host)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":" + commonHelpers.GetConfigurationStringValue(`app.port`)))

}

func catch() {
	if r := recover(); r != nil {
		fmt.Println("Error occured", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}

func registerValidator() {
	en := en.New()
	uni = ut.New(en, en)

	trans, _ := uni.GetTranslator("en")
	translator = trans

	validatorDriver = validator.New()
	enTranslations.RegisterDefaultTranslations(validatorDriver, translator)
}
