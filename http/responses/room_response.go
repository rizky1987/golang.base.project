package responses

type RoomResponse struct {
	ID      string `json:"Id"`
	FloorId string `json:"floorId"`
	Code    string `json:"code"`
	Number  int    `json:"number"`
}

type RoomSuccessResponse struct {
	CommonBaseResponse
	Data FloorResponse `json:"data"`
}

type RoomFailedResponse struct {
	CommonBaseResponse
}
