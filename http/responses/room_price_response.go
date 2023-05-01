package responses

type RoomPriceResponse struct {
	ID    string `json:"Id"`
	Code  string `json:"code"`
	Type  string `json:"type"`
	Price string `json:"price"`
}

type RoomPriceSuccessResponse struct {
	CommonBaseResponse
	Data FloorResponse `json:"data"`
}

type RoomPriceFailedResponse struct {
	CommonBaseResponse
}
