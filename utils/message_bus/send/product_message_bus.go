package message_bus_send

import (
	"encoding/json"
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

// begin example for Send Data to Message bus with object type struct
// 1. Ini Contoh Dummy Struct
type DummyStruct struct {
	Nama   string `json:nama`
	Alamat string `json:alamat`
}

//2. fungsi ini hanya contoh saja. Sebenernya fungsi SendCreateProduct
//ini bisa di panggil dr tempait lain dalam folder "http"
func example() {

	dataToSend := DummyStruct{
		Nama:   "Rizky Unyu",
		Alamat: "Jeddah",
	}

	byteofObject, err := json.Marshal(dataToSend)
	if err != nil {
		// save to log
	}

	SendCreateProduct(byteofObject)
}

// end example for Send Data to Message bus with object type struct

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
