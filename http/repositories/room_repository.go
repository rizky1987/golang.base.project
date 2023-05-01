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
						rp.Price as RoomPricePrice,
						rp.Id as RoomTypeId,
						rp.Code as RoomTypeCode,
						rp.[Type] as RoomTypeName 
					FROM Room r
						left join RoomPrice rp on rp.Id = r.RoomPriceId
						left join [Floor] f on rp.FloorId = f.Id
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
	floorNumber, roomNumber int, roomTypeName string, startPrice, endPrice int) ([]*SQLEntity.TempAvaibilityRoom, error) {

	var err error
	var results []*SQLEntity.TempAvaibilityRoom

	additionalQuery := additionalQueryAvailibilityRoom(floorNumber, roomNumber, roomTypeName, startPrice, endPrice)

	err = ctx.DB.Raw(`SELECT 
						masterData.FloorNumber,
						masterData.RoomNumber,
						masterData.RoomPriceType,
						masterData.RoomPricePrice,
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
							r.Id as RoomId,
							r.Number as RoomNumber,
							rp.[Type] as RoomPriceType,
							rp.Price as RoomPricePrice
						FROM Room r 
							left join RoomPrice rp on rp.Id = r.RoomPriceId
							left join [Floor] f on rp.FloorId = f.Id
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
	startPrice, endPrice int) string {

	result := ""
	addtionalQueries := []string{}
	if floorNumber > 0 {
		addtionalQueries = append(addtionalQueries, fmt.Sprintf("masterData.FloorNumber = %d", floorNumber))
	}

	if roomNumber > 0 {
		addtionalQueries = append(addtionalQueries, fmt.Sprintf("masterData.RoomNumber = %d", roomNumber))
	}

	if roomTypeName != "" {
		addtionalQueries = append(addtionalQueries, fmt.Sprintf("lower(masterData.RoomPriceType) like lower('%s')", roomTypeName))
	}

	if startPrice > 0 && endPrice > 0 {
		addtionalQueries = append(addtionalQueries, fmt.Sprintf("masterData.RoomPricePrice Between %d and %d", startPrice, endPrice))
	}

	if len(addtionalQueries) > 0 {
		result = fmt.Sprintf("WHERE %s", strings.Join(addtionalQueries, " AND "))
	}

	return result
}
