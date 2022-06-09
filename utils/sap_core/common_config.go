package sap_core

import (
	config "example/config/env"
)

func GetSAPCoreBaseURL() string {
	config := config.NewViperConfig()
	baseURL := config.GetString("third_part.sap_core.base_url")

	return baseURL
}
