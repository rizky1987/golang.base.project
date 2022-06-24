package message_bus_receive

import (
	"encoding/json"
	config "example/config/env"
	messageBusLibrary "example/utils/message_bus/library"

	"gopkg.in/mgo.v2"
)

type ProductMessageBusCredentialData struct {
	NameSpace       string
	KeyName         string
	KeyValue        string
	SubcriptionName string
}

type DummyStruct struct {
	Nama   string `json:nama`
	Alamat string `json:alamat`
}

func getProductMessageBusCredential() ProductMessageBusCredentialData {
	config := config.NewViperConfig()

	mesageBusCredential := messageBusLibrary.GetCredential()

	return ProductMessageBusCredentialData{
		NameSpace:       mesageBusCredential.NameSpace,
		KeyValue:        mesageBusCredential.KeyValue,
		KeyName:         mesageBusCredential.KeyName,
		SubcriptionName: config.GetString("third_part.service_bus.catalog_topic_and_subcription.product_subcription_name"),
	}

}

func ReceiveCreateProduct(mongoDBSession *mgo.Session, databaseName string) {

	config := config.NewViperConfig()
	topicName := config.GetString("third_part.service_bus.catalog_topic_and_subcription.create_product_topic_name")

	productMesageBusCredential := getProductMessageBusCredential()

	client := messageBusLibrary.New(messageBusLibrary.Topic,
		productMesageBusCredential.NameSpace, productMesageBusCredential.KeyName, productMesageBusCredential.KeyValue)
	client.SetSubscription(productMesageBusCredential.SubcriptionName)

	for {

		message, err := client.PeekLockMessage(topicName, 10)

		if err != nil {

			// do something with the error
			// return message bus does not exist or not connected
		}

		if message != nil && message.Body != nil {

			// begin put your code here

			jsonData := []byte(message.Body)
			var dummyStruct *DummyStruct

			var err = json.Unmarshal(jsonData, &dummyStruct)
			if err != nil {
				return
			}

			// end put your code here
		}
	}
}
