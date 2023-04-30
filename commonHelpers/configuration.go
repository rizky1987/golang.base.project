package commonHelpers

import (
	config "hotel/config/env"
	"os"
	"strconv"
)

var configuration = config.NewViperConfig()
var isRunOnLocal = configuration.GetBool("app.is_run_on_local")

func getEnvironmentVariableStringValue(input string) string {

	input = StringReplace(input, ".", "__")
	environmentVariableValue := os.Getenv(input)
	if environmentVariableValue != "" {
		return environmentVariableValue
	}

	return ""
}

func getEnvironmentVariableBoolValue(input string) bool {

	input = StringReplace(input, ".", "__")
	environmentVariableValue := os.Getenv(input)

	if environmentVariableValue != "" {
		boolValue, err := strconv.ParseBool(environmentVariableValue)
		if err != nil {
			return false
		}

		return boolValue
	}

	return false
}

func getEnvironmentVariableIntValue(input string) int {

	input = StringReplace(input, ".", "__")
	environmentVariableValue := os.Getenv(input)

	if environmentVariableValue != "" {
		integerValue, err := strconv.Atoi(environmentVariableValue)
		if err != nil {
			return 0
		}

		return integerValue
	}

	return 0
}

func GetConfigurationIntegerValue(input string) int {

	if isRunOnLocal {
		return configuration.GetInt(input)
	}

	return getEnvironmentVariableIntValue(input)
}

func GetConfigurationStringValue(input string) string {

	if isRunOnLocal {
		return configuration.GetString(input)
	}

	return getEnvironmentVariableStringValue(input)
}

func GetConfigurationBoolValue(input string) bool {

	if isRunOnLocal {
		return configuration.GetBool(input)
	}

	return getEnvironmentVariableBoolValue(input)
}
