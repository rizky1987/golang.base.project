package message_bus_library

import (
	config "example/config/env"
)

type MessageBusCredentialData struct {
	NameSpace string
	KeyName   string
	KeyValue  string
}

func GetCredential() MessageBusCredentialData {
	config := config.NewViperConfig()
	return MessageBusCredentialData{
		NameSpace: config.GetString("third_part.service_bus.credential.namespace"),
		KeyName:   config.GetString("third_part.service_bus.credential.key_name"),
		KeyValue:  config.GetString("third_part.service_bus.credential.key_value"),
	}
}
