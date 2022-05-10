package responses

import (
	"gopkg.in/mgo.v2/bson"
)

type UserResponse struct {
	Id               bson.ObjectId `json:"id"`
	Image            string        `json:"image"`
	FullName         string        `json:"fullname"`
	PhoneNumber      string        `json:"phone_number"`
	RT               int           `json:"rt"`
	HomeNumberId     bson.ObjectId `json:"home_number_id"`
	HomeBlockId      bson.ObjectId `json:"home_block_id"`
	SecurityGroupId  bson.ObjectId `json:"security_group_id"`
	IsSecurityAdmin  bool          `json:"is_security_admin"`
	IsHeadOfBlock    bool          `json:"is_head_of_block"`
	IsHeadOfCitizens bool          `json:"is_head_of_citizens"`
	AdditionalInfo   string        `json:"additional_info"`
	Token            string        `json:"token,omitempty"`
}

type UserSuccessResponse struct {
	CommonBaseResponse
	Data UserResponse `json:"data"`
}

type UserFailedResponse struct {
	CommonBaseResponse
}
