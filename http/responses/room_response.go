package responses

import (
	"hotel/commonHelpers"
	SQLEntity "hotel/databases/entities/sql"
)

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

type Floor struct {
	FloorNumber int     `json:"floorNumber"`
	Rooms       []*Room `json:"rooms"`
}

type Room struct {
	RoomNumber   int    `json:"roomNumber"`
	RoomTypeName string `json:"roomTypeName"`
	FloorPrice   int    `json:"price"`
	IsBooked     bool   `json:"isBooked"`
}

type TempRoomSuccessResponse struct {
	ITotalRecords        int      `json:"totalRecords"`
	ITotalDisplayRecords int      `json:"displayRecords"`
	Floors               []*Floor `json:"floors"`
}

func AddFloorResponse(data *SQLEntity.TempAvaibilityRoom) *Floor {

	floorResponse := &Floor{
		FloorNumber: data.FloorNumber,
	}

	return floorResponse
}

func AddRoomResponse(data *SQLEntity.TempAvaibilityRoom) *Room {

	floorResponse := &Room{
		RoomNumber:   data.RoomNumber,
		RoomTypeName: data.RoomTypeName,
		FloorPrice:   data.FloorPrice,
		IsBooked:     commonHelpers.ConvertIntegerToBoolen(data.IsBooked),
	}

	return floorResponse
}

func ConvertAvailibilityRoomEntityToResponse(TempAvaibilityRooms []*SQLEntity.TempAvaibilityRoom) *TempRoomSuccessResponse {

	result := &TempRoomSuccessResponse{
		ITotalDisplayRecords: len(TempAvaibilityRooms),
		ITotalRecords:        len(TempAvaibilityRooms),
	}
	previousFloorNumber := -1
	floorNumberCount := -1

	previousRoomNumber := -1
	roomCount := -1

	for _, data := range TempAvaibilityRooms {

		if previousFloorNumber == -1 || previousFloorNumber != data.FloorNumber {
			previousFloorNumber = data.FloorNumber
			roomCount = -1

			// begin add floor
			floorResponse := AddFloorResponse(data)

			result.Floors = append(result.Floors, floorResponse)
			floorNumberCount++
			// end add floor

			// begin add room
			roomResponse := AddRoomResponse(data)

			result.Floors[floorNumberCount].Rooms = append(result.Floors[floorNumberCount].Rooms, roomResponse)

			roomCount++
			previousRoomNumber = data.RoomNumber
			// end add room

		} else {

			if previousRoomNumber == -1 || previousRoomNumber != data.RoomNumber {
				// begin add room
				roomResponse := AddRoomResponse(data)

				result.Floors[floorNumberCount].Rooms = append(result.Floors[floorNumberCount].Rooms, roomResponse)

				roomCount++
				// end add floor
			} else {
				continue
			}

		}
	}

	return result
}
