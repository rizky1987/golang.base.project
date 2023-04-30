package responses

type FloorResponse struct {
	ID         string `json:"Id"`
	RoomTypeId string `json:"room_type_id"`
	Number     string `json:"number"`
	Price      string `json:"price"`
}

type FloorSuccessResponse struct {
	CommonBaseResponse
	Data FloorResponse `json:"data"`
}

type FloorFailedResponse struct {
	CommonBaseResponse
}
