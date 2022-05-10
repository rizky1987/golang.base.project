package utils

import (
	"fmt"
	"os"
	"time"

	config "example/config/env"

	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

func SaveErrorToApplicationInsight(errorCode, errorType, errorMessage, fileLocation string, fileLine int) {
	config := config.NewViperConfig()

	instrumentationkey := config.GetString("third_part.application_insight.instrumentationkey")
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

	isApplicationOnDebugMode := config.GetBool("app.debug")

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

	config := config.NewViperConfig()

	currentTime := time.Now().Format("20060102150405")

	errorCode := fmt.Sprintf("[%s] %s", config.GetString("app.name"), currentTime)

	return errorCode
}
