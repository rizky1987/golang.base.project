package repositories

import (
	"errors"
	"fmt"
	SQLEntity "hotel/databases/entities/sql"
	interfaces "hotel/http/interfaces"
	"strings"

	"gorm.io/gorm"
)

type respositoryRoom struct {
	DB *gorm.DB
}

func NewRoomRepository(db *gorm.DB) interfaces.RoomInterface {
	return &respositoryRoom{db}
}

func (ctx respositoryRoom) Create(transaction *gorm.DB, data SQLEntity.Room) (*SQLEntity.Room, error) {
	var err error

	if transaction != nil {
		res := transaction.Create(&data)
		err = res.Error
	} else {
		res := ctx.DB.Create(&data)
		err = res.Error
	}
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (ctx respositoryRoom) GetRoomByCode(code string) (*SQLEntity.Room, error) {
	var room *SQLEntity.Room
	res := ctx.DB.Where(
		&SQLEntity.Room{
			Code: code,
		},
	).First(&room)
	err := res.Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return room, nil
}

func (ctx respositoryRoom) GetRoomDetailByRoomIds(roomIds []string) ([]*SQLEntity.TempRoomDetail, error) {
	var err error
	var results []*SQLEntity.TempRoomDetail

	err = ctx.DB.Raw(`SELECT 
						r.Id as RoomId,
						r.Code as RoomCode,
						r.Number as RoomNumber,
						f.Id as FloorId,
						f.Number as FloorNumber,
						f.Price as FloorPrice,
						rt.Id as RoomTypeId,
						rt.Code as RoomTypeCode,
						rt.[Name] as RoomTypeName 
					FROM Room r
						left join [Floor] f on r.FloorId = f.Id
						left join [RoomType] rt on f.RoomTypeId = rt.Id
						WHERE r.Id in (?) `, strings.Join(roomIds, ",")).
		Find(&results).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return results, nil
}

func (ctx respositoryRoom) GetAvailibilityRooms(stardDate, endDate string,
	floorNumber, roomNumber int, roomTypeName string, startFloorPrice, endfloorPrice int) ([]*SQLEntity.TempAvaibilityRoom, error) {

	var err error
	var results []*SQLEntity.TempAvaibilityRoom

	additionalQuery := additionalQueryAvailibilityRoom(floorNumber, roomNumber, roomTypeName, startFloorPrice, endfloorPrice)

	err = ctx.DB.Raw(`SELECT 
						masterData.FloorNumber,
						masterData.RoomNumber,
						masterData.RoomTypeName,
						masterData.FloorPrice,
						(
							CASE
								WHEN masterData.RoomId = transactionData.RoomId 
								THEN 1
								ELSE 0
							END
						) as IsBooked					
					FROM
					(
						SELECT 
						f.Number as FloorNumber,
						f.Price as FloorPrice,
						r.Code as RoomCode,
						r.Id as RoomId,
						r.Number as RoomNumber,
						rt.[Name] as RoomTypeName
						FROM Room r 
						left join [Floor] f on f.Id = r.FloorId
						left join RoomType rt on rt.Id = f.RoomTypeId
					)  masterData
					left join
					(
						SELECT 
							bd.* 
						FROM Booking b 
							left join BookingDetail bd on b.Id = bd.BookingId
						WHERE (CONVERT(DATE,b.StartDate) <=  (?) and CONVERT(DATE,b.EndDate) >=  (?))
							or (CONVERT(DATE,b.StartDate) <=  (?) and CONVERT(DATE,b.EndDate) >=  (?))
					) transactionData
					on masterData.RoomId = transactionData.RoomId
					`+additionalQuery+`
					ORDER BY masterData.FloorNumber, masterData.RoomNumber `,
		stardDate, stardDate, endDate, endDate).Find(&results).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return results, nil
}

func additionalQueryAvailibilityRoom(floorNumber, roomNumber int, roomTypeName string,
	startFloorPrice, endfloorPrice int) string {

	result := ""
	addtionalQueries := []string{}
	if floorNumber > 0 {
		addtionalQueries = append(addtionalQueries, fmt.Sprintf("masterData.FloorNumber = %d", floorNumber))
	}

	if roomNumber > 0 {
		addtionalQueries = append(addtionalQueries, fmt.Sprintf("masterData.RoomNumber = %d", roomNumber))
	}

	if roomTypeName != "" {
		addtionalQueries = append(addtionalQueries, fmt.Sprintf("lower(masterData.RoomTypeName) like lower('%s')", roomTypeName))
	}

	if startFloorPrice > 0 && endfloorPrice > 0 {
		addtionalQueries = append(addtionalQueries, fmt.Sprintf("masterData.FloorPrice Between %d and %d", startFloorPrice, endfloorPrice))
	}

	if len(addtionalQueries) > 0 {
		result = fmt.Sprintf("WHERE %s", strings.Join(addtionalQueries, " AND"))
	}

	return result
}
