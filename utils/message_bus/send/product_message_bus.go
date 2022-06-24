package message_bus_send

import (
	config "example/config/env"
	messageBusLibrary "example/utils/message_bus/library"

	"github.com/michaelbironneau/asbclient"
)

type ProductMessageBusCredentialData struct {
	NameSpace       string
	KeyName         string
	KeyValue        string
	SubcriptionName string
}

func getProductMessageBusCredential() ProductMessageBusCredentialData {
	config := config.NewViperConfig()

	mesageBusCredential := messageBusLibrary.GetCredential()

	return ProductMessageBusCredentialData{
		NameSpace:       mesageBusCredential.NameSpace,
		KeyValue:        mesageBusCredential.KeyValue,
		KeyName:         mesageBusCredential.KeyName,
		SubcriptionName: config.GetString("third_part.service_bus.catalog_topic_and_subcription.create_product_topic_name"),
	}

}

func SendCreateProduct(byteofObject []byte) error {
	config := config.NewViperConfig()
	mesageBusCredential := getProductMessageBusCredential()
	topicName := config.GetString("third_part.service_bus.catalog_topic_and_subcription.subcription_name")

	client := asbclient.New(asbclient.Topic, mesageBusCredential.NameSpace, mesageBusCredential.KeyName, mesageBusCredential.KeyValue)
	client.SetSubscription(mesageBusCredential.SubcriptionName)

	err := client.Send(topicName, &asbclient.Message{
		Body: byteofObject,
	})

	if err != nil {
		return err
	}

	return nil
}
