package responses

import (
	"gopkg.in/mgo.v2/bson"
)

type ProductResponse struct {
	Id                     bson.ObjectId `json:"id"`
	ProductCode            string        `json:"productCode"`
	DosageDescription      string        `json:"dosageDescription"`
	UsabilityDescription   string        `json:"usabilityDescription"`
	CompositionDescription string        `json:"composition"`
	HowToUseDescription    string        `json:"howToUseDescription"`
}

type ProductSuccessResponse struct {
	CommonBaseResponse
	Data ProductResponse `json:"data"`
}

type ProductFailedResponse struct {
	CommonBaseResponse
}
