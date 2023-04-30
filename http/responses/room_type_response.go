package responses

type RoomTypeResponse struct {
	ID   string `json:"Id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type RoomTypeSuccessResponse struct {
	CommonBaseResponse
	Data FloorResponse `json:"data"`
}

type RoomTypeFailedResponse struct {
	CommonBaseResponse
}
