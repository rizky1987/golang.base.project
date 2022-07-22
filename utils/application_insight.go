package utils

import (
	"fmt"
	"os"
	"time"

	"example/commonHelpers"

	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

func SaveErrorToApplicationInsight(errorCode, errorType, errorMessage, fileLocation string, fileLine int) {

	instrumentationkey := commonHelpers.GetConfigurationStringValue("third_part.application_insight.instrumentationkey")
	telemetryConfig := appinsights.NewTelemetryConfiguration(instrumentationkey)

	client := appinsights.NewTelemetryClientFromConfig(telemetryConfig)

	errorFormating := fmt.Sprintf("%s || %s || %s || %s:%d",
		errorCode,
		errorType,
		errorMessage,
		fileLocation,
		fileLine)

	trace := appinsights.NewTraceTelemetry(errorFormating, appinsights.Error)
	trace.Timestamp = time.Now()
	client.Track(trace)

	isApplicationOnDebugMode := commonHelpers.GetConfigurationBoolValue("app.debug")

	if isApplicationOnDebugMode {

		currentTime := time.Now().Format("20060102")

		fileName := fmt.Sprintf("logs/%s.txt", currentTime)
		f, _ := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		defer f.Close()
		f.WriteString(errorFormating + "\n")
	}

	defer appinsights.TrackPanic(client, false)
}

func GenerateErrorCode() string {

	currentTime := time.Now().Format("20060102150405")

	errorCode := fmt.Sprintf("[%s] %s", commonHelpers.GetConfigurationStringValue("app.name"), currentTime)

	return errorCode
}
